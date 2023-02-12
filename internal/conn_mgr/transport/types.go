package transport

import "net/http"

const (
	Unreliable DelieveryOption = iota
	Reliable   DelieveryOption = iota
)

type DelieveryOption int

type ConnectionRequest struct {
	Connection Connection
}

type Connection interface {
	WriteMessage(delieveryOption DelieveryOption, data []byte) error
}

type Establisher interface {
	EstablishConnection(w http.ResponseWriter, r *http.Request) (Connection, error)
}
