package main

import (
	"log"

	"github.com/gkits/kurz/internal/app"
)

func main() {
	a := app.New()

	log.Fatal(a.Run())
}
