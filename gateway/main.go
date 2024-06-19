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
	pb "github.com/matizaj/oms/common/api"
	"github.com/matizaj/oms/common/discovery/consul"
	"github.com/matizaj/oms/gateway/gateway"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
var (
	webPort = common.EnvString("HTTP_ADDR", ":3030")
	consulAddrs = common.EnvString("CONSUL_ADDR", "localhost:8500")
	serviceName = "gateway"
)

func main() {
	registry, err := consul.NewRegistry(consulAddrs, serviceName)
	if err != nil {
		panic(err)
	}
	instanceId := strconv.Itoa(rand.Int())
	if err := registry.Register(context.Background(), instanceId, serviceName, webPort); err != nil {
		panic(err)
	}
	go func() {
		for {
			if err := registry.HeallhCheck(instanceId, serviceName); err != nil {
				log.Fatal("failed to health check")
			}
			time.Sleep(time.Second*1)
		}
	}()

	defer registry.Unregister(context.Background(), instanceId, serviceName)

	mux := http.NewServeMux()
	ordersGtw := gateway.NewGrpcGateway(registry)
	handler:= NewHandler(ordersGtw)
	handler.RegisterRoutes(mux)

	log.Printf("Server is running on port: %s",webPort )

	if err := http.ListenAndServe(webPort, mux); err != nil {
		log.Fatal("Failed to start server", err)
	}

}