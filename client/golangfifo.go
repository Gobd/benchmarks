package client

import fifo "github.com/scalalang2/golang-fifo/v2"

type FIFO[K comparable, V any] struct {
	client fifo.Cache[K, V]
}

func (c *FIFO[K, V]) Init(cap int) {
	client := fifo.NewS3FIFO[K, V](cap)
	c.client = client
}

func (c *FIFO[K, V]) Get(key K) (V, bool) {
	return c.client.Get(key)
}

func (c *FIFO[K, V]) Set(key K, value V) {
	c.client.Set(key, value)
}

func (c *FIFO[K, V]) Name() string {
	return "S3FIFO"
}

func (c *FIFO[K, V]) Close() {
	c.client.Clean()
	c.client = nil
}
