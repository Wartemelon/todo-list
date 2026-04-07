package main

import (
	"log"
	"os"

	"github.com/Wartemelon/TODO-list/pkg/db"
	"github.com/Wartemelon/TODO-list/pkg/server"
)

func main() {
	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		log.Fatalln("Environmental value: TODO_DBFILE must be not empty")
	}

	err := db.Init(dbFile)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	server.Run()
}
