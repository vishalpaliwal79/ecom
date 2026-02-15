package main

import (
	"log"
)

func main() {
	server := api.newAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
