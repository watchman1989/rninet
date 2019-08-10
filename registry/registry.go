package registry

import "context"


type Registry interface {
	Name() string
	Init(ctx context.Context, opts ...Option) (error)
	Register (ctx context.Context, service *Service) (error)
	Deregister (ctx context.Context, service *Service) (error)
	QueryService (ctx context.Context, name string) (map[string]*Service, error)
	SyncService (ctx context.Context, name string) (chan map[string]*Service)
}
