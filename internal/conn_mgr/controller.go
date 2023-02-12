package conn_mgr

import (
	"fmt"

	"github.com/ganzz96/gdev-cloud/internal/conn_mgr/storage"
	"github.com/ganzz96/gdev-cloud/internal/conn_mgr/transport"
)

type controller struct {
	connStorage storage.ConnectionStorage
}

func New(connStorage storage.ConnectionStorage) *controller {
	return &controller{connStorage: connStorage}
}

func (c *controller) Connect(req transport.ConnectionRequest) error {
	// need to be implemented

	fmt.Println("New connection has been established")
	return nil
}

func (c *controller) Deliver(message string, toClientIDs []string, guarantee transport.DelieveryOption) error {
	// need to be implemented

	return nil
}
