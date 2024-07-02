package main

import (
	"context"
	"log"
	"net"

	pb "github.com/matizaj/oms/common/proto"
	"google.golang.org/grpc"
)
const grpcAddr="localhost:50051"
var serviceName = "orders"

type server struct {	
	pb.OrderServiceServer
}
func main() {

	store := NewStore()
	service := NewOrderService(store)
	service.CreateOrder(context.Background())


	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen on grpc port %v\n", err)
	}
	defer l.Close()
	
	grpcServer := grpc.NewServer()
	server := server{}
	pb.RegisterOrderServiceServer(grpcServer, &server)
	
	log.Print("gRPC server is running")
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to run gRPC Server %v\n", err)
	}
}