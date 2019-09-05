package redis

import (
	"context"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/watchman1989/rninet/plugin"
	"github.com/watchman1989/rninet/plugin/broker"
	"time"
)

type redisBroker struct {
	pool *redis.Pool
	options *Options
}

type publication struct {
	channel string
	message *broker.Message
}

type subscriber struct {
	conn *redis.PubSubConn
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

	b.options = &Options{}
	for _, opt := range opts {
		opt.(Option)(b.options)
	}

	return nil
}


func (b *redisBroker) Connect () error {

	b.pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", b.options.Host, b.options.Port))
			if err != nil {
				fmt.Println("NEWRDSPOOL_DIAL_ERROR: ", err)
				return nil, err
			}
			if len(b.options.Pass) > 0 {
				if _, err = c.Do("AUTH", b.options.Pass); err != nil {
					c.Close()
					fmt.Println("NEWRDSPOOL_AUTH_ERROR: ", err)
					return nil, err
				}
			}

			if _, err = c.Do("SELECT", b.options.Num); err != nil {
				c.Close()
				fmt.Println("NEWRDSPOOL_SELECT_ERROR: ", err)
				return nil, err
			}
			return c, nil
		},
		MaxIdle:     3000,
		MaxActive:   3000,
		IdleTimeout: 300 * time.Second,

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}


	return nil
}


func (b *redisBroker) Disconnect () error {

	if err := b.pool.Close(); err != nil {
		return nil
	}
	b.pool = nil

	return nil
}


func (b *redisBroker) Publish (topic string, m *broker.Message) error {

}


func (b *redisBroker) Subcribe (topic string) {

}
