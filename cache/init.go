package cache

import (
	"AnimeRss/model"
	"sync"
	"time"
)

type Cache struct {
	list sync.Map
}

func NewCache() model.CacheModel {
	return &Cache{}
}

func (c *Cache) Set(path string, data []byte) {
	c.list.Store(path, model.CacheItem{
		Data:       data,
		CreateTime: time.Now().Unix(),
	})
}

func (c *Cache) Get(path string) []byte {
	data, ok := c.list.Load(path)
	if ok {
		switch data.(type) {
		case model.CacheItem:
			return data.(model.CacheItem).Data
		}
	}
	return nil
}
