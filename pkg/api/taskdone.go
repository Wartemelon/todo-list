package api

import (
	"net/http"
	"time"

	"github.com/Wartemelon/TODO-list/pkg/db"
	"github.com/Wartemelon/TODO-list/pkg/service"
)

func taskDoneHandler(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	t, err := db.GetTask(id)
	if err != nil {
		writeJson(res, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}

	if t.Repeat == "" {
		if err := db.DeleteTask(id); err != nil {
			writeJson(res, map[string]string{"error": err.Error()}, http.StatusBadRequest)
			return
		}
		writeJson(res, map[string]any{}, http.StatusOK)
		return
	}
	next, err := service.NextDate(time.Now(), t.Date, t.Repeat)
	if err != nil {
		writeJson(res, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}
	if err := db.UpdateDate(next, id); err != nil {
		writeJson(res, map[string]string{"error": err.Error()}, http.StatusInternalServerError)
		return
	}
	writeJson(res, map[string]any{}, http.StatusOK)
}
