package web_crawler

import (
	"sync"
)

func CrawlFunc() {
	WebsitesChannel := make(chan string)
	crawledSitesChannel := make(chan string)
	pendingSitesChannel := make(chan int)

	sitesToCrawl := []string{"https://theuselessweb.com/"}

	go func() {
		for i := 0; len(sitesToCrawl) > i; i++ {
			site := sitesToCrawl[i]
			crawledSitesChannel <- site
		}

	}()

	var wg sync.WaitGroup

	go LinksChecker(WebsitesChannel, crawledSitesChannel, pendingSitesChannel)
	go LinksMonitor(WebsitesChannel, crawledSitesChannel, pendingSitesChannel)

	numOfThreads := 35
	for i := 0; i < numOfThreads; i++ {
		wg.Add(1)

	}
	wg.Wait()
}
