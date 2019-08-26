package main

import (
	"context"
	"fmt"
	"github.com/watchman1989/rninet/infra"
	"github.com/watchman1989/rninet/infra/etcd"
)

func main() {
	infra.AddStarter(context.Background(),
		"etcd",
		&etcd.EtcdStarter{},
		etcd.WithAddrs([]string{"127.0.0.1:2379"}),
		etcd.WithTimeout(3),
	)

	etcdCom := etcd.NewEtcdComponent()
	fmt.Println("ETCD_COM: ", etcdCom)
}
