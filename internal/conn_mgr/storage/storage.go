package storage

type ConnectionStorage interface {
	SetConnection(clientID string, conn ConnectionEntry) error
	GetConnection(clientID string) (ConnectionEntry, error)
}
