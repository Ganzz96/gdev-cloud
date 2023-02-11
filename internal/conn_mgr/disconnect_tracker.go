package conn_mgr

type disconnectTracker struct {
}

func NewDisconnectTracker() *disconnectTracker {
	return &disconnectTracker{}
}

func (d *disconnectTracker) Serve() {
	// need to be implemented
}
