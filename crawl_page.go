package main

import (
	"fmt"
	"net/url"
)

const crawlMsg = "crawling page '%s'\n"
const errMsgParsing = "Error - crawlPage: couldn't parse URL '%s': %v\n"
const errMsgNormalization = "Error - normalizedURL: %v\n"
const errMsgHTML = "Error - getHTML: %v\n"

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf(errMsgParsing, rawCurrentURL, err)
		return
	}

	// stay within the same site
	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normalizedURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf(errMsgNormalization, err)
		return
	}

	if cfg.reachedMaxPages() {
		return
	}

	// Only proceed the first time we see this normalized URL
	isFirst := cfg.addPageVisit(normalizedURL)
	if !isFirst {
		return
	}

	fmt.Printf(crawlMsg, rawCurrentURL)

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf(errMsgHTML, err)
		return
	}

	// Extract all the data we care about and store it
	pageData := extractPageData(htmlBody, rawCurrentURL)
	cfg.setPageData(normalizedURL, pageData)

	// Recurse using the already-extracted outgoing links
	for _, nextURL := range pageData.OutgoingLinks {
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
	}
}
