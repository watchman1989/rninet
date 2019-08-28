package main

import (
	"context"
	"fmt"
	"github.com/watchman1989/rninet/plugin"
	"github.com/watchman1989/rninet/plugin/registry"
	"time"
)

func main() {
	reg, err := plugin.InitRegistry(
		context.TODO(),
		"etcd",
		registry.WithAddrs([]string{"127.0.0.1:2379"}),
		registry.WithTimeout(3),
		registry.WithTTL(3),
		registry.WithInterval(1),
	)


	fmt.Println("INIT_REGISTRY_OVER")

	fmt.Println("PULUGINS: ", plugin.GetPlugins())

	if err != nil {
		fmt.Printf("INIT_REGISTRY_ERROR: %v\n", err)
		return
	}


	service := &registry.Service{Name: "test.srv", Addr: "http://127.0.0.1:8080"}

	reg.Register(context.TODO(), service)


	time.Sleep(5 * time.Second)

	reg.Deregister(context.TODO(), service)



	s1 := &registry.Service{Name: "s1", Addr: "http://127.0.0.1:8080"}
	reg.Register(context.TODO(), s1)
	s2 := &registry.Service{Name: "s1", Addr: "http://127.0.0.1:8081"}
	reg.Register(context.TODO(), s2)
	s3 := &registry.Service{Name: "s2", Addr: "http://127.0.0.2:8080"}
	reg.Register(context.TODO(), s3)
	s4 := &registry.Service{Name: "s3", Addr: "http://127.0.0.1:8082"}
	reg.Register(context.TODO(), s4)
	s5 := &registry.Service{Name: "s3", Addr: "http://127.0.0.1:8083"}
	reg.Register(context.TODO(), s5)

	time.Sleep(5 * time.Second)
	mapSrv, _ := reg.QueryService(context.TODO(), "s1")
	fmt.Println("QUERY_SERVICE:", mapSrv)


	go func() {

		srvChan := reg.SyncService(context.TODO(), "s3")
		for srvMap := range srvChan {
			fmt.Println("SYNC_SRV: ", srvMap)
		}

	}()

	time.Sleep(2 * time.Second)
	reg.Deregister(context.TODO(), s5)

	s6 := &registry.Service{Name: "s3", Addr: "http://127.0.0.1:8088"}
	reg.Register(context.TODO(), s6)
}
