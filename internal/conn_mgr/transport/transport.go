package transport

type Transport interface {
	Deliever(recipient interface{}, msg []byte) error // not clear how to represent recipient
}
