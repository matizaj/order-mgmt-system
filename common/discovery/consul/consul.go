package consul

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/consul/api"
	consul "github.com/hashicorp/consul/api"
)
type Registry struct {
	client *consul.Client
}

func NewRegistry(addr, serviceName string) (*Registry, error) {
	config := consul.DefaultConfig()
	config.Address = addr

	client, err := consul.NewClient(config)
	if err != nil {
		log.Println("cant connect consul")
		return nil, err
	}

	return &Registry{client}, nil
}

func (r *Registry) Register(ctx context.Context, instanceId, serverName, hostPort string) error {
	log.Printf("Registering service %s", serverName)

	parts := strings.Split(hostPort, ":")
	if len(parts)!= 2 {
		return errors.New("invalid host:port format")
	}
	port, err := strconv.Atoi(parts[1])
	if err != nil {
		return err
	}
	host := parts[0]

	return r.client.Agent().ServiceRegister(&consul.AgentServiceRegistration{
		ID: instanceId,
		Name: serverName,
		Address: host,
		Port: port,
		Check: &consul.AgentServiceCheck{
			CheckID: instanceId,
			TLSSkipVerify: true,
			TTL: "5s",
			Timeout: "1s",
			DeregisterCriticalServiceAfter: "10s",
		},
	})

}

func (r *Registry)	Unregister(ctx context.Context, instanceId, serverName string) error {
	log.Printf("Unregistering service %s", serverName)
	return r.client.Agent().CheckDeregister(instanceId) 
}

func (r *Registry)	Discover(ctx context.Context, serverName string)([]string, error) {
	entries, _, err := r.client.Health().Service(serverName, "", true, nil)
	if err != nil {
		return nil, err
	}
	var instances []string
	for _, entry := range entries {
		instances = append(instances, fmt.Sprintf("%s:%d", entry.Service.Address, entry.Service.Port))
	}
	return instances, nil 
}

func (r *Registry)	HeallhCheck(instanceId, serverName string) error {
	return r.client.Agent().UpdateTTL(instanceId, "online", api.HealthPassing) 
}