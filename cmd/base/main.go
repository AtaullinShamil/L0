package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/AtaullinShamil/L0/pkg/config"
	"github.com/AtaullinShamil/L0/pkg/db"
	nats2 "github.com/AtaullinShamil/L0/pkg/nats"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"log"
	"sync"
)

var configPath string = ""

func init() {
	flag.StringVar(&configPath, "config", "/Users/shamil/Desktop/L0/deploy", "config folder")
	flag.Parse()
}

func main() {
	ctx := context.Background()

	var cfg config.Config
	err := confita.NewLoader(
		file.NewBackend(fmt.Sprintf("%s/default.yaml", configPath)),
		env.NewBackend(),
	).Load(ctx, &cfg)
	if err != nil {
		fmt.Printf("failed to parse config: %s\n", err.Error())
		return
	}
	//db
	postgres, err := db.NewDatabase(cfg.Postgres)
	if err != nil {
		log.Fatalf("failed to connect postgresql: %s\n", err.Error())
	}

	err = postgres.Init(ctx)
	if err != nil {
		log.Fatalf("failed to init db: %s\n", err.Error())
	}

	//nats
	nc, err := nats2.NewNats(nats.DefaultURL)
	if err != nil {
		log.Fatalf("failed to connect nats: %s\n", err.Error())
	}
	defer nc.Close()

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
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
	if _, err := ec.Subscribe("updates", func(s *db.Order) {
		fmt.Println(s)
		wg.Done()
	}); err != nil {
		log.Fatal(err)
	}

	// Wait for a message to come in
	wg.Wait()
}
