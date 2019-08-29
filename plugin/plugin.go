package plugin

import (
	"context"
	"errors"
	"github.com/watchman1989/rninet/plugin/broker"

	"github.com/watchman1989/rninet/plugin/registry"
	"sync"
)

var (
	pluginManager = &PluginManager{
		registryPlugins: make(map[string]registry.Registry),
		brokerPlugins: make(map[string]broker.Broker),
	}
)

type PluginManager struct {
	registryPlugins map[string]registry.Registry
	brokerPlugins map[string]broker.Broker
	lock sync.Mutex
}

func (p *PluginManager) installPlugin (pluginType string, plugin interface{}) (error) {

	var (
		err error
		ok bool
	)

	p.lock.Lock()
	defer  p.lock.Unlock()

	switch pluginType {
	case "registry":
		if _, ok = p.registryPlugins[plugin.(registry.Registry).Name()]; ok {
			err = errors.New("PLUGIN_IS_EXISTS")
			return err
		}
		p.registryPlugins[plugin.(registry.Registry).Name()] = plugin.(registry.Registry)
	case "broker":
		if _, ok = p.brokerPlugins[plugin.(broker.Broker).Name()]; ok {
			err = errors.New("PLUGIN_IS_EXISTS")
			return err
		}
		p.brokerPlugins[plugin.(broker.Broker).Name()] = plugin.(broker.Broker)

	}

	return nil
}


func (p *PluginManager) initRegistry (ctx context.Context, name string, opts ...interface{}) (registry.Registry, error) {
	var (
		err    error
		ok     bool
		plugin registry.Registry
	)
	p.lock.Lock()
	defer p.lock.Unlock()

	if plugin, ok = p.registryPlugins[name]; !ok {
		err = errors.New("PLUGIN_IS_NOT_EXISTS")
		return nil, err
	}

	err = plugin.Init(ctx, opts...)

	return plugin, nil
}


func (p *PluginManager) initBroker (ctx context.Context, name string, opts ...interface{}) (broker.Broker, error) {
	var (
		err    error
		ok     bool
		plugin broker.Broker
	)
	p.lock.Lock()
	defer p.lock.Unlock()

	if plugin, ok = p.brokerPlugins[name]; !ok {
		err = errors.New("PLUGIN_IS_NOT_EXISTS")
		return nil, err
	}

	err = plugin.Init(ctx, opts...)

	return plugin, nil
}


func InstallPlugin(pluginType string, plugin interface{}) error {

	return pluginManager.installPlugin(pluginType, plugin)
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

