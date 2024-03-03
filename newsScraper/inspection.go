package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

func inspectArticle(c *colly.Collector, articleUrl string) (string, error){
	var articleContent strings.Builder

	c.OnHTML("h1.article__title", func(element *colly.HTMLElement) {
		articleContent.WriteString("Title: " + element.Text + "\n")
	})

	c.OnHTML("div.article-content p", func(element *colly.HTMLElement) {
		if element.DOM.Parent().Is("div.article-content") && element.Attr("class") != "wp-caption-text" {
			articleContent.WriteString(element.Text)
		}
	})

	err := c.Visit(articleUrl)
	if err != nil {
		log.Printf("failed to visit url: %v\n", err)
		return "", fmt.Errorf("failed to visit url: %v", err)
	}
	return articleContent.String(), nil
}