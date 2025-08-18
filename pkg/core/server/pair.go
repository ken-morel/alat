package server

import (
	"alat/pkg/core/pbuf"
	"fmt"
	"net/http"

	"google.golang.org/protobuf/proto"
)

func handlePairRequest(w http.ResponseWriter, r *http.Request) {
	var length int
	fmt.Sscanf(r.Header.Get("Content-Length"), "%d", &length)
	if length < 10 || length > 1000 {
		http.Error(w, "received invalid content length", http.StatusBadRequest)
		return
	}
	data := make([]byte, length)
	r.Body.Read(data)
	var request pbuf.PairRequest
	err := proto.Unmarshal(data, &request)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(config.OnPairRequest(&request))
}

func handlePairResponse(w http.ResponseWriter, r *http.Request) {
	var length int
	fmt.Sscanf(r.Header.Get("Content-Length"), "%d", &length)
	if length < 10 || length > 1000 {
		http.Error(w, "received invalid content length", http.StatusBadRequest)
		return
	}
	data := make([]byte, length)
	r.Body.Read(data)
	var response pbuf.PairResponse
	err := proto.Unmarshal(data, &response)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	w.WriteHeader(config.OnPairResponse(&response))
}
