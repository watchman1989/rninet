package main

import (
	"context"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"projects/rninet/cmd"
	"projects/rninet/infra"
	"projects/rninet/infra/etcd"
	"projects/rninet/registry"
	//_ "projects/rninet/registry/consul"
	//_ "projects/rninet/registry/etcd"
	"runtime"
	"time"
)



func TestEtcd() {
	reg, err := registry.InitRegistry(
		context.TODO(),
		"etcd",
		registry.WithAddrs([]string{"127.0.0.1:2379"}),
		registry.WithTimeout(3),
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



func TestEtcdCom() {

	infra.AddStarter(context.Background(),
		"etcd",
		&etcd.EtcdStarter{},
		etcd.WithAddrs([]string{"10.42.6.161"}),
		etcd.WithTimeout(3),
		)

	etcdCom := etcd.NewEtcdComponent()
	fmt.Println("ETCD_COM: ", etcdCom)


}


func TestConsul() {

	reg, err := registry.InitRegistry(
		context.TODO(),
		"consul",
		registry.WithAddrs([]string{"127.0.0.1:8500"}),
		registry.WithInterval(5),
		registry.WithTTL(10),
	)

	if err != nil {
		fmt.Printf("INIT_CONSUL_ERROR: %v\n", err)
		return
	}

	fmt.Println("INIT_CONSUL_REGISTRY_OVER")

	fmt.Println("PULUGINS: ", registry.GetPlugins())

	s0 := registry.Service{Name: "test_consul", Ip: "192.168.1.10", Port: 5000}
	reg.Register(context.TODO(), &s0)

	go reg.SyncService(context.TODO(), "test_consul")



	time.Sleep(5 * time.Second)

	s1 := registry.Service{Name: "test_consul", Ip: "192.168.1.14", Port: 5005}
	reg.Register(context.TODO(), &s1)

}


func InitEnv() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}


func InitArg() {

}



func init () {
	InitEnv()
	InitArg()

}




func main () {

	fmt.Println("MAIN_START")

	app := cli.NewApp()
	app.Name = "rninet"
	app.Usage = "A micro service frame"

	app.Commands = cmd.Commands


	if err := app.Run(os.Args); err != nil {
		fmt.Printf("APP_RUN_ERROR: %v\n", err)
	}


	//time.Sleep(1000 * time.Second)

}
