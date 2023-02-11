package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ganzz96/gdev-cloud/internal/conn_mgr/transport"
	ws "github.com/gorilla/websocket"
)

type connector interface {
	Connect(req transport.ConnectionRequest) error
}

type transporter struct {
	upgrader  ws.Upgrader
	connector connector
}

func New(connector connector) *transporter {
	t := &transporter{
		upgrader: ws.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
		connector: connector,
	}

	t.upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	return t
}

func (t *transporter) RegisterEndpoints(connectionURL string) {
	http.HandleFunc(connectionURL, t.connectionEndpoint)
}

// TODO: improve error handling
func (t *transporter) connectionEndpoint(w http.ResponseWriter, r *http.Request) {
	socket, err := t.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	queue := make(chan []byte)

	if err = t.connector.Connect(transport.ConnectionRequest{Sender: queue}); err != nil {
		log.Println(err)
		return
	}

	// TODO: handle channel closing
	for msg := range queue {
		if err := socket.WriteMessage(ws.TextMessage, msg); err != nil {
			log.Println(err)
			return
		}

		fmt.Println("Successful delievery")
	}
}
