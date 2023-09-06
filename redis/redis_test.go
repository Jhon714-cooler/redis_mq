package redis

import (
	"context"
	"log"
	"testing"
)

const (
	addr     = "127.0.0.1:6379"
	password = ""
	db       = 2
)

func TestConnet(t *testing.T) {
	redisclient := NewClient(addr, password)
	conne := redisclient.GetConn()
	ctx := context.Background()
	if err := conne.Get(ctx, "liu"); err != nil {
		log.Fatal(err)
	}

}
