package nats2

import (
	"github.com/nats-io/nats.go"
	"log"
)

func NewNats(natsURL string) (*nats.Conn, error) {
	nc, err := nats.Connect(natsURL,
		nats.ErrorHandler(func(nc *nats.Conn, s *nats.Subscription, err error) {
			if s != nil {
				log.Printf("Async error in %q/%q: %v", s.Subject, s.Queue, err)
			} else {
				log.Printf("Async error outside subscription: %v", err)
			}
		}))
	if err != nil {
		return nil, err
	}
	return nc, nil
}
