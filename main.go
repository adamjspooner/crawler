package main

import (
	"log"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args[1:]
	if len(args) < 3 {
		log.Fatal("too few arguments provided")
	}
	if len(args) > 3 {
		log.Fatal("too many arguments provided")
	}

	u := args[0]

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		log.Fatal(err)
	}

	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		log.Fatal(err)
	}

	baseURL, err := url.Parse(u)
	if err != nil {
		log.Fatal(err)
	}

	cfg := &config{
		pages:              make(map[string]int),
		maxPages:           maxPages,
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(u)
	cfg.wg.Wait()

	cfg.printReport()
}
