package messaging

import (
	"context"
	"log"

	"github.com/mixdjoker/chain-notes/internal/app/commitservice"
	"github.com/nats-io/nats.go"
)

type NATSClient struct {
	nc *nats.Conn
}

func New(nc *nats.Conn) *NATSClient {
	return &NATSClient{
		nc: nc,
	}
}

func (n *NATSClient) Subscribe(
	ctx context.Context,
	subject,
	queue string,
	handler func(commitservice.IncomingMessage),
) error {
	sub, err := n.nc.QueueSubscribe(subject, queue, func(m *nats.Msg) {
		wrapped := &natsMessage{msg: m}
		handler(wrapped)
	})
	if err != nil {
		return err
	}

	log.Printf("[nats] subscribed to %s", subject)

	go func() {
		<-ctx.Done()
		_ = sub.Unsubscribe()
	}()

	return nil
}

type natsMessage struct {
	msg *nats.Msg
}

func (m *natsMessage) Data() []byte {
	return m.msg.Data
}

func (m *natsMessage) Respond(data []byte) error {
	return m.msg.Respond(data)
}
