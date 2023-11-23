package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"log"
	"sync"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "wb"
	password = "wb"
	dbname   = "wb"
)

func main() {
	//db
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, dbname, password)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Println("cannot connect to db")
		return
	}
	defer db.Close()

	//lifeCheck
	err = db.Ping()
	if err != nil {
		fmt.Println("db doesn't work:ping")
	}

	////nats-streaming
	//nc, err := nats.Connect(nats.DefaultURL)
	//if err != nil {
	//	fmt.Println("nats doens't work")
	//	//log.Fatal(err)
	//}
	//defer nc.Close()
	//
	//// Subscribe
	//sub, err := nc.SubscribeSync("updates")
	//if err != nil {
	//	fmt.Println("cannot subscribe nc")
	//	//log.Fatal(err)
	//}
	//
	//// Wait for a message
	//msg, err := sub.NextMsg(10 * time.Second)
	//if err != nil {
	//	fmt.Println("cannot have messages")
	//	//log.Fatal(err)
	//}
	//
	//fmt.Println(string(msg.Data))

	nc, err := nats.Connect(nats.DefaultURL,
		nats.ErrorHandler(func(nc *nats.Conn, s *nats.Subscription, err error) {
			if s != nil {
				log.Printf("Async error in %q/%q: %v", s.Subject, s.Queue, err)
			} else {
				log.Printf("Async error outside subscription: %v", err)
			}
		}))
	if err != nil {
		fmt.Println("error connection nats")
		//log.Fatal(err)
	}
	defer nc.Close()
	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		fmt.Println("encoding problems")
		//log.Fatal(err)
	}
	defer ec.Close()

	// Define the object
	type stock struct {
		Symbol string
		Price  int
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	// Subscribe
	// Decoding errors will be passed to the function supplied via
	// nats.ErrorHandler above, and the callback supplied here will
	// not be invoked.
	if _, err := ec.Subscribe("updates", func(s *stock) {
		fmt.Printf("Stock: %s - Price: %v\n", s.Symbol, s.Price)
		//log.Printf("Stock: %s - Price: %v", s.Symbol, s.Price)
		wg.Done()
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for a message to come in
	wg.Wait()

	fmt.Println("Woo-Hoo!")
}
