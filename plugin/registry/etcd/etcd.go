package etcd


import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/watchman1989/rninet/common/utils"
	"github.com/watchman1989/rninet/plugin"
	"github.com/watchman1989/rninet/plugin/registry"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/mvcc/mvccpb"
	"path"
	"sync"
	"time"
)

const (
	DEFAULT_CHANNEL_LENGTH = 10
	RNINET_SERVICE_PATH = "/SERVICE/"
)

// SERVICE_APTH: /SERVICE/name/id

var (
	etcdRegistry *EtcdRegistry = &EtcdRegistry{
		registerChannel: make(chan *registry.Service, DEFAULT_CHANNEL_LENGTH),
		deregisterChannel: make(chan *registry.Service, DEFAULT_CHANNEL_LENGTH),
		allServices: make(map[string]map[string]string),
		lidMap: make(map[string]clientv3.LeaseID),
	}
)


type EtcdRegistry struct {
	options *registry.Options
	client *clientv3.Client
	lock sync.Mutex
	registerChannel chan *registry.Service
	deregisterChannel chan *registry.Service
	allServices map[string]map[string]string
	lidMap map[string]clientv3.LeaseID
}


func init () {

	fmt.Println("ETCD_INIT")

	plugin.InstallPlugin("registry", etcdRegistry)

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

	e.GetAllServices()

	fmt.Println(e.allServices)

	go etcdRegistry.watch()
	go etcdRegistry.listen()

	return nil
}


func (e *EtcdRegistry) Register (ctx context.Context, service *registry.Service) error {

	select {
	case e.registerChannel <- service:
		fmt.Println("PUSH_INFO_TO_REGISTER_CHANNEL")
	default:
		return errors.New("REGISTER_CHANNEL_FULL")

	}

	return nil
}


func (e *EtcdRegistry) Deregister (ctx context.Context, service *registry.Service) error {

	fmt.Println("CALL_DEREGISTER")

	select {
	case e.deregisterChannel <- service:
		fmt.Println("PUSH_INFO_TO_DEREGISTER_CHANNEL")
	default:
		return errors.New("DEREGISTER_CHANNEL_FULL")
	}

	return nil
}


func (e *EtcdRegistry) QueryService (ctx context.Context, name string) (map[string]*registry.Service, error) {
	fmt.Println("ALL_SERVICE: ", e.allServices)
	srvInfo, ok := e.allServices[name]
	if !ok {
		fmt.Printf("SERVICE_NOT_EXISTS\n")
		return nil, errors.New("SERVICE_NOT_EXISTS")
	}

	srvMap := make(map[string]*registry.Service)
	for srvId, srvVal := range srvInfo {
		srv := registry.Service{}
		if err := json.Unmarshal([]byte(srvVal), &srv); err != nil {
			fmt.Println("UNMARSHAL_ERROR: ", err)
			continue
		}
		srvMap[srvId] = &srv
	}

	return srvMap, nil
}


func (e *EtcdRegistry) SyncService (ctx context.Context, name string) (chan map[string]*registry.Service) {

	srvChan := make(chan map[string]*registry.Service, DEFAULT_CHANNEL_LENGTH)
	srvMap := make(map[string]*registry.Service)
	srvMap, err := e.QueryService(context.TODO(), name)
	if err == nil {
		srvChan <- srvMap
	}

	go func() {
		watchPath := fmt.Sprintf("%s%s/", RNINET_SERVICE_PATH, name)
		wch := e.client.Watch(context.Background(), watchPath, clientv3.WithPrefix())
		for wresp := range wch {
			if len(wresp.Events) != 0 {
				rsp, err := e.client.Get(context.TODO(), watchPath, clientv3.WithPrefix())
				if err != nil {
					fmt.Printf("ETCD_GET_ERROR: %v\n", err)
					continue
				}

				for _, ev := range rsp.Kvs {
					fmt.Printf("%s: %s\n", ev.Key, ev.Value)
					srvId := path.Base(string(ev.Key))
					srv := registry.Service{}
					if err := json.Unmarshal(ev.Value, &srv); err != nil {
						continue
					}
					srvMap[srvId] = &srv
				}
				srvChan <- srvMap
			}
		}

	}()

	return srvChan
}



func (e *EtcdRegistry) watch () {

	fmt.Println("START_WATCH: ", RNINET_SERVICE_PATH)

	wch := e.client.Watch(context.Background(), RNINET_SERVICE_PATH, clientv3.WithPrefix())
	for wresp := range wch {
		e.lock.Lock()
		for _, ev := range wresp.Events {
			fmt.Printf("%s %s: %s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			srvId := path.Base(string(ev.Kv.Key))
			srvName := path.Base(path.Dir(string(ev.Kv.Key)))
			if false {

			} else if ev.IsCreate() {
				fmt.Println("IS_CREATE")
				if _, ok := e.allServices[srvName]; ok {
					if _, ok := e.allServices[srvName][srvId]; ok {

					} else {
						e.allServices[srvName][srvId] = string(ev.Kv.Value)
					}

				} else {
					e.allServices[srvName] = make(map[string]string)
					e.allServices[srvName][srvId] = string(ev.Kv.Value)
				}

			} else if ev.IsModify() {
				fmt.Println("IS_MODIFY")
				e.allServices[srvName][srvId] = string(ev.Kv.Value)
			} else if ev.Type == mvccpb.DELETE {
				fmt.Println("IS_DELETE")
				delete(e.allServices[srvName], srvId)
				if len(e.allServices[srvName]) == 0 {
					delete(e.allServices, srvName)
				}
			} else {

			}
		}
		e.lock.Unlock()
	}
}



func (e *EtcdRegistry) listen () {

	fmt.Println("START_LISTEN", e)

	for {
		select {
		case service := <-e.registerChannel:
			fmt.Println("GET_REGGISTER_INFO")
			go e.ProcessRegister(service)
		case service := <-e.deregisterChannel:
			fmt.Println("GET_DEREGGISTER_INFO")
			e.ProcessDeregister(service)
		}
	}
}


func (e *EtcdRegistry) GetAllServices() {

	rsp, err := e.client.Get(context.TODO(), RNINET_SERVICE_PATH, clientv3.WithPrefix())
	if err != nil {
		fmt.Printf("ETCD_GET_ERROR: %v\n", err)
		return
	}

	for _, ev := range rsp.Kvs {
		fmt.Printf("%s: %s\n",  ev.Key, ev.Value)

		srvId := path.Base(string(ev.Key))
		//fmt.Println(srvId)
		srvName := path.Base(path.Dir(string(ev.Key)))
		//fmt.Println(srvName)

		if _, ok := e.allServices[srvName]; ok {
			if _, ok := e.allServices[srvName][srvId]; ok {

			} else {
				e.allServices[srvName][srvId] = string(ev.Value)
			}

		} else {
			e.allServices[srvName] = make(map[string]string)
			e.allServices[srvName][srvId] = string(ev.Value)
		}

	}
}


func (e *EtcdRegistry) ProcessRegister (service *registry.Service) {

	fmt.Println("PROCESS_REGISTER:", service.Name, service.Addr)

	srvName := service.Name
	srvId := utils.GetMd5(service.Addr)
	srvValue, _ := json.Marshal(*service)

	rsp, err := e.client.Grant(context.TODO(), e.options.TTL)
	if err != nil {
		fmt.Printf("GRANT_ERROR: %v\n", err)
		return
	}

	leaseID := rsp.ID

	_, err = e.client.Put(context.TODO(), fmt.Sprintf("%s%s/%s", RNINET_SERVICE_PATH, srvName, srvId), string(srvValue),  clientv3.WithLease(leaseID))
	if err != nil {
		fmt.Printf("PUT_ERROR: %v\n", err)
		return
	}

	ka, err := e.client.KeepAlive(context.TODO(), rsp.ID)
	if err != nil {
		fmt.Printf("KEEPALIVE_ERROR: %v\n", err)
		return
	}
	e.lock.Lock()
	e.lidMap[srvId] = leaseID
	e.lock.Unlock()
	for {
		select {
		case _, ok := <-ka:
			if ok {
				//fmt.Printf("KEEPALIVE: %s %s %d\n", srvName, srvId, rsp.TTL)
			} else {
				/*
				e.lock.Lock()
				if _, ok := e.lidMap[srvId]; ok {
					fmt.Printf("REVOKE: %s %s\n", srvName, srvId)
					if _, err = e.client.Revoke(context.TODO(), leaseID); err != nil {
						fmt.Printf("REVOKE_ERROR: %v\n", err)
					}
				}
				e.lock.Unlock()
				*/
				return
			}
		}
	}
}


func (e *EtcdRegistry) ProcessDeregister (service *registry.Service) {

	fmt.Println("PROCESS_DEREGISTER:", service.Name, service.Addr)
	srvName := service.Name
	srvId := utils.GetMd5(service.Addr)

	e.lock.Lock()
	defer e.lock.Unlock()
	//fmt.Println(e.lidMap)
	leaseID, ok := e.lidMap[srvId]
	if ok {
		delete(e.lidMap, srvId)
		fmt.Printf("REVOKE: %s %s\n", srvName, srvId)
		if _, err := e.client.Revoke(context.TODO(), leaseID); err != nil {
			fmt.Printf("REVOKE_ERROR: %v\n", err)
		}
	}
}
