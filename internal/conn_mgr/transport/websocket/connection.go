package websocket

import (
	"github.com/ganzz96/gdev-cloud/internal/conn_mgr/transport"
	ws "github.com/gorilla/websocket"
)

type connection struct {
	ws *ws.Conn
}

func (c *connection) WriteMessage(delieveryOption transport.DelieveryOption, data []byte) error {
	return c.ws.WriteMessage(ws.TextMessage, data)
}
