package main

import (
	"context"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/matizaj/oms/common"
	"github.com/matizaj/oms/common/discovery/consul"
	pb "github.com/matizaj/oms/common/proto"
	"github.com/matizaj/oms/gateway/gateway"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
var  (
	gtwAddr = common.EnvString("GTW_ADDR", ":7001")
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:50051")
	consulAddr = common.EnvString("CONSUL_ADDR", "localhost:8500")
	serviceName = "gateway"
)

func main() {

	// grpcConn, err := grpc.NewClient(grpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("failed to connect grpc server %v\n", err)
	// }
	// defer grpcConn.Close()

	// log.Printf("gRPC client connected to server %s", grpcAddr)
	instanceId :=  strconv.Itoa(rand.Int())
	registry, err := consul.NewRegistry(consulAddr, serviceName)
	if err != nil {
		log.Fatalf("failed to register service %v\n")
	}

	if err := registry.Register(context.Background(),instanceId, serviceName, gtwAddr); err != nil {
		log.Fatalf("failed to register gateway %v\n")
	}

	go func(){
		for {
			if err := registry.HealthCheck(instanceId, serviceName); err != nil {

			}
			time.Sleep(time.Second *3)
		}
	}()

	defer registry.Unregister(context.Background(), instanceId, serviceName)

	ordersGateway := gateway.NewGrpcGateway(registry)

	mux := http.NewServeMux()
	handler := NewHttpHandler()
	handler.registerRoutes(mux)

	log.Println("Gateway service is up and running on port ", gtwAddr)

	if err := http.ListenAndServe(gtwAddr, mux); err != nil {
		log.Panicf("failed to start gateway service %v\n", err)
	}
}