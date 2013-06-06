package memory

import (
	"fmt"
	"github.com/pierrre/imageproxy"
	lru_impl "github.com/pierrre/imageproxy/cache/memory/lru"
)

type MemoryCache struct {
	lru *lru_impl.LRUCache
}

func New(capacity uint64) *MemoryCache {
	return &MemoryCache{
		lru: lru_impl.NewLRUCache(capacity),
	}
}

func (cache *MemoryCache) Get(key string) (image *imageproxy.Image, err error) {
	value, ok := cache.lru.Get(key)
	if !ok {
		err = fmt.Errorf("Image not found")
		return
	}
	item, ok := value.(*item)
	if !ok {
		err = fmt.Errorf("The cache value is not an image")
		return
	}
	image = item.image
	return
}

func (cache *MemoryCache) Set(key string, image *imageproxy.Image) (err error) {
	item := &item{
		image: image,
	}
	cache.lru.Set(key, item)
	return
}

type item struct {
	image *imageproxy.Image
}

func (item *item) Size() int {
	return len(item.image.Data)
}