package redis

import (
	"context"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)


var (
	redisConnects []*RedisComponent = make([]*RedisComponent, 16)
)


type RedisComponent struct {
	pool *redis.Pool
	debugFlag bool
}


type RedisStarter struct {
	options *Options
}


func (s *RedisStarter) Init (ctx context.Context, opts ...interface{}) error {

	s.options = &Options{}
	for _, opt := range opts {
		opt.(Option)(s.options)
	}

	for dbNum := 0; dbNum < 16; dbNum++ {
		pool := redis.Pool{
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", fmt.Sprintf("%s:%d", s.options.Host, s.options.Port))
				if err != nil {
					fmt.Println("NEWRDSPOOL_DIAL_ERROR: ", err)
					return nil, err
				}
				if len(s.options.Pass) > 0 {
					if _, err = c.Do("AUTH", s.options.Pass); err != nil {
						c.Close()
						fmt.Println("NEWRDSPOOL_AUTH_ERROR: ", err)
						return nil, err
					}
				}

				if _, err = c.Do("SELECT", dbNum); err != nil {
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

		rc := RedisComponent{}
		rc.pool = &pool
		rc.debugFlag = false

		redisConnects = append(redisConnects, &rc)
	}

	return nil
}


func (s *RedisStarter) Stop () error {

	return nil
}


func (r *RedisComponent) Do (commandName string, args ...interface{}) (reply interface{}, err error) {
	c := r.pool.Get()
	defer c.Close()

	tBegin := time.Now().UnixNano()
	reply, err = c.Do(commandName, args...)
	if err != nil {
		e := c.Err()
		if e != nil {
			fmt.Println("RDS_DO_ERROR", err, e)
		}
	}
	tEnd := time.Now().UnixNano()
	if r.debugFlag {
		fmt.Printf("[REDIS_INFO] [%dus] cmd=%s, err=%s, args=%v, reply=%v\n", (tEnd - tBegin)/1000, commandName, err, args, reply)
	}

	return reply, err
}


func GetRedisComponent(dbNum int) *RedisComponent {

	return redisConnects[dbNum]
}

