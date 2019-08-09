package main

import (
	"context"
	"fmt"
	"projects/rninet/registry"
	_ "projects/rninet/registry/etcd"
	"time"
)

func main () {

	fmt.Println("MAIN_START")

	reg, err := registry.InitRegistry(
		context.TODO(),
		"etcd",
		registry.WithAddrs([]string{"10.42.6.161:2379"}),
		registry.WithTimeout(1),
		registry.WithTTL(3),
		registry.WithInterval(1),
	)


	fmt.Println("INIT_REGISTRY_OVER")

	fmt.Println("PULUGINS: ", registry.GetPlugins())

	if err != nil {
		fmt.Printf("INIT_REGISTRY_ERROR: %v\n", err)
		return
	}


	service := &registry.Service{Name: "test.srv", Addr: "http://127.0.0.1:8080"}

	reg.Register(context.TODO(), service)



	time.Sleep(10 * time.Second)

}
