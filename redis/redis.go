package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Client struct {
	Opts *ClientOptions
	c    *redis.Client
}

func NewClient(address, password string, opts ...options) *Client {
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
		Addr:         client.Opts.address,
		Network:      client.Opts.network,
		Password:     client.Opts.password,
		PoolSize:     client.Opts.poolSize,
		PoolTimeout:  time.Duration(client.Opts.poolTimeout) * time.Second,
		DB:           client.Opts.db,
		MaxIdleConns: client.Opts.maxIdleConns,
	})
	return &client
}
func (c *Client) GetConn() *redis.Conn { // 遗弃！！！
	return c.c.Conn()
}

// 投递消息
func (c *Client) XADD(ctx context.Context, topic string, maxLen int, key, val string) (res string, err error) {
	if topic == "" {
		return "", errors.New("redis XADD topic can't be empty")
	}
	return c.c.XAdd(ctx, &redis.XAddArgs{
		Stream: topic,
		ID:     "*",
		MaxLen: int64(maxLen),
		Values: map[string]interface{}{key: val},
	}).Result()
}

type MsgEntity struct {
	MsgID string
	value map[string]interface{}
}

// 消费消息
func (c *Client) XReadGroup(ctx context.Context, groupID, consumerID, topic string) (msg *MsgEntity, err error) {
	if groupID == "" || consumerID == "" || topic == "" {
		return nil, errors.New("redis XREADGROUP groupID/consumerID/topic can't be empty")
	}
	res, err := c.c.XReadGroup(ctx, &redis.XReadGroupArgs{
		Group:    groupID,
		Consumer: consumerID,
		Streams:  []string{topic, ">"},
	}).Result()
	if err != nil {
		return nil, err
	}
	for _, val := range res {
		for _, res := range val.Messages {
			msg.MsgID = res.ID
			msg.value = res.Values
		}
	}
	return msg, nil
}

func (c *Client) XAck(ctx context.Context, topic, groupID, msgID string) error {
	if topic == "" || groupID == "" || msgID == "" {
		return errors.New("redis XACK topic | group_id | msg_ id can't be empty")
	}
	reply, err := c.c.XAck(ctx, topic, groupID, msgID).Result()
	if err != nil {
		return err
	}
	if reply != 1 {
		return fmt.Errorf("invalid reply: %d", reply)
	}

	return nil
}
