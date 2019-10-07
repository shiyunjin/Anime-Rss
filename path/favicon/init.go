package favicon

import "AnimeRss/model"

type FaviconHandler struct {
	cache model.CacheModel
}

func NewFaviconHandler(cacheModel model.CacheModel) *FaviconHandler {
	return &FaviconHandler{
		cache: cacheModel,
	}
}
