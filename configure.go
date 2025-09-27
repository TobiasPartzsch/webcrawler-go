package main

import (
	"fmt"
	"net/url"
	"sync"
)

type cliInput struct {
	baseURL        string
	maxConcurrency int
	maxPages       int
}
type config struct {
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

// addPageVisit returns true if this is the first time we see the URL.
// We insert a placeholder PageData to mark it as visited; it’ll be replaced later.
func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, visited := cfg.pages[normalizedURL]; visited {
		return false
	}

	cfg.pages[normalizedURL] = PageData{URL: normalizedURL}
	return true
}

// setPageData safely stores the final PageData for a URL.
func (cfg *config) setPageData(normalizedURL string, data PageData) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	cfg.pages[normalizedURL] = data
}

func (cfg *config) reachedMaxPages() bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages) >= cfg.maxPages
}

func configure(ci cliInput) (*config, error) {
	baseURL, err := url.Parse(ci.baseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	return &config{
		pages:              make(map[string]PageData),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, ci.maxConcurrency),
		wg:                 &sync.WaitGroup{},
		maxPages:           ci.maxPages,
	}, nil
}
