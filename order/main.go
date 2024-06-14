package main

import (
	"context"
	"log"
	"net"

	"github.com/matizaj/oms/common"
	"google.golang.org/grpc"
)
var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:3000")
)
func main() {
	grpcServer := grpc.NewServer()

	listen, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer listen.Close()

	orderStore := NewStore()
	orderSvc := NewService(orderStore)
	orderSvc.CreateOrder(context.TODO())

	NewGrpcHandler(grpcServer, orderSvc)

	log.Println("GRPC Server started ", grpcAddr)

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatal(err.Error())
	}
}