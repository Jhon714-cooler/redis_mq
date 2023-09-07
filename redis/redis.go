package redis

import (
	"context"
	"errors"
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

	client.c = redis.NewClient(&redis.Options{
		Addr: client.Opts.address,
		Network: client.Opts.network,
		Password: client.Opts.password,
		PoolSize: client.Opts.poolSize,
		PoolTimeout: time.Duration(client.Opts.poolTimeout)*time.Second,
		DB: client.Opts.db,
		MaxIdleConns: client.Opts.maxIdleConns,
	})
	return &client
}
func (c *Client)GetConn()*redis.Conn{// 遗弃！！！
	return c.c.Conn()
}

func (c *Client)XADD(ctx context.Context,topic string,maxLen int,key,val string)(res string,err error) {
	if topic ==""{
		return "",errors.New("redis XADD topic can't be empty")
	}
	return c.c.XAdd(ctx,&redis.XAddArgs{
		Stream: topic,
		ID: "*",
		MaxLen: int64(maxLen),
		Values: map[string]interface{}{key: val},
	}).Result()
}