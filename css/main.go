package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gocolly/colly"
)

type CSSSelector struct {
	Selector    string `json:"selector"`
	Example     string `json:"example"`
	Description string `json:"description"`
}

func main() {
	allCSS := make([]CSSSelector, 0)
	c := colly.NewCollector()
	c.OnHTML("tbody > tr", func(h *colly.HTMLElement) {
		selector := h.ChildText("td > em") // (or) td > a
		if selector == "" {
			selector = h.ChildText("th")
		}
		if selector == "" {
			selector = h.ChildText("td > a")
		}
		if selector == "" {
			selector = h.ChildText("td > i > a")
		}
		example := h.ChildText("td.notranslate")
		if example == "" {
			example = h.ChildText("td")
		}
		if example == "" {
			example = h.ChildText("th")
		}
		desc := h.ChildText("td")
		if desc == "" {
			desc = h.ChildText("th")
		}
		css := CSSSelector{
			Selector:    selector,
			Example:     example,
			Description: desc,
		}
		allCSS = append(allCSS, css)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Visit("https://www.w3schools.com/cssref/css_selectors.php")
	writeJSON(allCSS)
}

func writeJSON(data []CSSSelector) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Println("Unable to create json file")
		return
	}
	_ = ioutil.WriteFile("css.json", file, 0644)
}
