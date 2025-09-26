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

const startingCrawlMsg = `starting crawl of: %s\n`

func main() {
	baseURL, err := validateArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf(startingCrawlMsg, baseURL)
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
