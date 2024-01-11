package main

import "github.com/web-crawler/web_crawler"

func main() {
	sitesToCrawl := []string{"https://theuselessweb.com/"}
	web_crawler.CrawlFunc(sitesToCrawl)
}
