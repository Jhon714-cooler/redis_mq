package mq

import (
	"context"

	"github.com/redismq/redis"
)

type Producer struct {
	opts	*ProducerOptions
	client *redis.Client
}

func NewProducer(client redis.Client,opts ProducerOption)( *Producer){
	p := Producer{
		client:&client,
		opts: &ProducerOptions{},
	}
	opts(p.opts)
	return &p
}
//发送消息
func (p *Producer) SendMsg(ctx context.Context, topic, key, val string) (string, error) {
	return p.client.XADD(ctx, topic, p.opts.MsgQuenueLen, key, val)
}
