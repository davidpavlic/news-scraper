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
	"time"
)

func main() {

	clearTerminal()

	setUp()

	c := colly.NewCollector(colly.AllowedDomains("https://techcrunch.com", "www.techcrunch.com", "techcrunch.com"))

	c.OnHTML("h1.article__title", func(element *colly.HTMLElement) {
		fmt.Printf("Title: ")
		fmt.Println(element.Text)
	})

	c.OnHTML("div.article-content p", func(element *colly.HTMLElement) {
		if element.DOM.Parent().Is("div.article-content") && element.Attr("class") != "wp-caption-text" {
			fmt.Println(element.Text)
		}
	})

	err := c.Visit("https://techcrunch.com/2023/12/12/comun-local-banking-latino-immigrants-fintech/")
	if err != nil {
		log.Printf("failed to visit url: %v\n", err)
		return
	}
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
func setUp() {
	var err error

	// Setup your client however you need it. This is simply an example
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

	req, err := http.NewRequest("GET", "https://techcrunch.com/2023/12/12/comun-local-banking-latino-immigrants-fintech/", nil)
	if err != nil {
		return
	}

	_, err = client.Do(req)
	if err != nil {
		return
	}

}
