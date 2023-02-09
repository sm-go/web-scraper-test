package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gocolly/colly/v2"
	"github.com/smith-golang/Sites/ds"
	"github.com/smith-golang/Sites/model"
)

type BBCService struct {
	ds    *ds.DataSource
	colly *colly.Collector
}

func NewBBCService(ds *ds.DataSource) *BBCService {

	return &BBCService{
		ds: ds,
	}
}

type question struct {
	title, image, sub_title, paragraph string
}

func main() {

	c := colly.NewCollector()
	var bbclists []model.Post
	var details []model.PostDetail

	c.OnHTML("div.bbc-1hwgxf", func(e *colly.HTMLElement) {

		// fmt.Println(e.Text)
		q := question{}
		q.title = e.ChildText("h1#content")
		q.image = e.ChildAttr("img", "src")
		fmt.Println("one ??>>", q.title, "////")
		q.sub_title = e.ChildText("h2.bbc-1uafkdy")
		q.paragraph = e.ChildText("p.bbc-14tryw8")

		fmt.Println("two ... >>", q.image, "////")
		fmt.Println("three", q.sub_title, " sub title  >>>> ////")
		fmt.Println("four is parag", q.paragraph, " paragraph  >>>> ////")

	})

	fmt.Println(bbclists, "qssssssssssss >>>>>>>>>>>>")
	fmt.Println(details, "qssssssssssss >>>>>>>>>>>>")

	c.Visit("https://www.bbc.com/burmese/articles/cll3jeqe53zo")

	content, err := json.Marshal(details)
	if err != nil {
		fmt.Println(err.Error())
	}
	os.WriteFile("newData.json", content, 0644)
	fmt.Println("Total employees: ", len(details))

}
