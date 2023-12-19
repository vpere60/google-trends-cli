package main

import (
	"fmt"
	"encoding/xml"
	"io"
	"net/http"
	"os"
)

type RSS struct {
	XMLName		xml.Name	`xml:"rss"`
	Channel		*Channel	`xml:"channel"`
}

type Channel struct {
	Title		string		`xml:"title"`
	ItemList	[]Item		`xml:"item"`
}

type Item struct {
	Title		string		`xml:"title"`
	Link		string		`xml:"link"`
	Traffic		string		`xml:"approx_traffic"`
	NewsItems	[]News		`xml:"news_item"`
}

type News struct {
	Headline		string	`xml:"news_item_title"`
	HeadlineLink	string	`xml:"news_item_url"`
}

func main() {
	var r RSS
	data := readGoogleTrends()
}

func getGoogleTrends() *http.Response {
	resp, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=US")

	if err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
}

func readGoogleTrends() []byte {
	getGoogleTrends
}


