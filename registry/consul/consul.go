package consul

import (
	"context"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"google.golang.org/grpc/health/grpc_health_v1"
	"projects/rninet/registry"
	"time"
)

var (
	consulRegistry *ConsulRegistry = &ConsulRegistry{}
)


type ConsulRegistry struct {
	options *registry.Options
	client *consulapi.Client
}

func init () {
	fmt.Println("CONSUL_INIT")

	registry.RegisterPlugin(consulRegistry)

	fmt.Println("CONSUL_INIT_OVER")
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

	interval := time.Duration(c.options.Interval) * time.Second
	deregister := time.Duration(c.options.TTL) * time.Minute

	reg := &consulapi.AgentServiceRegistration{
		ID: fmt.Sprintf("%s-%s-%d", service.Name, service.Addr, service.Port),
		Name: service.Name,
		Address: service.Ip,
		Port: service.Port,
		Tags: tags,
		Check: &consulapi.AgentServiceCheck{
			Interval: interval.String(),
			GRPC: fmt.Sprintf("%s:%d/%s", service.Addr, service.Port, service.Name),
			DeregisterCriticalServiceAfter: deregister.String(),
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



func (c *ConsulRegistry) QueryService (ctx context.Context, name string) (map[string]*registry.Service, error) {


	return nil, nil
}


func (c *ConsulRegistry) SyncService (ctx context.Context, name string) chan map[string]*registry.Service {
	var lastIndex uint64 = 0
	for {
		services, metainfo, err := c.client.Health().Service(name, "", true, &consulapi.QueryOptions{
			WaitIndex: lastIndex,
		})

		if err != nil {
			fmt.Printf("CLIENT_HEALTH_SERVICE_ERROR: %v\n", err)
		}

		fmt.Println(metainfo.LastIndex)

		lastIndex = metainfo.LastIndex

		fmt.Println(metainfo)

		fmt.Println(services)
	}



	return nil
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