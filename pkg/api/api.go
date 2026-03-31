package api

import (
	"net/http"
)

const dateFormat = "20060102"

func Init() {
	http.HandleFunc("/api/nextdate", nextDayHandler)
}
