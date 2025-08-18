package sysinfo

import (
	"net/http"

	"google.golang.org/protobuf/proto"
)

func HandleGetRequest(w http.ResponseWriter, r *http.Request) {
	if !config.Enabled {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	info, err := Get()
	if err != nil {
		http.Error(w, "Failed to get system info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := proto.Marshal(info)
	if err != nil {
		http.Error(w, "Failed to marshal system info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/protobuf")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
