package websocket

import (
	"context"
	"fmt"
	"sfw/biz/mw/jwt"
	"sfw/biz/service"
	"sfw/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/websocket"
)

var upgrader = websocket.HertzUpgrader{}

var (
	badMessage = `{"code": %v, "msg": "%v"}`
	successMsg = `{"code": 0, "msg": "success"}`
)

func Handler(ctx context.Context, c *app.RequestContext) {
	err := upgrader.Upgrade(c, func(conn *websocket.Conn) {
		_, err := jwt.CovertJWTPayloadToString(ctx, c)
		if err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(badMessage,
				errno.AccessTokenInvalidErrorCode, errno.AccessTokenInvalidErrorMsg)))
			return
		}
		conn.WriteMessage(websocket.TextMessage, []byte(successMsg))

		s := service.NewChatService(ctx, c, conn)

		if err := s.Login(); err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(badMessage,
				errno.InternalServerError, err.Error())))
			return
		}
		defer s.Logout()

		if err := s.ReadOfflineMessage(); err != nil {
			conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(badMessage,
				errno.InternalServerError, err.Error())))
			return
		}

		for {
			_, _, err := conn.ReadMessage()
			if err != nil {
				conn.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf(badMessage,
					errno.InternalServerError, err.Error())))
				return
			}

			conn.WriteMessage(websocket.TextMessage, []byte("you shouldn't send message to me"))
		}
	})

	if err != nil {
		c.JSON(consts.StatusOK, utils.H{
			"code": errno.InternalServerError,
			"msg":  err.Error(),
		})
		return
	}
}
