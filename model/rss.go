package model

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

type ItemCategory struct {
	Domain string `xml:"domain,attr"`
	Value  string `xml:",chardata"`
}

type Item struct {
	Title       string          `xml:"title"`
	Link        string          `xml:"link"`
	PubDate     Date            `xml:"pubDate"`
	Description string          `xml:"description"`
	Enclosure   []ItemEnclosure `xml:"enclosure"`
	Author      string          `xml:"author"`
	GUID        string          `xml:"guid"`
	Category    ItemCategory    `xml:"category"`
}

type Date string
