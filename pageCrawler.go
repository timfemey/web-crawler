package web_crawler

import (
	"fmt"
	"sync"
)

func CrawlWebPage(wg *sync.WaitGroup, WebsitesChannel chan string,
	crawledSitesChannel chan string,
	pendingSitesChannel chan int) {
	crawled := 0

	// for url := range WebsitesChannel {
	// 	extractContent(url,crawledSitesChannel)
	// 	pendingSitesChannel <- -1
	// 	crawled++
	// }

	fmt.Println("Crawled ", crawled, " web sites")

	wg.Done()

}
