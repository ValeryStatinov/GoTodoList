package main

import (
	"todolist/src/server"

	"github.com/joho/godotenv"
)

type P struct {
	PR  int    `json:"projects"`
	FUU string `json:"fuck"`
}

func main() {
	godotenv.Load(".env")

	server.StartServer()
}
