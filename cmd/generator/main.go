package main

import (
	"encoding/json"
	"fmt"
	"github.com/AtaullinShamil/L0/pkg/db"
	"github.com/nats-io/nats.go"
	"log"
	"math/rand"
	"os"
	"time"
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

	ticker := time.NewTicker(5 * time.Second)
	for _ = range ticker.C {
		test.OrderUID = randomString(10)
		count := rand.Intn(10) + 1
		if len(test.Items) > count {
			test.Items = test.Items[0 : count+1]
		} else {
			for len(test.Items) < count {
				test.Items = append(test.Items, test.Items[0])
			}
		}
		if err := ec.Publish("updates", &test); err != nil {
			log.Fatal(err)
		}
		fmt.Println(test.OrderUID)
		fmt.Println(len(test.Items))
	}
	// Publish the message

}

func randomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
