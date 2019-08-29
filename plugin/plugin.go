package plugin

import (
	"context"
	"errors"
	"fmt"
	"github.com/watchman1989/rninet/plugin/broker"

	"github.com/watchman1989/rninet/plugin/registry"
	"sync"
)

var (
	pluginManager = &PluginManager{
		registryPlugins: make(map[string]func()registry.Registry),
		brokerPlugins: make(map[string]func()broker.Broker),
	}
)

type PluginManager struct {
	registryPlugins map[string]func()registry.Registry
	brokerPlugins map[string]func()broker.Broker
	lock sync.Mutex
}

func (p *PluginManager) installRegistryPlugin (pluginName string, pluginNewFunc func()registry.Registry) (error) {

	var (
		err error
		ok bool
	)

	p.lock.Lock()
	defer  p.lock.Unlock()

	if _, ok = p.registryPlugins[pluginName]; ok {
		err = errors.New("PLUGIN_IS_EXISTS")
		return err
	}
	p.registryPlugins[pluginName] = pluginNewFunc

	return nil
}

func (p *PluginManager) installBrokerPlugin (pluginName string, pluginNewFunc func()broker.Broker) (error) {

	var (
		err error
		ok bool
	)

	p.lock.Lock()
	defer  p.lock.Unlock()

	if _, ok = p.registryPlugins[pluginName]; ok {
		err = errors.New("PLUGIN_IS_EXISTS")
		return err
	}
	p.brokerPlugins[pluginName] = pluginNewFunc

	return nil
}


func (p *PluginManager) initRegistry (ctx context.Context, name string, opts ...interface{}) (reg registry.Registry, err error) {
	var (
		ok     bool
		regNewFunc func()registry.Registry
	)
	p.lock.Lock()
	defer p.lock.Unlock()

	regNewFunc, ok = p.registryPlugins[name]
	if !ok {
		fmt.Printf("%s NOT_EXISTS\n", name)
		return reg, errors.New("NEW_%s_NOT_EXISTS")
	}

	reg = regNewFunc()
	reg.Init(ctx, opts...)

	return
}


func (p *PluginManager) initBroker (ctx context.Context, name string, opts ...interface{}) (bro broker.Broker, err error) {
	var (
		ok     bool
		broNewFunc func()broker.Broker
	)
	p.lock.Lock()
	defer p.lock.Unlock()

	broNewFunc, ok = p.brokerPlugins[name]
	if !ok {
		fmt.Printf("%s NOT_EXISTS\n", name)
		return bro, errors.New("NEW_%s_NOT_EXISTS")
	}

	bro = broNewFunc()
	bro.Init(ctx, opts...)

	return
}


func InstallRegistryPlugin(pluginName string, newPluginFunc func()registry.Registry) error {

	return pluginManager.installRegistryPlugin(pluginName, newPluginFunc)
}

func InstallBrokerPlugin(pluginName string, newPluginFunc func()broker.Broker) error {

	return pluginManager.installBrokerPlugin(pluginName, newPluginFunc)
}

func InitRegistry (ctx context.Context, name string, opts ...interface{}) (registry.Registry, error) {

	return pluginManager.initRegistry(ctx, name, opts...)
}

func InitBroker (ctx context.Context, name string, opts ...interface{}) (broker.Broker, error) {

	return pluginManager.initBroker(ctx, name, opts...)
}

func GetPlugins() *PluginManager {

	return pluginManager
}

