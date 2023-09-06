package redis

const (
	// 默认连接池超过 10 s 释放连接
	DefaultIdleTimeoutSeconds = 10
	// 默认最大空闲连接数
	DefaultMaxIdle = 20
)

type ClientOptions struct {
	poolTimeout	int
	poolSize     int
	db           int
	maxIdleConns int
	// 必填参数
	network  string
	address  string
	password string
}
type options func(c *ClientOptions)

func WithPoolTimeout(poolTimeout int) options {
	return func(c *ClientOptions) {
		c.poolTimeout = poolTimeout
	}
}
func WithPoolsize(poolSize int) options {
	return func(c *ClientOptions) {
		c.poolSize = poolSize
	}
}

func WithDb(dbnum int) options {
	return func(c *ClientOptions) {
		c.db = dbnum
	}
}

func WithMaxIdleConns(MaxIdleConns int) options {
	return func (c *ClientOptions)  {
		c.maxIdleConns = MaxIdleConns
	}
}

func checkParm(c *ClientOptions){
	if c.maxIdleConns < 0 {
		c.maxIdleConns = DefaultMaxIdle
	}
	if c.poolTimeout < 0 {
		c.poolTimeout= DefaultIdleTimeoutSeconds
	}
	if c.poolSize < 0 {
		c.poolSize = 20
	}
	if c.db < 0{
		c.db = 0
	}
}