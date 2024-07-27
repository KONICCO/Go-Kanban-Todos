package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func main() {
	// Create a new instance of the struct
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	//use ../.env because main.go inside /cmd
	err = godotenv.Load(filepath.Join(pwd, "../.env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
