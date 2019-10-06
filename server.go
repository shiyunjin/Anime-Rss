package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var (
	port      string
	rssCookie string
)

func main() {
	flag.StringVar(&port, "port", ":2333", "server port [default :2333]")
	flag.StringVar(&rssCookie, "rssCookie", "", "dmhy rss cookie [default: null] like :COOKIE:uid=****;rsspass=******")
	flag.Parse()
	Start(port)

}

func Start(port string) {
	http.HandleFunc("/", Rss)
	fmt.Println("Server Now Runing At", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

type rss struct {
	Channel Channel `xml:"channel"`
}

type Channel struct {
	Title         string  `xml:"title"`
	Link          string  `xml:"link"`
	Description   string  `xml:"description"`
	Language      string  `xml:"language"`
	LastBuildDate Date    `xml:"lastBuildDate"`
	Item          []*Item `xml:"item"`
	PubDate       Date    `xml:"pubDate"`
}

type ItemEnclosure struct {
	URL    string `xml:"url,attr"`
	Type   string `xml:"type,attr"`
	Length int64  `xml:"length,attr"`
}

type Item struct {
	Title       string          `xml:"title"`
	Link        string          `xml:"link"`
	PubDate     Date            `xml:"pubDate"`
	Description string          `xml:"description"`
	Enclosure   []ItemEnclosure `xml:"enclosure"`
	Author      string          `xml:"author"`
	GUID        string          `xml:"guid"`
	Category    string          `xml:"category"`
	CategoryUrl string          `xml:"category,attr"`
}

type Date string

func Rss(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://share.dmhy.org/topics/rss/sort_id/2/rss.xml" + rssCookie)
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
	var rssFilter []*Item
	for _, item := range rss.Channel.Item {
		if item.Category == "動畫" {
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
