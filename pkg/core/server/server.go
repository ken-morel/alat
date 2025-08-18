// Package server: holds the p2p server
package server

import (
	"alat/pkg/core/service/rcfile"
	"fmt"
	"log"
	"net/http"
)

var Running bool = false

func startServer() {
	Running = true
	fmt.Println("Creating server")
	server := http.NewServeMux()
	server.HandleFunc("/alat-info", handleInfo)
	server.HandleFunc("/pair/request", handlePairRequest)
	server.HandleFunc("/pair/response", handlePairResponse)
	server.HandleFunc("/rcfile/send", rcfile.HandleSendRequest)
	http.Handle("/", server)
	srv := &http.Server{
		Addr:    config.DeviceInfo.Address.String(),
		Handler: server,
	}
	fmt.Println("server listening")
	log.Fatal(srv.ListenAndServe())
	Running = false
}

func Start() {
	go startServer()
}
