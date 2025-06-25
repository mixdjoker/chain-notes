package natsx

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

// Connect establishes a connection to a NATS server using the provided URL.
func Connect(url string) (*nats.Conn, error) {
	opts := []nats.Option{
		nats.Name("ChainNotes-NATS"),
		nats.MaxReconnects(10),              // До 10 попыток переподключения
		nats.ReconnectWait(2 * time.Second), // Интервал между попытками
		nats.Timeout(5 * time.Second),       // Таймаут на попытку соединения
		nats.RetryOnFailedConnect(true),     // Включает ретраи при первом подключении
		nats.ReconnectHandler(func(_ *nats.Conn) {
			log.Println("[natsx] reconnected to NATS")
		}),
		nats.ClosedHandler(func(_ *nats.Conn) {
			log.Println("[natsx] connection closed")
		}),
		nats.DisconnectErrHandler(func(_ *nats.Conn, err error) {
			log.Printf("[natsx] disconnected due to: %v", err)
		}),
	}

	nc, err := nats.Connect(url, opts...)
	if err != nil {
		return nil, err
	}

	return nc, nil
}
