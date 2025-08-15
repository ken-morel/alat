// Package server: holds the p2p server
package server

import (
	"fmt"
	"log"
	"net/http"
)

func startServer() {
	fmt.Println("Creating server")
	server := http.NewServeMux()
	server.HandleFunc("/alat-info", handleInfo)
	http.Handle("/", server)
	srv := &http.Server{
		Addr:    "192.168.1.192:60000",
		Handler: server,
	}
	fmt.Println("server listening")
	log.Fatal(srv.ListenAndServe())
}

func Start() {
	go startServer()
}
