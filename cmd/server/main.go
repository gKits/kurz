package main

import (
	"log"

	"github.com/gkits/kurz/config"
	"github.com/gkits/kurz/internal/app"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %e", err)
	}

	a, err := app.New(cfg)
	if err != nil {
		log.Fatalf("failed to initialize app: %v", err)
	}

	log.Fatal(a.Run())
}
