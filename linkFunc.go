package web_crawler

import (
	"strings"

	"golang.org/x/net/html"
)

func isAncorTag(tokenType html.TokenType, token html.Token) bool {
	return tokenType == html.StartTagToken && token.DataAtom.String() == "a"
}

func formatURL(baseURL string, linkURL string) string {
	base := strings.TrimSuffix(baseURL, "/")

	switch {
	case strings.HasPrefix(linkURL, "https://"):
	case strings.HasPrefix(linkURL, "http://"):
		if strings.Contains(linkURL, base) {
			return linkURL
		}
		return ""
	case strings.HasPrefix(linkURL, "/"):
		return base + linkURL
	}
	return ""
}

func extractLinks(token html.Token, url string) (string, bool) {
	for _, attr := range token.Attr {
		if attr.Key == "href" {
			link := attr.Val
			url := formatURL(url, link)
			if url == "" {
				break
			}
			return url, true
		}
	}
	return "", false
}
