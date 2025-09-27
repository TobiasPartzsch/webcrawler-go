package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

var (
	errNotEnoughArgs = errors.New("not enough arguments provided")
	errTooManyArg    = errors.New("too many arguments provided")
	errArgConversion = errors.New("couldn't convert argument to int")
)

const startingCrawlMsg = "starting crawl of: %s\n"

func main() {
	cliInput, err := validateArgs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf(startingCrawlMsg, cliInput.baseURL)

	cfg, err := configure(cliInput)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(cliInput.baseURL)
	cfg.wg.Wait()

	if err := writeCSVRport(cfg.pages, "report.csv"); err != nil {
		fmt.Printf("error writing CSV: %v\n", err)
		os.Exit(1)
	}
}

func validateArgs() (cliInput, error) {
	switch len(os.Args) {
	case 1, 2, 3:
		return cliInput{}, errNotEnoughArgs
	case 4:
		maxConcurrency, err := strconv.Atoi(os.Args[2])
		if err != nil {
			return cliInput{}, errArgConversion
		}
		maxPages, err := strconv.Atoi(os.Args[3])
		if err != nil {
			return cliInput{}, errArgConversion
		}
		return cliInput{
			baseURL:        os.Args[1],
			maxConcurrency: maxConcurrency,
			maxPages:       maxPages,
		}, nil
	default:
		return cliInput{}, errTooManyArg
	}
}
