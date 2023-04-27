package counter

type Receivable interface {
	Receive(handler func(b byte) uint)
}

type BytesReceiver struct {
	receiving    <-chan byte
	accumulating chan<- uint
}

// NewBytesReceiver factory to initialize bytes receiver
func NewBytesReceiver(
	receiving <-chan byte,
	accumulating chan<- uint) *BytesReceiver {
	return &BytesReceiver{
		receiving:    receiving,
		accumulating: accumulating,
	}
}

// Receive Receiving bytes and perform
// the given logic for the received bytes
func (r BytesReceiver) Receive(handler func(b byte) uint) {
	defer close(r.accumulating)
	for b := range r.receiving {
		r.accumulating <- handler(b)
	}
}
