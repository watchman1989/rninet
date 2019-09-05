package mongo

import (
	"context"
	"gopkg.in/mgo.v2"
)


var (
	mongoComponent *MongoComponent = &MongoComponent{}
)

func NewmongoComponent () *MongoComponent {
	return mongoComponent
}


type MongoComponent struct {
	Name string
	Session *mgo.Session
}


type MongoStarter struct {
	options *Options
}


func (e *MongoStarter) Init (ctx context.Context, opts ...interface{}) error {

	var (
		err error
		session *mgo.Session
	)

	e.options = &Options{}
	for _, opt := range opts {
		opt.(Option)(e.options)
	}

	session, err = mgo.Dial(e.options.Url)
	if err != nil {
		return err
	}

	mongoComponent.Session = session

	return nil
}


func (e *MongoStarter) Stop () error {
	return nil
}

