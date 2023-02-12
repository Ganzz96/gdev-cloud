package websocket

import (
	"net/http"

	"github.com/ganzz96/gdev-cloud/internal/conn_mgr/transport"
	ws "github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type transporter struct {
	upgrader ws.Upgrader
}

func New() *transporter {
	t := &transporter{
		upgrader: ws.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}

	t.upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	return t
}

func (t *transporter) EstablishConnection(w http.ResponseWriter, r *http.Request) (transport.Connection, error) {
	socket, err := t.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to upgrade connection to web socket type")
	}

	return &connection{ws: socket}, nil
}
