package main

import (
	"github.com/caffix/cloudflare-roundtripper/cfrt"
	"net"
	"net/http"
	"time"
)

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