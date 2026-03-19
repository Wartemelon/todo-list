package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Run() {
	curDir, err := os.Getwd()
	fmt.Println(curDir)
	if err != nil {
		log.Fatal(err)
	}

	webDir := filepath.Join(curDir, "web")
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
