package commitservice

import "context"

type Messaging interface {
	Subscribe(ctx context.Context, subject, queue string, handler func(IncomingMessage)) error
}

type IncomingMessage interface {
	Data() []byte
	Respond([]byte) error
}
