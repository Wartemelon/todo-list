package api

import (
	"net/http"
	"time"

	"github.com/Wartemelon/TODO-list/pkg/service"
)

func nextDateHandler(res http.ResponseWriter, req *http.Request) {
	nowStr := req.FormValue("now")
	var now time.Time
	var err error

	if nowStr == "" {
		now = time.Now()
	} else {
		now, err = time.Parse(dateFormat, nowStr)
		if err != nil {
			http.Error(res, "invalid 'now' date format, must be YYYYMMDD", http.StatusBadRequest)
			return
		}
	}

	dstart := req.FormValue("date")
	repeat := req.FormValue("repeat")
	next, err := service.NextDate(now, dstart, repeat)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "text/plain")
	res.Write([]byte(next))
}
