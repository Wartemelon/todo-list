package api

import (
	"github.com/Wartemelon/TODO-list/pkg/db"
	"net/http"
	"time"
)

type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(res http.ResponseWriter, req *http.Request) {
	var tasks []*db.Task
	limit := 50

	search := req.URL.Query().Get("search")
	t, err := time.Parse("02.01.2006", search)
	if err == nil {
		date := t.Format("20060102")
		tasks, err = db.TasksByDate(date, limit)
	} else {
		tasks, err = db.TasksByText(search, limit)
	}

	if err != nil {
		writeJson(res, map[string]string{"error": err.Error()})
		return
	}
	writeJson(res, TasksResp{
		Tasks: tasks,
	})
}
