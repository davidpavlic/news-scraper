package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

var allowedDomains = []string{"https://techcrunch.com", "www.techcrunch.com", "techcrunch.com"}

func main() {
	c := colly.NewCollector(colly.AllowedDomains(allowedDomains...))

	clearTerminal()
	displayOptions()

	var articleUrl string

	for {
		userInput, err := readUserInput()

		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}

		switch userInput {
		case "1":
			articleUrl = getFixedArticle()
		case "2":
			articleUrl = getLatestArticle(c)
		default:
			fmt.Print("Invalid option\n>")
			displayOptions() // Remind the user of the options if they input an invalid option
			continue
		}
		break
	}
	setUp(articleUrl)
	inspectArticle(c, articleUrl)
}