package rss

import (
	"AnimeRss/cache"
	"AnimeRss/model"
	"encoding/xml"
	"io/ioutil"
	"net/http"
)

type rss struct {
	Channel model.Channel `xml:"channel"`
}

func rssHandle(w http.ResponseWriter, r *http.Request, category string) {
	resp, err := http.Get("https://share.dmhy.org/topics/rss/sort_id/" + category + "/rss.xml" + cache.RssCookie)
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
	var rss rss
	if err := xml.Unmarshal(body, &rss); err != nil {
		w.WriteHeader(404)
		return
	}
	var rssFilter []*model.Item
	for _, item := range rss.Channel.Item {
		if item.Category.Domain == "http://share.dmhy.org/topics/list/sort_id/"+category {
			rssFilter = append(rssFilter, item)
		}
	}
	rss.Channel.Item = rssFilter
	rawBody, err := xml.Marshal(&rss)
	if err != nil {
		w.WriteHeader(404)
		return
	}
	w.Header().Set("content-type", "text/xml; charset=utf-8")
	w.Write(rawBody)
	return
}
