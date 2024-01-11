package web_crawler

import (
	"fmt"
	"net/http"
	"time"
)

func ConnectToWebsite(url string) (*http.Response, bool) {
	nilResponse := http.Response{}
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error Occured while Connecting to Website", err)
		return &nilResponse, false
	}

	request.Header.Set("User-Agent", "CrawlerBot v1.0 - This bot retrieves links and content.")
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error Occured during Request-Response Cycle", err)
		return &nilResponse, false
	}
	return response, true
}
