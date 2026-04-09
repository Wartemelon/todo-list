package api

import (
	"encoding/json"
	"log"
	"net/http"
)

const dateFormat = "20060102"

func Init() {
	http.HandleFunc("/api/nextdate", nextDateHandler)
	http.HandleFunc("/api/task", auth(taskHandler))
	http.HandleFunc("/api/tasks", auth(tasksHandler))
	http.HandleFunc("/api/task/done", auth(taskDoneHandler))
	http.HandleFunc("/api/signin", signinHandler)
}

func writeJson(res http.ResponseWriter, data any, code int) {
	resp, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		return
	}

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.WriteHeader(code)
	_, err = res.Write(resp)
	if err != nil {
		log.Println(err)
		return
	}
}
