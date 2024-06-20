package main

import (
	"log"
	"net/http"
)

const addr = ":5000"

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)
	
	log.Printf("Gateway is up and running on port %s \n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("gateway can start")
	}
}