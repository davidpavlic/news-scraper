package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
)

func inspectArticle(c *colly.Collector, articleUrl string) {
	c.OnHTML("h1.article__title", func(element *colly.HTMLElement) {
		fmt.Printf("Title: ")
		fmt.Println(element.Text)
	})

	c.OnHTML("div.article-content p", func(element *colly.HTMLElement) {
		if element.DOM.Parent().Is("div.article-content") && element.Attr("class") != "wp-caption-text" {
			fmt.Println(element.Text)
		}
	})

	err := c.Visit(articleUrl)
	if err != nil {
		log.Printf("failed to visit url: %v\n", err)
		return
	}
}