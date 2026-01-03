package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

type news struct {
	name string
	link string
}

func main() {
	var s []string
	var results []news

	for i := 1; i < 10; i++ {
		s = append(s, fmt.Sprintf("https://news.ycombinator.com/?p=%d", i))
	}

	c := colly.NewCollector()

	c.OnHTML(".titleline > a", func(e *colly.HTMLElement) {

		title := e.Text
		link := e.Attr("href")
		
		results = append(results, news{
			name: title,
			link: link,
		})
		
		fmt.Printf("Title: %s\nLink: %s\n\n", title, link)
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Printf("Visit: %s\n", r.URL)
	})

	for _, value := range s {
		c.Visit(value)
	}

	fmt.Printf("%d\n", len(results))
}