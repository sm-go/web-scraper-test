package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		// colly.Debugger(&debug.LogDebugger{}),
		colly.MaxDepth(1), //no further link
	// colly.AllowedDomains("go-colly.org", "www.go-colly.org"),
	)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	// c.OnError(func(r *colly.Response, err error) {
	// 	log.Println("Something went wrong:", err)
	// })

	// c.OnHTML("a[href]", func(h *colly.HTMLElement) {
	// 	h.Request.Visit(h.Attr("href"))
	// })

	// c.OnHTML("tr td:nth-of-type(1)", func(h *colly.HTMLElement) {
	// 	fmt.Println("first colum table row:", h.Text)
	// })

	// c.OnXML("//h1", func(x *colly.XMLElement) {
	// 	fmt.Println(x.Text)
	// })

	// c.OnScraped(func(r *colly.Response) {
	// 	fmt.Println("finished", r.Request.URL)
	// })

	// c.WithTransport(&http.Transport{
	// 	Proxy: http.ProxyFromEnvironment,
	// 	DialContext: (&net.Dialer{
	// 		Timeout:   30 * time.Second,
	// 		DualStack: true,
	// 		KeepAlive: 30 * time.Second,
	// 	}).DialContext,
	// 	MaxIdleConns:          100,
	// 	IdleConnTimeout:       90 * time.Second,
	// 	TLSClientConfig:       &tls.Config{},
	// 	TLSHandshakeTimeout:   10 * time.Second,
	// 	ExpectContinueTimeout: 1 * time.Second,
	// })

	c.Visit("http://go-colly.org/")
}
