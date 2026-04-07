package api

import (
	"net/http"

	"github.com/Wartemelon/TODO-list/pkg/db"
)

func deleteTaskHandler(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		writeJson(res, map[string]string{"error": "id can not be empty"})
		return
	}

	err := db.DeleteTask(id)
	if err != nil {
		writeJson(res, map[string]string{"error": err.Error()})
		return
	}

	writeJson(res, map[string]any{})
}
