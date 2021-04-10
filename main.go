package main

import (
	"log"

	"github.com/JavDomGom/sn-test/db"
	"github.com/JavDomGom/sn-test/handlers"
)

func main() {
	if !db.CheckConnection() {
		log.Fatal("Something went wrong with the database connection.")
		return
	}
	handlers.Handlers()
}
