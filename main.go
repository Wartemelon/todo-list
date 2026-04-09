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
		dbFile = "scheduler.db"
	}

	err := db.Init(dbFile)
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	server.Run()
}
