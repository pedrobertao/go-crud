package main

import (
	"fmt"
	"log"

	"github.com/pedrobertao/go-crud/logging"
	auth "github.com/pedrobertao/go-crud/middlewares/authorization"
)

func main() {
	token, err := auth.Generate()
	if err != nil {
		logging.L.Error(err)
		log.Fatal(err)
	}
	fmt.Println(token)
}
