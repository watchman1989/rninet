package etcd



import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)


var (
	etcdComponent *EtcdComponent = &EtcdComponent{}
)

func NewEtcdComponent () *EtcdComponent {
	return etcdComponent
}


type EtcdComponent struct {
	Name string
	Client *clientv3.Client
}


type EtcdStarter struct {
	options *Options
}


func (e *EtcdStarter) Init (ctx context.Context, opts ...interface{}) error {

	var (
		config clientv3.Config
		err error
		client *clientv3.Client
	)

	e.options = &Options{}
	for _, opt := range opts {
		opt.(Option)(e.options)
	}

	fmt.Println("ETCD_START_INIT")

	config = clientv3.Config{
		Endpoints: e.options.Addrs,
		DialTimeout: time.Duration(e.options.Timeout) * time.Second,
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Printf("ETCD_NEW_ERROR: %v\n", err)
		return err
	}

	etcdComponent.Client = client
	etcdComponent.Name = "etcd"

	return nil
}


func (e *EtcdStarter) Stop () error {
	return nil
}


func (e *EtcdComponent) Put (key string, val string) error {

	_, err := e.Client.Put(context.Background(), key, val)
	if err != nil {
		fmt.Printf("ETCD_PUT_ERROR: %v\n", err)
		return err
	}

	return nil
}


func (e *EtcdComponent) Del (key string) error {

	_, err := e.Client.Delete(context.Background(), key)
	if err != nil {
		fmt.Printf("ETCD_DEL_ERROR: %v\n", err)
		return err
	}

	return nil
}


func (e *EtcdComponent) Get (key string) (string, error) {

	rsp, err := e.Client.Get(context.Background(), key)
	if err != nil {
		fmt.Printf("ETCD_GET_ERROR: %v\n", err)
		return "", err
	}

	fmt.Println(rsp.Kvs)
	for _, ev := range(rsp.Kvs) {
		fmt.Println(ev.Key, ev.Value)
	}

	return string(rsp.Kvs[0].Value), nil
}


func (e *EtcdComponent) GetWithPrefix (prefix string) (map[string]string, error) {

	rsp, err := e.Client.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		fmt.Printf("ETCD_GET_ERROR: %v\n", err)
		return nil, err
	}

	fmt.Println(rsp.Kvs)
	for _, ev := range(rsp.Kvs) {
		fmt.Println(ev.Key, ev.Value)
	}

	return nil, nil
}