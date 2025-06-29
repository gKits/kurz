package main

import (
	"log"

	fs "github.com/gkits/kurz"
	"github.com/gkits/kurz/config"
	"github.com/pressly/goose/v3"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("migrate: failed to load config: %e", err)
	}

	log.Println(cfg)

	db, err := cfg.Database.Open()
	if err != nil {
		log.Fatalf("migrate: failed to open database connection: %e", err)
	}

	goose.SetBaseFS(fs.Migrations())
	if err := goose.SetDialect(cfg.Database.Driver.String()); err != nil {
		log.Fatalf("migrate: failed to set dialect: %e", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
}
