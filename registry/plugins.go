package registry

import (
	"context"
	"errors"
	"sync"
)

var (
	pluginManager = &PluginManager{
		plugins: make(map[string]Registry),
	}
)

type PluginManager struct {
	plugins map[string]Registry
	lock sync.Mutex
}

func (p *PluginManager) registerPlugin (plugin Registry) (error) {

	var (
		err error
		ok bool
	)

	p.lock.Lock()
	defer  p.lock.Unlock()

	if _, ok = p.plugins[plugin.Name()]; ok {
		err = errors.New("PLUGIN_IS_EXISTS")
		return err
	}
	p.plugins[plugin.Name()] = plugin

	return nil
}

func (p *PluginManager) initRegister (ctx context.Context, name string, opts ...Option) (Registry, error) {

	var (
		err error
		ok bool
		plugin Registry
	)
	p.lock.Lock()
	defer p.lock.Unlock()

	if plugin, ok = p.plugins[name]; !ok {
		err = errors.New("PLUGIN_IS_NOT_EXISTS")
		return nil, err
	}

	err = plugin.Init(ctx, opts...)

	return plugin, nil
}


func RegisterPlugin(registry Registry) error {

	return pluginManager.registerPlugin(registry)
}

func InitRegistry (ctx context.Context, name string, opts ...Option) (Registry, error) {

	return pluginManager.initRegister(ctx, name, opts...)
}

func GetPlugins() *PluginManager {

	return pluginManager
}
