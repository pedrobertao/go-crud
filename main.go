package main

import (
	"log"
	"os"

	"github.com/pedrobertao/go-crud/server"
)

func main() {
	if err := server.Start(":3030"); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
