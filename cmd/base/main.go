package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/AtaullinShamil/L0/pkg/config"
	"github.com/AtaullinShamil/L0/pkg/db"
	"github.com/AtaullinShamil/L0/pkg/handler"
	brocker "github.com/AtaullinShamil/L0/pkg/nats"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	_ "github.com/lib/pq"
	"log"
	"net/http"
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
	//
	//err = postgres.Init(ctx)
	//if err != nil {
	//	log.Fatalf("failed to init db: %s\n", err.Error())
	//}

	//nats
	go brocker.NatsCycle()

	handlers := handler.NewHandler(postgres, fmt.Sprintf("%s:%s", cfg.Host, cfg.Port))
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: handlers.Router(),
	}
	fmt.Println("server started")
	_ = srv.ListenAndServe()
}
