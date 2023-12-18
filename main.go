package main

import (
	"fmt"
	"github.com/caffix/cloudflare-roundtripper/cfrt"
	"github.com/gocolly/colly"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"time"
)

var allowedDomains = []string{"https://techcrunch.com", "www.techcrunch.com", "techcrunch.com"}

func main() {
	c := colly.NewCollector(colly.AllowedDomains(allowedDomains...))
	var userInput string
	var articleUrl string

	clearTerminal()

	fmt.Println("Which Option would you like to use?")
	fmt.Println("[1] - Paste in a fixed URL from an Article")
	fmt.Println("[2] - Get the Latest Article")

	for {
		_, err := fmt.Scanln(&userInput)

		if err != nil {
			fmt.Println("Error reading input: %v", err)
			continue
		}

		switch userInput {
		case "1":
			articleUrl = getFixedArticle()
			// Perform action for option 1
		case "2":
			articleUrl = getLatestArticle(c)
			// Perform action for option 2
		default:
			fmt.Print("Invalid option\n>")
			continue
		}
		break
	}
	setUp(articleUrl)
	inspectArticle(c, articleUrl)
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		return
	}
}

// Reduces Cloudflare Detection
func setUp(articleUrl string) {
	var err error
	// Set up your client however you need it. This is simply an example
	client := &http.Client{
		Timeout: 15 * time.Second,
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   15 * time.Second,
				KeepAlive: 15 * time.Second,
				DualStack: true,
			}).DialContext,
		},
	}
	// Set the client Transport to the RoundTripper that solves the Cloudflare anti-bot
	client.Transport, err = cfrt.New(client.Transport)
	if err != nil {
		return
	}

	req, err := http.NewRequest("GET", articleUrl, nil)
	if err != nil {
		return
	}

	_, err = client.Do(req)
	if err != nil {
		return
	}

}

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
