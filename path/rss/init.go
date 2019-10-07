package rss

import "AnimeRss/model"

type RssHandler struct {
	cache model.CacheModel
}

func NewRssHandler(cacheModel model.CacheModel) *RssHandler {
	return &RssHandler{
		cache: cacheModel,
	}
}
