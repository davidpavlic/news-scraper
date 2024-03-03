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
			fmt.Print("Invalid option\n")
			displayOptions() // Remind the user of the options if they input an invalid option
			continue
		}
		break
	}
	setUp(articleUrl)

	articleContent, err := inspectArticle(c, articleUrl)
	if err != nil {
		fmt.Printf("Error inspecting article: %v\n", err)
		return
	} else {
		fmt.Printf(articleContent + "\n\n\n\n\n\n\n\n")
	}

	// Now summarize the content using ChatGPT
	summary, err := SummarizeContent(articleContent)
	if err != nil {
		fmt.Printf("Error summarizing content: %v\n", err)
	} else {
		fmt.Println("Summary:", summary)
	}
}