package main

import (
	"github.com/gocolly/colly"
	"strconv"
	"time"
	"strings"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func main() {
	URL := "http://www.nraam.org/events"
	queue := detURLarr(URL)
	events := ""
	for _, url := range queue{
		fmt.Println(url)
		event := scrape(url)
		events = events + event
	}
	fmt.Println(events)
	defer writeEvt(events)
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

func (e * evt)handleEvt()string{
	return "<p>Event Name: "+e.name+"</p> <p>Date: "+e.date+"</p> <p>location: "+e.location+"</p>"
}

func scrape(url string) string{
	c := colly.NewCollector()
	item := evt{}
	c.OnHTML("div[class]", func(e *colly.HTMLElement){
		cls := e.Attr("class")
		if cls == "col-md-24 details" {
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
	//fmt.Println(item.name)
	c.Visit(url)
	return item.handleEvt()
}

func writeEvt(s string) error{
	path, _ := filepath.Abs("guncontrol/html")
	file := "events.html"
	f, _ := ioutil.ReadFile(filepath.Join(path, file))
	 file_content := string(f)
	 f_arr := strings.Split(file_content, "<!--put shit here -->")
	 new_cnt := f_arr[0]+"<!--put shit here -->"+s+"<!--put shit here -->"+f_arr[len(f_arr)-1]
	 fmt.Println("DONE")
	 return ioutil.WriteFile(filepath.Join(path, file), []byte(new_cnt), 0755)
}