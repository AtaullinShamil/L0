package brocker

import (
	"context"
	"encoding/json"
	"github.com/AtaullinShamil/L0/pkg/db"
	"github.com/nats-io/nats.go"
	"log"
	"sync"
	"time"
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

func NatsCycle(ctx context.Context, postgres *db.Database, cache *sync.Map) {
	nc, err := NewNats(nats.DefaultURL)
	if err != nil {
		log.Fatalf("failed to connect nats: %s\n", err.Error())
	}
	defer nc.Close()

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Close()

	_, err = ec.Subscribe("updates", func(order *db.Order) {
		err := postgres.SaveOrder(ctx, order)
		if err != nil {
			log.Fatal(err)
		}

		err = postgres.SaveDelivery(ctx, order)
		if err != nil {
			log.Fatal(err)
		}

		err = postgres.SavePayment(ctx, order)
		if err != nil {
			log.Fatal(err)
		}

		err = postgres.SaveItems(ctx, order)
		if err != nil {
			log.Fatal(err)
		}

		jsonData, err := json.Marshal(order)
		if err != nil {
			log.Fatal(err)
		}
		cache.Store(order.OrderUID, jsonData)
	})
	if err != nil {
		log.Fatal(err)
	}

	// Wait for a message to come in
	for {
		time.Sleep(time.Second)
	}
}
