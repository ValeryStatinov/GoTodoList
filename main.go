package main

import (
	"todolist/src/server"
	"todolist/src/systemlogger"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		systemlogger.Log("Failed to read .env")
	}
	server.StartServer()
}
