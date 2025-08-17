package rcfile

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/send", HandleSendRequest)
	return router
}
