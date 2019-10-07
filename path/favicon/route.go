package favicon

import "net/http"

func (t *FaviconHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	faviconHandle(w, r, t.cache)
}
