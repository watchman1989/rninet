package broker

import "context"

type Broker interface {
	Name() string
	Init(ctx context.Context, opts ...interface{}) error
	Connect() error
	Disconnect() error
	Publish(topic string, m *Message) error
	Subcribe(topic string)
}

type Message struct {
	Meta map[string]string
	Body string
}


