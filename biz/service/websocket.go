package service

import (
	"context"
	"encoding/json"
	"strconv"

	"sfw/biz/dal"
	"sfw/biz/dal/model"
	"sfw/biz/mw/jwt"
	"sfw/pkg/errno"
	"sfw/pkg/utils/encrypt"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/hertz-contrib/websocket"
)

type ChatService struct {
	ctx  context.Context
	c    *app.RequestContext
	conn *websocket.Conn
}

type _user struct {
	username string
	conn     *websocket.Conn
	rsa      *encrypt.RsaService
}

var userMap = make(map[int64]*_user)

func NewChatService(ctx context.Context, c *app.RequestContext, conn *websocket.Conn) *ChatService {
	return &ChatService{
		ctx:  ctx,
		c:    c,
		conn: conn,
	}
}

func (service *ChatService) Login() error {
	token := string(service.c.GetHeader("Access-Token"))
	uids, err := jwt.AccessTokenJwtMiddleware.ExtractPayloadFromToken(token)
	if err != nil {
		return err
	}
	uid, _ := strconv.ParseInt(uids, 10, 64)
	u := dal.Executor.User
	user, err := u.WithContext(context.Background()).Where(u.ID.Eq(uid)).First()
	if err != nil {
		return err
	}
	rsaClientKey := service.c.GetHeader(`rsa_public_key`)
	r := encrypt.NewRsaService()
	if err := r.Build(rsaClientKey); err != nil {
		hlog.Info(err)
		return errno.InternalServerError
	}
	userMap[uid] = &_user{conn: service.conn, username: user.Username, rsa: r}
	publicKey, err := r.GetPublicKeyPemFormat()
	if err != nil {
		return errno.InternalServerError
	}
	if err := service.conn.WriteMessage(websocket.TextMessage, []byte(publicKey)); err != nil {
		return errno.InternalServerError
	}

	return nil
}

func (service *ChatService) Logout() error {
	token := string(service.c.GetHeader("Access-Token"))
	uids, err := jwt.AccessTokenJwtMiddleware.ExtractPayloadFromToken(token)
	if err != nil {
		return err
	}
	uid, _ := strconv.ParseInt(uids, 10, 64)
	userMap[uid] = nil
	return nil
}

func (service *ChatService) ReadOfflineMessage() error {
	token := string(service.c.GetHeader("Access-Token"))
	uids, err := jwt.AccessTokenJwtMiddleware.ExtractPayloadFromToken(token)
	if err != nil {
		return errno.AccessTokenInvalid
	}
	uid, _ := strconv.ParseInt(uids, 10, 64)
	m := dal.Executor.Message
	list, err := m.WithContext(context.Background()).Where(m.ToUserID.Eq(uid)).Find()
	if err != nil {
		return errno.InternalServerError
	}
	toConn := userMap[uid]
	for _, item := range list {
		jsonMessage, err := json.Marshal(*item)
		if err != nil {
			return errno.InternalServerError
		}
		ciphertext, err := toConn.rsa.Encode(jsonMessage)
		if err != nil {
			return errno.InternalServerError
		}
		err = service.conn.WriteMessage(websocket.BinaryMessage, ciphertext)
		if err != nil {
			return errno.InternalServerError
		}
	}
	return nil
}

func (service *ChatService) PushOnlineMessage(message *model.Message) error {
	toConn := userMap[message.ToUserID]
	if toConn == nil {
		return errno.UserOffline
	}
	jsonMessage, err := json.Marshal(*message)
	if err != nil {
		return errno.InternalServerError
	}
	ciphertext, err := toConn.rsa.Encode(jsonMessage)
	if err != nil {
		return errno.InternalServerError
	}
	err = toConn.conn.WriteMessage(websocket.BinaryMessage, ciphertext)
	if err != nil {
		return errno.InternalServerError
	}
	return nil
}

func isUserOnline(uid int64) bool {
	return userMap[uid] != nil
}
