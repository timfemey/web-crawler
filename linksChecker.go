package web_crawler

func LinksChecker(WebsitesChannel chan string,
	crawledSitesChannel chan string,
	pendingSitesChannel chan int) {
	discoveredLinks := make(map[string]bool)

	for i := range crawledSitesChannel {
		if !discoveredLinks[i] {
			discoveredLinks[i] = true
			pendingSitesChannel <- 1
			WebsitesChannel <- i
		}
	}
}
