package sysinfo

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/get", HandleGetRequest)
	return router
}
