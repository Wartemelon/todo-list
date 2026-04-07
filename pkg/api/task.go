package api

import (
	"net/http"
)

func taskHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		addTaskHandler(res, req)
	case http.MethodGet:
		getTaskHandler(res, req)
	case http.MethodPut:
		changeTaskHandler(res, req)
	case http.MethodDelete:
		deleteTaskHandler(res, req)
	}
}
