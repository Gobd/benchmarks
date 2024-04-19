package client

import (
	"github.com/phuslu/lru"
)

type PHULRU[K comparable, V any] struct {
	client *lru.LRUCache[K, V]
}

func (c *PHULRU[K, V]) Init(capacity int) {
	client := lru.NewLRUCache[K, V](capacity)
	c.client = client
}

func (c *PHULRU[K, V]) Get(key K) (V, bool) {
	return c.client.Get(key)
}

func (c *PHULRU[K, V]) Set(key K, value V) {
	c.client.Set(key, value)
}

func (c *PHULRU[K, V]) Name() string {
	return "phuslu"
}

func (c *PHULRU[K, V]) Close() {
	c.client = nil
}
