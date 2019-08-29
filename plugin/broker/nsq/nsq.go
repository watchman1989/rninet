package nsq

type nsqBroker struct {
	options *Options
}



/*
type NsqProducerComponent struct {
	Producer *nsq.Producer
}

type NsqConsumerComponent struct {
	Consumer *nsq.Consumer
	ReceiveCount int64
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

	producer, err = nsq.NewProducer(Addr, config)
	if err != nil {
		return err
	}



	return nil
}


func (n *NsqProducerComponent) Publish (topic string, message []byte) error {

	err := n.Producer.Publish(topic, message)
	if err != nil {
		return nil
	}

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
	config.LookupdPollInterval = time.Duration(Interval) * time.Second

	consumer, err = nsq.NewConsumer(Topic, Channel, config)
	if err != nil {
		return err
	}

	consumer.AddHandler(handler)

	if err = consumer.ConnectToNSQLookupd(Addr); err != nil {
		return err
	}

	return nil
}
*/