package rss

import "net/http"

func (t *RssHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/category/2.xml":
		rssHandle(w, r, "2")
	default:
		w.WriteHeader(404)
	}
}
