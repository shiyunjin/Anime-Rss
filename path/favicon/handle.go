package favicon

import (
	"AnimeRss/model"
	"io/ioutil"
	"net/http"
)

func faviconHandle(w http.ResponseWriter, r *http.Request, cache model.CacheModel) {
	w.Header().Set("content-type", "image/x-icon")
	cacheData := cache.Get(r.URL.Path)
	if cacheData == nil {
		resp, err := http.Get("https://share.dmhy.org/favicon.ico")
		if err != nil {
			w.WriteHeader(404)
			return
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			w.WriteHeader(404)
			return
		}

		w.Write(body)
		cache.Set(r.URL.Path, body)
		return
	}
	w.Write(cacheData)
	return
}
