package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"regexp"
)

func getFixedArticle() string {
	var userInputURL string

	fmt.Print("Which Article of www.techcrunch.com would you like to visit? \n>")
	_, err := fmt.Scanln(&userInputURL)
	if err != nil {
		log.Fatalf("Error reading input: %v", err)
	}

	re := regexp.MustCompile(`^(https:\/\/techcrunch\.com|www\.techcrunch\.com).*`)
	if !re.MatchString(userInputURL) {
		log.Fatalf("Invalid URL: %s. URL must start with www.techcrunch.com", userInputURL)
	}

	return userInputURL
}

func getLatestArticle(c *colly.Collector) string {
	var articleUrl string
	var found bool

	c.OnHTML("a.post-block__title__link", func(element *colly.HTMLElement) {
		if found {
			return
		}
		found = true
		articleUrl = element.Attr("href")
	})

	err := c.Visit("https://techcrunch.com")
	if err != nil {
		log.Printf("failed to visit url: %v\n", err)
		return ""
	}

	return articleUrl
}