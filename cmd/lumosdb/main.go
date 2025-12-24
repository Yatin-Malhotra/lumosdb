package main

import (
	"log"

	"github.com/Yatin-Malhotra/lumosdb/internal/server"
)

func main() {
	srv := server.New(":6379")
	log.Fatal(srv.Start())
}
