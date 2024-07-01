package discovery

import "context"

type Registry interface {
	Register(ctx context.Context, instanveId, serverName, hostPort string) error
	Unregister(ctx context.Context, instanveId, serverName string) error
	Dicover(ctx context.Context, serverName string) ([]string, error)
	HealthCheck(instanveId, serverName string) error
}