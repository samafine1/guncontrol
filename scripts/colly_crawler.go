package main

import (
	"github.com/gocolly/colly"
	//"fmt"
	"strconv"
	"time"
	"strings"
	"fmt"
)

func main() {
	URL := "http://www.nraam.org/events"
	queue := detURLarr(URL)
	
	for _, url := range queue{
		fmt.Println(url)
		scrape(url)
	}
}
// vvv determines the urls that are relevant to this project
func detURLarr(url string) []string {
	baseUrl := "http://www.nraam.org"
	c := colly.NewCollector()
	arr := []string{}
	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		evtUrls := e.Attr("href")
		if strings.HasPrefix(evtUrls, "/events/"+strconv.Itoa(time.Now().Year())) {
			arr = append(arr, baseUrl+evtUrls)
		}
	})
	c.Visit(url)
	return arr
}

type evt struct {
	name, date, location, all_data string
}
func scrape(url string){
	c := colly.NewCollector()
	info := []string{}
	c.OnHTML("div[class]", func(e *colly.HTMLElement) {
		cls := e.Attr("class")
		if cls == "col-md-24 details" {
			info = append(info, e.Text)
			fmt.Println(e.Text)
		}
	})
	c.Visit(url)
}