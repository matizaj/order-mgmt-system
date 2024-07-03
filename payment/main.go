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
	stripeProcessor "github.com/matizaj/oms/payment/processors/stripe"
	"github.com/stripe/stripe-go/v79"
	"google.golang.org/grpc"
	_ "github.com/joho/godotenv/autoload"
)

var (
	serviceName = "payment"
	consulAddr = common.EnvString("CONSUL_ADDR", "localhost:8500")
	grpcAddr = common.EnvString("GRPC_ADDR","localhost:50052")
	amqpUser = common.EnvString("AMQP_USER","guest")
	amqpPass = common.EnvString("AMQP_PASS","guest")
	amqpHost = common.EnvString("AMQP_HOST","localhost")
	amqpPort = common.EnvString("AMQP_PORT","5672")
	stripeKey = common.EnvString("STRIPE_KEY","")
)

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

	//stripe setup
	stripe.Key = stripeKey
	

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

	stripeProcessor := stripeProcessor.NewStripeProcessor()
	service := NewPaymentService(stripeProcessor)
	amqpConsumer := NewConsumer(service)
	amqpConsumer.Listen(channel)
	
	grpcServer := grpc.NewServer()
	
	log.Print("gRPC server is running")
	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("failed to run gRPC Server %v\n", err)
	}
}