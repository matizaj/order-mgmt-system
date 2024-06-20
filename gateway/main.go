package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/matizaj/oms/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/matizaj/oms/common/api"
)

var (
	addr = common.EnvString("GTW_ADDR", ":5000")
	orderServiceAddr = common.EnvString("GTW_ADDR", ":9000")
)

func main() {
	conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	log.Println("Dialing orders service at ", orderServiceAddr)
	client := pb.NewOrderServiceClient(conn)

	mux := http.NewServeMux()
	handler := NewHandler(client)
	handler.registerRoutes(mux)
	
	log.Printf("Gateway is up and running on port %s \n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal("gateway can start")
	}
}