package main

import (
	"log"

	"github.com/tapiaw38/concejapi/db"
	"github.com/tapiaw38/concejapi/handlers"
)

func main() {
	if db.CheckConnection() == false {
		log.Fatal("No database connection")

		return
	}

	handlers.HandlerServer()
}
