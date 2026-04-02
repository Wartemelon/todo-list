package api

import (
	"encoding/json"
	"net/http"
)

const dateFormat = "20060102"

func Init() {
	http.HandleFunc("/api/nextdate", nextDateHandler)
	http.HandleFunc("/api/task", taskHandler)
}

func writeJson(res http.ResponseWriter, data any) error {
	resp, err := json.Marshal(data)
	if err != nil {
		return err
	}

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	res.Write(resp)

	return nil
}
