package main

import (
	"bytes"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
	"log"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(colly.AllowURLRevisit())

	// Rotate two socks5 proxies
	rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:1337", "socks5://127.0.0.1:1338")
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	// Print the response
	c.OnResponse(func(r *colly.Response) {
		log.Printf("%s\n", bytes.Replace(r.Body, []byte("\n"), nil, -1))
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Fetch httpbin.org/ip five times
	for i := 0; i < 5; i++ {
		c.Visit("https://httpbin.org/ip")
	}
}
