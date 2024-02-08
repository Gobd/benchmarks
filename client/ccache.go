package client

import (
	"github.com/karlseguin/ccache/v3"
	"runtime"
	"time"
)

type Ccache[V any] struct {
	client *ccache.Cache[V]
}

func (c *Ccache[V]) Init(cap int) {
	client := ccache.New(
		ccache.Configure[V]().
			MaxSize(int64(cap)).
			Buckets(uint32(16 * runtime.GOMAXPROCS(0))),
	)
	c.client = client
}

func (c *Ccache[V]) Name() string {
	return "Ccache"
}

func (c *Ccache[V]) Get(key string) (V, bool) {
	item := c.client.Get(key)
	if item == nil || item.Expired() {
		var value V
		return value, false
	}

	return item.Value(), true
}

func (c *Ccache[V]) Set(key string, value V) {
	c.client.Set(key, value, time.Hour)
}

func (c *Ccache[V]) Close() {
	c.client.Clear()
	c.client.Stop()
}
