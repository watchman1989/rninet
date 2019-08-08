package etcd


import (
	"context"
	"errors"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"projects/rninet/registry"
	"sync"
	"time"
)

const (
	SERVICE_CHAN_LENGTH = 10
	ETCD_SERVICE_PATH = "/SERVICE/"
)


var (
	etcdRegistry *EtcdRegistry = &EtcdRegistry{
		registerChannel: make(chan *registry.Service, SERVICE_CHAN_LENGTH),
	}
)


type EtcdRegistry struct {
	options *registry.Options
	client *clientv3.Client
	lock sync.Mutex
	registerChannel chan *registry.Service
}


func init () {

	registry.RegisterPlugin(etcdRegistry)

	go etcdRegistry.run()
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



func (e *EtcdRegistry) run () {

	for {
		select {
		case service := <-e.registerChannel:
			service = service
		}
	}
}


