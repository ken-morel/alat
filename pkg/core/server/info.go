package server

import (
	"fmt"
	"net/http"
)

func handleInfo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Add("Alat-Device", "true")
	w.Header().Add("Content-Length", fmt.Sprintf("%d", len(infoResponse)))

	w.Write(infoResponse)
}
