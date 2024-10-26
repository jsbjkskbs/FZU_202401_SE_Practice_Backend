package webs

import (
	"sfw/biz/handler/websocket"

	"github.com/cloudwego/hertz/pkg/app/server"
)

func register(h *server.Hertz) {
	h.GET(`/`, append(_homeMW(), websocket.Handler)...)
}
