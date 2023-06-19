package main

import (
	"log"

	"github.com/57th-street/oza-gueser/api"
)

func main() {
	api.NewRouter()
	log.Println("server start at port 8080")
}
