package main

import (
	"context"
	"log"
	"net"

	"github.com/matizaj/oms/common"
	"google.golang.org/grpc"
)
var (
	grpcAddr=common.EnvString("GRPC_ADDR", ":9000")
)

func main() {
	grpcServer := grpc.NewServer()
	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	NewGrpcHandler(grpcServer)
	
	log.Println("Grpc Server is up and running, listen on port: ", grpcAddr)

	orderStore := NewOrderStore()
	orderSvc := NewOrderService(orderStore)
	orderSvc.CreateOrder(context.Background())

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal("grpc server failed to start")
	}
}