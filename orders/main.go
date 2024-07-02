package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/matizaj/oms/common"
	"github.com/matizaj/oms/common/broker"
	"github.com/matizaj/oms/common/discovery/consul"
	pb "github.com/matizaj/oms/common/proto"
	"google.golang.org/grpc"
)

var (
	serviceName = "orders"
	consulAddr = common.EnvString("CONSUL_ADDR", "localhost:8500")
	grpcAddr="localhost:50051"
	amqpUser = "guest"
	amqpPass ="guest"
	amqpHost = "localhost"
	amqpPort = "5672"
)

// type server struct {	
// 	pb.OrderServiceServer
// }
func main() {
	instanceId :=  strconv.Itoa(rand.Int())
	registry, err := consul.NewRegistry(consulAddr, serviceName)
	if err != nil {
		log.Fatalf("failed to register service %v\n", err)
	}

	if err := registry.Register(context.Background(),instanceId, serviceName, grpcAddr); err != nil {
		log.Fatalf("failed to register gateway %v\n", err)
	}
	go func(){
		for {
			if err := registry.HealthCheck(instanceId, serviceName); err != nil {
				log.Printf("health check failed %v\n", err)
			}
			time.Sleep(time.Second *3)
		}
	}()

	defer registry.Unregister(context.Background(), instanceId, serviceName)
	
	store := NewStore()
	service := NewOrderService(store)
	service.CreateOrder(context.Background())


	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen on grpc port %v\n", err)
	}
	defer l.Close()

	channel, close := broker.ConnectBroker(amqpUser, amqpPass, amqpHost, amqpPort)
	defer func() {
		close()
		channel.Close()
	}()
	
	grpcServer := grpc.NewServer()
	grpcHandler := NewGrpcHandler(service, channel)
	pb.RegisterOrderServiceServer(grpcServer, grpcHandler)
	
	log.Print("gRPC server is running")
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to run gRPC Server %v\n", err)
	}
}