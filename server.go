package main

import (
	"AnimeRss/cache"
	"AnimeRss/model"
	"AnimeRss/path/favicon"
	"AnimeRss/path/rss"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	flag.StringVar(&cache.Port, "port", ":2333", "server port [default :2333]")
	flag.StringVar(&cache.RssCookie, "rssCookie", "", "dmhy rss cookie [default: null] like :COOKIE:uid=****;rsspass=******")
	flag.Parse()

	Start(cache.Port, cache.NewCache())
}

func Start(port string, cache model.CacheModel) {

	http.Handle("/category/", rss.NewRssHandler(cache))
	http.Handle("/favicon.ico", favicon.NewFaviconHandler(cache))

	fmt.Println("Server Now Runing At", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
