package main

import (
	"github.com/gocolly/colly"
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
	name, date, location, locLink, all_data string
}

func scrape(url string) {
	c := colly.NewCollector()
	c.OnHTML("div[class]", func(e *colly.HTMLElement){
		cls := e.Attr("class")
		if cls == "col-md-24 details" {
			item := evt{}
			item.name = e.ChildText("div > h2")
			item.all_data = e.ChildText("div > p")
			adARR := strings.Split(item.all_data, "|")
			if len(adARR) == 3{
				item.date = adARR[0]
				item.location = adARR[1]
				return
			} else {
				return
			}
		}
	})
	fmt.Println(item.name)
	c.Visit(url)
}