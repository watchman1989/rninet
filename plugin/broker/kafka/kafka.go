package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/watchman1989/rninet/plugin"
	"github.com/watchman1989/rninet/plugin/broker"
)

type kafkaBroker struct {
	options *Options
	client sarama.Client
}

func init () {
	plugin.InstallBrokerPlugin("kafka", NewKafkaBroker)
}


func NewKafkaBroker () broker.Broker {
	return &kafkaBroker{}
}

func (b *kafkaBroker) Name() string {
	return "kafka"
}


func (b *kafkaBroker) Init (ctx context.Context, opts ...interface{}) error {

	b.options = &Options{}
	for _, opt := range opts {
		opt.(Option)(b.options)
	}

	if len(b.options.Addrs) == 0 {
		b.options.Addrs = []string{"127.0.0.1:9092"}
	}

	return nil
}


func (b *kafkaBroker) Connect () error {


	return nil
}


func (b *kafkaBroker) Disconnect () error {


	return nil
}


func (b *kafkaBroker) Publish (topic string, m *broker.Message) {

}


func (b *kafkaBroker) Subcribe (topic string) {

}

