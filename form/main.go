package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	// create a new collector
	c := colly.NewCollector()

	c.OnHTML("html", func(h *colly.HTMLElement) {
		// fmt.Println(h.Text)
		h.Request.PostMultipart("https://luckyadmin.huyi.buzz/api/admin/login", generateFormData())
		fmt.Println(generateFormData())
	})

	// start scraping
	c.Visit("https://luckyadmin.huyi.buzz")
}

func generateFormData() map[string][]byte {
	return map[string][]byte{
		"username": []byte("admin"),
		"password": []byte("admin"),
		"otp":      []byte("987654321"),
	}
}
