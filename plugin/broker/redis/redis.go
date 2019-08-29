package redis

import (
	"context"
	"github.com/watchman1989/rninet/plugin"
	"github.com/watchman1989/rninet/plugin/broker"
)

type redisBroker struct {

}

func init () {
	plugin.InstallBrokerPlugin("redis", NewRedisBroker)
}


func NewRedisBroker () broker.Broker {
	return &redisBroker{}
}

func (b *redisBroker) Name() string {
	return "redis"
}


func (b *redisBroker) Init (ctx context.Context, opts ...interface{}) error {


	return nil
}


func (b *redisBroker) Connect () error {


	return nil
}


func (b *redisBroker) Disconnect () error {


	return nil
}


func (b *redisBroker) Publish (topic string, m *broker.Message) {

}


func (b *redisBroker) Subcribe (topic string) {

}
