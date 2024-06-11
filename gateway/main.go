package main

import (
	"fmt"
	"log"
	"net/http"
)
const webPort="7070"

func main() {
	mux := http.NewServeMux()

	handler:= NewHandler()
	handler.RegisterRoutes(mux)

	if err := http.ListenAndServe(fmt.Sprintf(":%s", webPort), mux); err != nil {
		log.Fatal("Failed to start server", err)
	}

	log.Printf("Server is running on port: %s",webPort )
}