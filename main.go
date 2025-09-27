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

	pages := make(map[string]int)

	crawlPage(rawBaseURL, rawBaseURL, pages)

	for normalizedURL, count := range pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
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
