package natsx

import (
	"log"

	"github.com/nats-io/nats.go"
)

// Connect establishes a connection to a NATS server using the provided URL.
func Connect(url string) (*nats.Conn, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	// Optionally set up a connection close handler
	nc.SetClosedHandler(func(_ *nats.Conn) {
		log.Println("[natsx] connection closed")
	})

	nc.SetReconnectHandler(func(_ *nats.Conn) {
		log.Println("[natsx] reconnected to NATS")
	})

	return nc, nil
}
