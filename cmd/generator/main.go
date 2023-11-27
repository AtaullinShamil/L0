package main

import (
	"encoding/json"
	"github.com/AtaullinShamil/L0/pkg/db"
	"github.com/nats-io/nats.go"
	"log"
	"os"
)

func main() {
	file, err := os.Open("/Users/shamil/Desktop/L0/model.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var test db.Order
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&test)
	if err != nil {
		log.Fatal(err)
	}

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Close()

	// Publish the message
	if err := ec.Publish("updates", &test); err != nil {
		log.Fatal(err)
	}
}
