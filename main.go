package main

import (
	"errors"
	"fmt"
	"os"
)

var (
	errNoWebsite  = errors.New("no website provided")
	errTooManyArg = errors.New("too many arguments provided")
)

const startingCrawlMsg = "starting crawl of: %s\n"

func main() {
	rawBaseURL, err := validateArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf(startingCrawlMsg, rawBaseURL)

	const maxConcurrency = 3
	cfg, err := configure(rawBaseURL, maxConcurrency)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s...\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for normalizedURL, pd := range cfg.pages {
		// TODO: just print count
		fmt.Printf("%v - %s\n", pd, normalizedURL)
	}
}

func validateArgs() (string, error) {
	switch len(os.Args) {
	case 1:
		return "", errNoWebsite
	case 2:
		return os.Args[1], nil
	default:
		return "", errTooManyArg
	}
}
