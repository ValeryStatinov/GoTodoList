package main

import (
	"os"
	"todolist/src/server"
	"todolist/src/systemlogger"

	"github.com/joho/godotenv"
)

func main() {
	mode := os.Getenv("MODE")
	if mode != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			systemlogger.Log("Failed to read .env")
		}
	}

	server.StartServer()
}
