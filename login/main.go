package main

import (
	"log"

	"github.com/gocolly/colly"
)

func main() {
	// create a new collector
	c := colly.NewCollector()

	// authenticate
	err := c.Post("https://luckyadmin.huyi.buzz/api/admin/login", map[string]string{"username": "admin", "password": "admin", "otp": "987654321"})
	if err != nil {
		log.Fatal(err)
	}

	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.Ctx)
	})

	// start scraping
	c.Visit("https://luckyadmin.huyi.buzz")
}
