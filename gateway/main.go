package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/matizaj/oms/common"
	pb "github.com/matizaj/oms/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
var (
	webPort = common.EnvString("HTTP_ADDR", ":3030")
	orderServiceAddr = "localhost:3000"
)

func main() {
	conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("GRPC server can start ", err)
	}
	defer conn.Close()

	c := pb.NewOrderServiceClient(conn)
	mux := http.NewServeMux()

	handler:= NewHandler(c)
	handler.RegisterRoutes(mux)

	log.Printf("Server is running on port: %s",webPort )

	if err := http.ListenAndServe(webPort, mux); err != nil {
		log.Fatal("Failed to start server", err)
	}

}