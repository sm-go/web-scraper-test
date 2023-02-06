package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gocolly/colly"
)

type Fact struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
}

func main() {
	allFacts := make([]Fact, 0)
	c := colly.NewCollector(
		colly.AllowedDomains(
			"factretriever.com",
			"www.factretriever.com",
		),
	)
	c.OnHTML(".factsList li", func(h *colly.HTMLElement) {
		factId, err := strconv.Atoi(h.Attr("id"))
		if err != nil {
			log.Println("Could not get id")
		}
		factDesc := h.Text
		fact := Fact{
			Id:          factId,
			Description: factDesc,
		}
		allFacts = append(allFacts, fact)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit("https://www.factretriever.com/rhino-facts")
	writeJSON(allFacts)
}

func writeJSON(data []Fact) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}
	_ = ioutil.WriteFile("rhinofacts.json", file, 0644)
}
