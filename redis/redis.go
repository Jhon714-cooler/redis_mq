package redis

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	Opts	*ClientOptions
	c *redis.Client
}

func NewClient(address, password string,opts ...options) *Client {
	client := Client{
		Opts: &ClientOptions{
			network:  "tcp",
			address:  address,
			password: password,
		},
	}
	for _, opt := range opts {
		opt(client.Opts)
	}
	checkParm(client.Opts)
	o := redis.Options{
		Addr: client.Opts.address,
		Network: client.Opts.network,
		Password: client.Opts.password,
		PoolSize: client.Opts.poolSize,
		PoolTimeout: time.Duration(client.Opts.poolTimeout)*time.Second,
		DB: client.Opts.db,
		MaxIdleConns: client.Opts.maxIdleConns,
	}
	client.c = redis.NewClient(&o)
	return &client
}
func (c *Client)GetConn()*redis.Conn{
	return c.c.Conn()
}
