package api

import (
	"net/http"

	"github.com/Wartemelon/TODO-list/pkg/db"
)

func getTaskHandler(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id == "" {
		writeJson(res, map[string]string{"error": "id can not be empty"})
		return
	}

	t, err := db.GetTask(id)
	if err != nil {
		writeJson(res, map[string]string{"error": err.Error()})
		return
	}

	writeJson(res, t)
}
