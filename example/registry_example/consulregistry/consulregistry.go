package main

import (
	"context"
	"fmt"
	"github.com/watchman1989/rninet/plugin/registry/consul"
	"github.com/watchman1989/rninet/plugin"
	"github.com/watchman1989/rninet/plugin/registry"
	"time"
)

func main() {

	reg, err := plugin.InitRegistry(
		context.TODO(),
		"consul",
		consul.WithAddrs([]string{"127.0.0.1:8500"}),
		consul.WithInterval(5),
		consul.WithTTL(10),
	)

	if err != nil {
		fmt.Printf("INIT_CONSUL_ERROR: %v\n", err)
		return
	}

	fmt.Println("INIT_CONSUL_REGISTRY_OVER")

	fmt.Println("PULUGINS: ", plugin.GetPlugins())

	s0 := registry.Service{Name: "test_consul", Ip: "192.168.1.10", Port: 5000}
	reg.Register(context.TODO(), &s0)

	go reg.SyncService(context.TODO(), "test_consul")



	time.Sleep(5 * time.Second)

	s1 := registry.Service{Name: "test_consul", Ip: "192.168.1.14", Port: 5005}
	reg.Register(context.TODO(), &s1)
}
