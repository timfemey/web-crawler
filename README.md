# Web Crawler in Golang

This is a simple yet powerful web crawler built in Golang that utilizes GoRoutines for concurrent and fast crawling. The crawler is designed to prevent duplicated websites and efficiently parse HTML using the HTML Tokenizer.

## Features

- **Concurrent Crawling:** The web crawler uses GoRoutines to crawl multiple web pages concurrently, making the crawling process faster and more efficient.

- **Duplicate Prevention:** The crawler is equipped with mechanisms to prevent crawling duplicate websites, ensuring that each unique URL is crawled only once.

- **Efficient HTML Parsing:** The HTML Tokenizer is employed for efficient parsing of HTML content, making the crawler more robust and resource-friendly.

## Installation

1. Make sure you have Golang installed. If not, you can download it from [Golang Official Website](https://golang.org/dl/).

2. Clone this repository:

   ```bash
   git clone https://github.com/timfemey/web-crawler.git
   ```

3. Change to the project directory:

   ```bash
   cd web-crawler
   ```

4. Build and run the web crawler:

   ```bash
   go build
   ./web-crawler
   ```

## Usage

Modify the `main.go` file to specify the starting URL and other sites. Then, run the application following the installation steps above.

```go
// main.go

package main

import "github.com/web-crawler/web_crawler"

func main() {
	sitesToCrawl := []string{"https://theuselessweb.com/"}
	web_crawler.CrawlFunc(sitesToCrawl)
}

```

## Contributing

Feel free to contribute to the development of this web crawler. Create issues, submit pull requests, and help make it even better!

## Acknowledgments

- Thanks to the Golang community for providing a powerful language for building efficient and concurrent applications.

Happy Crawling! 🕷️🚀
