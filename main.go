package main

import (
	"todolist/src/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	server.StartServer()
}
