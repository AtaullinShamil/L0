package brocker

import (
	"fmt"
	"github.com/AtaullinShamil/L0/pkg/db"
	"github.com/nats-io/nats.go"
	"log"
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

//func NatsCycle() {
//	nc, err := NewNats(nats.DefaultURL)
//	if err != nil {
//		log.Fatalf("failed to connect nats: %s\n", err.Error())
//	}
//	defer nc.Close()
//
//	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer ec.Close()
//
//	wg := sync.WaitGroup{}
//	wg.Add(1)
//
//	// Subscribe
//	// Decoding errors will be passed to the function supplied via
//	// nats.ErrorHandler above, and the callback supplied here will
//	// not be invoked.
//	if _, err := ec.Subscribe("updates", func(s *db.Order) {
//		fmt.Println(s)
//		wg.Done()
//	}); err != nil {
//		log.Fatal(err)
//	}
//
//	// Wait for a message to come in
//	wg.Wait()
//}

func NatsCycle() {
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

	// Subscribe
	// Decoding errors will be passed to the function supplied via
	// nats.ErrorHandler above, and the callback supplied here will
	// not be invoked.
	_, err = ec.Subscribe("updates", func(s *db.Order) {
		fmt.Println(s)
	})
	if err != nil {
		log.Fatal(err)
	}

	// Wait for a message to come in
	for {
		time.Sleep(time.Second)
	}
}
