package elasticsearch

import (
	"context"
	"github.com/olivere/elastic"
)


var (
	elasticsearchComponent *ElasticsearchComponent = &ElasticsearchComponent{}
)


type ElasticsearchComponent struct {
	Client *elastic.Client
}


type ElasticsearchStarter struct {
	options *Options
}

func (e *ElasticsearchStarter) Init (ctx context.Context, opts ...interface{}) error {

	var (
		config []elastic.ClientOptionFunc
		client *elastic.Client
		err error
	)

	config = []elastic.ClientOptionFunc{
		elastic.SetURL(e.options.Urls...),
		elastic.SetSniff(e.options.Sniff),
	}

	client, err = elastic.NewClient(config...)
	if err != nil {
		return err
	}

	elasticsearchComponent.Client = client

	return nil
}
