package client

import (
	"strconv"

	lru "github.com/cloudxaas/gocache/lru/bytes"
)

type Cloudxaas struct {
	client *lru.ShardedCache
}

func (c *Cloudxaas) Init(capacity int) {
	client := lru.NewShardedCache(128, int64(capacity), 1)
	c.client = client
}

func (c *Cloudxaas) Get(key uint64) (uint64, bool) {
	b, ok := c.client.Get([]byte(strconv.FormatUint(key, 10)))
	if !ok {
		return 0, ok
	}
	val, err := strconv.ParseUint(string(b), 10, 64)
	if err != nil {
		panic(err)
	}
	return val, ok
}

func (c *Cloudxaas) Set(key, value uint64) {
	c.client.Put([]byte(strconv.FormatUint(key, 10)), []byte(strconv.FormatUint(value, 10)))
}

func (c *Cloudxaas) Name() string {
	return "cloudxaas"
}

func (c *Cloudxaas) Close() {
	c.client = nil
}

type CloudxaasString struct {
	client *lru.ShardedCache
}

func (c *CloudxaasString) Init(capacity int) {
	client := lru.NewShardedCache(128, int64(capacity), 1)
	c.client = client
}

func (c *CloudxaasString) Get(key string) (string, bool) {
	b, ok := c.client.Get([]byte(key))
	return string(b), ok
}

func (c *CloudxaasString) Set(key, value string) {
	c.client.Put([]byte(key), []byte(value))
}

func (c *CloudxaasString) Name() string {
	return "cloudxaas"
}

func (c *CloudxaasString) Close() {
	c.client = nil
}
