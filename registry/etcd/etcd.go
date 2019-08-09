package etcd


import (
	"context"
	"errors"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/v3/mvcc/mvccpb"
	"projects/rninet/registry"
	"sync"
	"time"
)

const (
	SERVICE_CHAN_LENGTH = 10
	SERVICE_PATH = "/SERVICE/"
)

// SERVICE_APTH: /SERVICE/name/id

var (
	etcdRegistry *EtcdRegistry = &EtcdRegistry{
		registerChannel: make(chan *registry.Service, SERVICE_CHAN_LENGTH),
		allServices: make(map[string]string),
	}
)


type EtcdRegistry struct {
	options *registry.Options
	client *clientv3.Client
	lock sync.Mutex
	registerChannel chan *registry.Service
	allServices map[string]string
}


func init () {

	fmt.Println("ETCD_INIT")

	registry.RegisterPlugin(etcdRegistry)

	fmt.Println("ETCD_INIT_OVER")
}


func (e *EtcdRegistry) Name () string {

	return "etcd"
}


func (e *EtcdRegistry) Init (ctx context.Context, opts ...registry.Option) error {
	var (
		err error
		opt registry.Option
	)
	e.options = &registry.Options{}
	for _, opt = range opts {
		opt(e.options)
	}

	if e.client, err = clientv3.New(clientv3.Config{Endpoints: e.options.Addrs, DialTimeout: time.Duration(e.options.Timeout) * time.Second,}); err != nil {
		fmt.Printf("ETCD_CLIENT_NEW_ERROR: %v\n", err)
		return err
	}

	go etcdRegistry.watch()
	go etcdRegistry.listen()

	return nil
}


func (e *EtcdRegistry) Register (ctx context.Context, service *registry.Service) error {

	select {
	case e.registerChannel <- service:
	default:
		return errors.New("REGISTER_CHANNEL_FULL")

	}

	return nil
}


func (e *EtcdRegistry) Deregister (ctx context.Context, service *registry.Service) error {

	return nil
}


func (e *EtcdRegistry) GetService (ctx context.Context, name string) (*registry.Service, error) {

	return nil, nil
}


func (e *EtcdRegistry) watch () {

	fmt.Println("START_WATCH: ", SERVICE_PATH)

	wch := e.client.Watch(context.Background(), SERVICE_PATH, clientv3.WithPrefix())
	for wresp := range wch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %s: %s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)

			if false {

			} else if ev.IsCreate() {
				fmt.Println("IS_CREATE")
			} else if ev.IsModify() {
				fmt.Println("IS_MODIFY")
			} else if ev.Type == mvccpb.DELETE {
				fmt.Println("IS_DELETE")
			} else {

			}
		}
	}
}



func (e *EtcdRegistry) listen () {

	fmt.Println("START_LISTEN", e)

	time.Sleep(20 * time.Second)

	for {
		select {
		case service := <-e.registerChannel:
			service = service
		}
	}
}




