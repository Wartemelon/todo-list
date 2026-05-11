package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Wartemelon/TODO-list/pkg/api"
)

func Run() {
	curDir, err := os.Getwd()
	fmt.Println(curDir)
	if err != nil {
		log.Fatal(err)
	}

	webDir := filepath.Join(curDir, "web")

	api.Init()
	http.Handle("/", http.FileServer(http.Dir(webDir)))

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = ":7540"
	}

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal("Server failed to start: ", err)
		return
	}
}
