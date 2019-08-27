package nsq

import (
	"context"
	"fmt"
	"github.com/nsqio/go-nsq"
	"time"
)


var (
	nsqProducerComponent *NsqProducerComponent = &NsqProducerComponent{}
	nsqConsumerComponent *NsqConsumerComponent = &NsqConsumerComponent{}
)

type NsqProducerComponent struct {
	Producer *nsq.Producer
}

type NsqConsumerComponent struct {
	Consumer *nsq.Consumer
}

type NsqProducerStarter struct {
	options *Options
}

type NsqConsumerStarter struct {
	options *Options
}

func (n *NsqProducerStarter) Init (ctx context.Context, opts ...interface{}) error {

	var (
		config *nsq.Config
		producer *nsq.Producer
		err error
	)

	n.options = &Options{}
	for _, opt := range opts {
		opt.(Option)(n.options)
	}

	config = nsq.NewConfig()

	producer, err = nsq.NewProducer(n.options.Addr, config)
	if err != nil {
		return err
	}

	nsqProducerComponent.Producer = producer

	return nil
}


type Handler struct {}

func (h *Handler) HandleMessage (msg *nsq.Message) error {

	fmt.Println(msg.NSQDAddress, string(msg.Body))

	return nil
}


func (n *NsqConsumerStarter) Init (ctx context.Context, handler nsq.Handler, opts ...interface{}) error {

	var (
		config *nsq.Config
		consumer *nsq.Consumer
		err error
	)

	n.options = &Options{}
	for _, opt := range opts {
		opt.(Option)(n.options)
	}

	config = nsq.NewConfig()
	config.LookupdPollInterval = time.Duration(n.options.Interval) * time.Second

	consumer, err = nsq.NewConsumer(n.options.Topic, n.options.Channel, config)
	if err != nil {
		return err
	}

	consumer.AddHandler(handler)

	if err = consumer.ConnectToNSQLookupd(n.options.Addr); err != nil {
		return err
	}

	return nil
}
