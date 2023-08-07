package main

import (
	"fmt"
	"log"

	auth "github.com/pedrobertao/go-crud/middlewares/Auth"
)

func main() {
	token, err := auth.Generate()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(token)
}
