package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(colly.AllowedDomains("www.mydomain.com")) //allowed domain

	c.OnHTML("a[href]", func(h *colly.HTMLElement) { //scrapes [href] element from site
		fmt.Println(h.Text)
	})

	c.Visit("www.mydomain.com/some/fany/url") //specific url
}
