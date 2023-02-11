package transport

const (
	Reliable   DelieveryOption = iota
	Unreliable DelieveryOption = iota
)

type DelieveryOption int

type ConnectionRequest struct {
	Sender chan []byte
}
