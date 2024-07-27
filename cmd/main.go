package main

import (
	"log"

	"github.com/KONICCO/go-industrious-boilerplate/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
