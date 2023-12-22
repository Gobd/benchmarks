package client

import (
	"github.com/hashicorp/golang-lru/arc/v2"
	lru "github.com/hashicorp/golang-lru/v2"
)

type LRU[K comparable, V any] struct {
	client *lru.Cache[K, V]
}

func (c *LRU[K, V]) Init(cap int) {
	client, err := lru.New[K, V](cap)
	if err != nil {
		panic(err)
	}
	c.client = client
}

func (c *LRU[K, V]) Get(key K) (V, bool) {
	return c.client.Get(key)
}

func (c *LRU[K, V]) Set(key K, value V) {
	c.client.Add(key, value)
}

func (c *LRU[K, V]) Name() string {
	return "LRU"
}

func (c *LRU[K, V]) Close() {
	c.client = nil
}

type ARC[K comparable, V any] struct {
	client *arc.ARCCache[K, V]
}

func (c *ARC[K, V]) Init(cap int) {
	client, err := arc.NewARC[K, V](cap)
	if err != nil {
		panic(err)
	}
	c.client = client
}

func (c *ARC[K, V]) Get(key K) (V, bool) {
	return c.client.Get(key)
}

func (c *ARC[K, V]) Set(key K, value V) {
	c.client.Add(key, value)
}

func (c *ARC[K, V]) Name() string {
	return "ARC"
}

func (c *ARC[K, V]) Close() {
	c.client = nil
}