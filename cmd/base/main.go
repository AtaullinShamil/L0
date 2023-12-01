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
	"sync"
)

var Path string = ""

func init() {
	flag.StringVar(&Path, "config", "/Users/shamil/Desktop/L0", "config folder")
	flag.Parse()
}

func main() {
	ctx := context.Background()

	var cfg config.Config
	err := confita.NewLoader(
		file.NewBackend(fmt.Sprintf("%s/deploy/default.yaml", Path)),
		env.NewBackend(),
	).Load(ctx, &cfg)
	if err != nil {
		fmt.Printf("failed to parse config: %s\n", err.Error())
		return
	}

	cache := &sync.Map{}
	//db
	postgres, err := db.NewDatabase(cfg.Postgres)
	if err != nil {
		log.Fatalf("failed to connect postgresql: %s\n", err.Error())
	}

	//if there are tables, then the data is stored in the cache, if not, tables are created
	if postgres.CheckTable(ctx, "Orders") {
		err := brocker.Recovery(postgres, cache)
		if err != nil {
			log.Fatalf("failed to recovery: %s\n", err.Error())
		}
	} else {
		err = postgres.Init(ctx)
		if err != nil {
			log.Fatalf("failed to init db: %s\n", err.Error())
		}
	}

	//nats
	go brocker.NatsCycle(ctx, postgres, cache)

	//web
	handlers := handler.NewHandler(postgres, fmt.Sprintf("%s:%s", cfg.Host, cfg.Port), cache, Path)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: handlers.Router(),
	}
	fmt.Println("server started")
	_ = srv.ListenAndServe()
}
