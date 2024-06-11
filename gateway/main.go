package main

import (
	"log"
	"net/http"
	 _ "github.com/joho/godotenv/autoload"
	"github.com/matizaj/oms/common"
)
var (
	webPort = common.EnvString("HTTP_ADDR", ":3030")
)

func main() {
	mux := http.NewServeMux()

	handler:= NewHandler()
	handler.RegisterRoutes(mux)

	log.Printf("Server is running on port: %s",webPort )

	if err := http.ListenAndServe(webPort, mux); err != nil {
		log.Fatal("Failed to start server", err)
	}

}