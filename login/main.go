package main

import (
	"github.com/gocolly/colly"
	"log"
)

func main() {
	// create a new collector
	c := colly.NewCollector()

	// authenticate
	// 这真能用？本地的 RabbitMQ 都登不上啊
	err := c.Post("http://localhost:15672/", map[string]string{"username": "guest", "password": "guest"})
	if err != nil {
		log.Fatal(err)
	}

	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
	})

	// start scraping
	c.Visit("https://example.com/")
}
