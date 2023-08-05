package main

import (
	"log"

	"github.com/pedrobertao/go-crud/server"
)

func main() {
	if err := server.Start(":3030"); err != nil {
		log.Fatal(err)
	}
}
