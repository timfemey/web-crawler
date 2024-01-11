package web_crawler

func LinksMonitor(WebsitesChannel chan string,
	crawledSitesChannel chan string,
	pendingSitesChannel chan int) {
	count := 0
	for c := range pendingSitesChannel {
		count += c
		if count == 0 {
			close(WebsitesChannel)
			close(crawledSitesChannel)
			close(pendingSitesChannel)
		}
	}
}
