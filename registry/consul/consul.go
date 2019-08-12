package consul

import (
	"context"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc/health/grpc_health_v1"
	"projects/rninet/registry"
)

type ConsulRegistry struct {
	options *registry.Options
	client *consulapi.Client
}

type Health struct {}


func (h *Health) Check (ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {

	return &grpc_health_v1.HealthCheckResponse {
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (h *Health) Watch (req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {

	return nil
}


func (c *ConsulRegistry) Name () string {
	return "consul"
}


func (c *ConsulRegistry) Init (ctx context.Context, opts ...registry.Option) error {

	var (
		err error
		opt registry.Option
	)

	c.options = &registry.Options{}
	for _, opt = range opts {
		opt(c.options)
	}

	config := consulapi.DefaultConfig()
	config.Address = c.options.Addrs[0]

	c.client, err = consulapi.NewClient(config)
	if err != nil {
		fmt.Printf("CONSUL_CLIENT_NEW_ERROR: %v\n", err)
		return err
	}

	return nil
}



func (c *ConsulRegistry) Register (ctx context.Context, service *registry.Service) error {

	tags := []string{}
	for key, _ := range service.Metadata {
		tags = append(tags, key)
	}

	reg := &consulapi.AgentServiceRegistration{
		Name: service.Name,
		Address: service.Ip,
		Port: service.Port,
		Tags: tags,
		Check: &consulapi.AgentServiceCheck{
			Interval: fmt.Sprintf("%d", c.options.Interval),
		},
	}

	agent := c.client.Agent()
	if err := agent.ServiceRegister(reg); err != nil {
		fmt.Printf("SERVICE_REGISTER_ERROR: %v\n", err)
		return err
	}

	return nil
}


func (c *ConsulRegistry) Deregister (ctx context.Context, service *registry.Service) error {


	return nil
}



func (c *ConsulRegistry) QueryService (ctx context.Context, name string) map[string]*registry.Service {


	return nil
}


func (c *ConsulRegistry) SyncService (ctx context.Context, name string) chan map[string]*registry.Service {


	return nil
}
