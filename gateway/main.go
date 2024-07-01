package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"
	"github.com/matizaj/oms/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "github.com/matizaj/oms/common/proto"
)
var  (
	gtwAddr = common.EnvString("GTW_ADDR", ":7001")
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:50051")
)

func main() {

	grpcConn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect grpc server %v\n", err)
	}
	defer grpcConn.Close()

	log.Printf("gRPC client connected to server %s", grpcAddr)

	grpcClient := pb.NewOrderServiceClient(grpcConn)

	mux := http.NewServeMux()
	handler := NewHttpHandler(grpcClient)
	handler.registerRoutes(mux)

	log.Println("Gateway service is up and running on port ", gtwAddr)

	if err := http.ListenAndServe(gtwAddr, mux); err != nil {
		log.Panicf("failed to start gateway service %v\n", err)
	}
}