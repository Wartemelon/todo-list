package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Wartemelon/TODO-list/pkg/db"
	"github.com/Wartemelon/TODO-list/pkg/service"
)

func addTaskHandler(res http.ResponseWriter, req *http.Request) {
	var task db.Task
	var buf bytes.Buffer

	_, err := buf.ReadFrom(req.Body)
	if err != nil {
		writeJson(res, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	if err = json.Unmarshal(buf.Bytes(), &task); err != nil {
		writeJson(res, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	if task.Title == "" {
		writeJson(res, map[string]string{"error": "title can not be empty"}, http.StatusBadRequest)
		return
	}

	err = service.CheckDate(&task)
	if err != nil {
		writeJson(res, map[string]string{"error": err.Error()}, http.StatusBadRequest)
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeJson(res, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}
	writeJson(res, map[string]int64{"id": id}, http.StatusCreated)
}
