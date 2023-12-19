package main

import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"os"
)

type Item struct {
	Title		string
	Link		string
	Traffic		string
	NewsItems	[]News
}

type News struct {
	Headline		string
	HeadlineLink	string
}

func main() {
	fmt.Println("Hello world!")
}

func main() {
	readGoogleTrends
}

func readGoogleTrends() {
	getGoogleTrends
}

func getGoogleTrends() {

}
