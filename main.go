package main

import (
	"net/http"
    "encoding/xml"
    "fmt"
	"io/ioutil"
	"log"	
	"reflect"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Channel struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Language    string `xml:"language"`
		Item        []struct {
			Title       string `xml:"title"`
			Description string `xml:"description"`
			Thumbnail   struct {
				Text   string `xml:",chardata"`
				Width  string `xml:"width,attr"`
				Height string `xml:"height,attr"`
				URL    string `xml:"url,attr"`
			} `xml:"thumbnail"`
			Link    string `xml:"link"`
			PubDate string `xml:"pubDate"`
			NlTitle string `xml:"nl-title"`
		} `xml:"item"`
	} `xml:"channel"`
} 

func main() {

	req, err := http.NewRequest("GET", "http://www.nzz.ch/recent.rss", nil)
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}


	type Loc string

	var data = []byte(string(bytes))
	

		var rss Rss
		xml.Unmarshal(data, &rss)

		v := reflect.ValueOf(rss.Channel)
		for i := 0; i< v.NumField(); i++ {
			fmt.Println(rss.Channel.Item[i].Title)
		}
			
}
