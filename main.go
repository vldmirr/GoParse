package main

import(
	// "net/http"
	"fmt"
	"github.com/gocolly/colly/v2"
	// "time"
)

type news struct{
	name string
	link string
}



func main(){
	var s []string
	// var result []news

	for i := 1; i < 10; i++ {
		s=append(s,fmt.Sprintf("https://news.ycombinator.com/?p=%d", i))
	}
	
	c := colly.NewCollector()

	for _,value:=range s{
		c.OnHTML(".title > span > a", func(e *colly.HTMLElement) {
			e.Request.Visit(e.Attr("href"))
		})	

		c.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting",r.URL)
		})

		c.Visit(value)
	}
	

}