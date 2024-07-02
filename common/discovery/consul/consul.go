package consul

import (
	"context"
	"errors"
	"strconv"
	"strings"

	consul "github.com/hashicorp/consul/agent/consul"
)

type Registry struct {
	client *consul.Client
}

func NewRegistry(addr, serviceName string) (*Registry, error) {
	config := consul.DefaultConfig()
	config.Address = addr

	client, err := consul.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &Registry{client}, nil
}

func (r *Registry) Register(ctx context.Context, instanveId, serverName, hostPort string) error {
	parts := strings.Split(hostPort, ":")
	if len(parts)!=2 {
		return errors.New("invalid host:port format")
	}

	port, err := strconv.Atoi(parts[1])
	if err!=nil {
		return err
	}

	host := parts[0]

	return r.client.Agent().ServiceRegister(&consul.AgentServiceRegistration{
		ID: instanveId,
		Address: host,
		Port: port,
		Name: serverName,
	}), nil
}