package discovery

import "context"

type Registry interface {
	Register(ctx context.Context, instanceId, serverName, hostPort string) error
	Unregister(ctx context.Context, instanceId, serverName string) error
	Discover(ctx context.Context, serverName string)([]string, error)
	HeallhCheck(instanceId, serverName string) error
}