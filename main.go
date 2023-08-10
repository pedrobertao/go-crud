package main

import (
	"log"
	"os"

	"github.com/pedrobertao/go-crud/database"
	"github.com/pedrobertao/go-crud/logging"
	"github.com/pedrobertao/go-crud/server"
)

func main() {
	if err := logging.Start(); err != nil {
		log.Fatal(err)
	}
	defer logging.Logger.Sync()

	database.Start()
	defer database.Shutdown()

	if err := server.Start(":3030"); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}
