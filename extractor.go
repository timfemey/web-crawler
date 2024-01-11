package web_crawler

import (
	"fmt"

	"golang.org/x/net/html"
)

func extractContent(url string, crawledSitesChannel chan string) {
	resp, success := ConnectToWebsite(url)

	if !success {
		fmt.Println("An Error Occured while connecting to this website: ", url)
		return
	}

	defer resp.Body.Close()

	tokenizer := html.NewTokenizer(resp.Body)

	for {
		tokenType := tokenizer.Next()

		if tokenType == html.ErrorToken {
			return
		}

		token := tokenizer.Token()

		if isAncorTag(tokenType, token) {
			link, ok := extractLinks(token, url)
			if ok {
				go func() {
					crawledSitesChannel <- link
				}()
			}
		}

	}
}
