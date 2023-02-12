package api

import (
	"encoding/json"
	"net/http"

	"github.com/ganzz96/gdev-cloud/internal/conn_mgr/api/generated"
	"github.com/ganzz96/gdev-cloud/internal/conn_mgr/transport"
)

type connector interface {
	Connect(req transport.ConnectionRequest) error
	Deliver(message string, toClientIDs []string, guarantee transport.DelieveryOption) error
}

type controller struct {
	connector   connector
	establisher transport.Establisher
}

func New(connector connector, establisher transport.Establisher) generated.ServerInterface {
	c := &controller{
		connector:   connector,
		establisher: establisher,
	}

	return c
}

func (c *controller) GetApiV1Connect(w http.ResponseWriter, r *http.Request) {
	conn, err := c.establisher.EstablishConnection(w, r)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, generated.Error{Message: err.Error()})
		return
	}

	if err = c.connector.Connect(transport.ConnectionRequest{Connection: conn}); err != nil {
		writeJSON(w, http.StatusInternalServerError, generated.Error{Message: err.Error()})
		return
	}
}

func (c *controller) PostApiV1Deliver(w http.ResponseWriter, r *http.Request) {
	var req generated.DeliveryRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, generated.Error{Message: err.Error()})
		return
	}

	if err := c.connector.Deliver(req.Message, req.Clients, transport.DelieveryOption(req.DeliveryType)); err != nil {
		writeJSON(w, http.StatusInternalServerError, generated.Error{Message: err.Error()})
		return
	}
}

func writeJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(code)

	b, err := json.Marshal(v)
	if err != nil {
		return
	}

	if _, err := w.Write(b); err != nil {
		return
	}
}
