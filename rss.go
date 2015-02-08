package main

import (
	"encoding/xml"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Items Items `xml:"channel"`
}

type Items struct {
	XMLName xml.Name `xml:"channel"`
	ItemList []Item `xml:"item"`
}

type Item struct {
	Title string `xml:"title"`
	Category string `xml:"category"`
	Link string `xml:"link"`
}

func ReadRSS( in []byte ) (RSS,error) {
	var rss RSS
	err := xml.Unmarshal( in, &rss )

	if err != nil {
		return rss, err
	}
	return rss, nil
}