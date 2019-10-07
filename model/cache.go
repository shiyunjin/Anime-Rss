package model

type CacheModel interface {
	Set(path string, data []byte)
	Get(path string) []byte
}

type CacheItem struct {
	Data       []byte
	CreateTime int64
}
