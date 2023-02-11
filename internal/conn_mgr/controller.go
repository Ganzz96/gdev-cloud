package conn_mgr

import (
	"github.com/ganzz96/gdev-cloud/internal/conn_mgr/storage"
	"github.com/ganzz96/gdev-cloud/internal/conn_mgr/transport"
)

type controller struct {
	connStorage storage.ConnectionStorage
	transport   transport.Transport
}

func New(connStorage storage.ConnectionStorage, transport transport.Transport) *controller {
	return &controller{connStorage: connStorage, transport: transport}
}

func (c *controller) Connect(req ConnectionRequest) (Connection, error) {
	// need to be implemented
	return Connection{}, nil
}

func (c *controller) SendMessage(message interface{}, toClientIDs []string, guarantee transport.DelieveryOption) error {
	// need to be implemented
	return nil
}
