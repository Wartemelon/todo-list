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
		writeJson(res, map[string]string{"error": err.Error()})
		return
	}

	if t.Repeat == "" {
		db.DeleteTask(id)
		writeJson(res, map[string]any{})
	} else {
		next, err := service.NextDate(time.Now(), t.Date, t.Repeat)
		if err != nil {
			writeJson(res, map[string]string{"error": err.Error()})
			return
		}
		db.UpdateDate(next, id)
		writeJson(res, map[string]any{})
	}

}
