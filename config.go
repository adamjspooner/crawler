package main

import (
	"fmt"
	"log"
	"net/url"
	"sort"
	"strings"
	"sync"
)

type config struct {
	pages              map[string]int
	maxPages           int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, exists := cfg.pages[normalizedURL]; exists {
		cfg.pages[normalizedURL]++
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()

	normalizedRawCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		return
	}

	if !strings.HasPrefix(rawCurrentURL, cfg.baseURL.String()) {
		return
	}

	cu, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Println(err)
	}

	if cu.Scheme != cfg.baseURL.Scheme {
		return
	}

	if first := cfg.addPageVisit(normalizedRawCurrentURL); !first {
		return
	}

	html, err := getHTML(rawCurrentURL)
	if err != nil {
		return
	}

	urls, err := getURLsFromHTML(html, cfg.baseURL.String())
	if err != nil {
		return
	}

	for _, u := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(u)
	}
}

func (cfg *config) printReport() {
	type pageCount struct {
		Page  string
		Count int
	}

	var pc []pageCount
	for p, c := range cfg.pages {
		pc = append(pc, pageCount{p, c})
	}

	sort.Slice(pc, func(i, j int) bool {
		return pc[i].Count > pc[j].Count
	})

	fmt.Printf(`=============================
  REPORT for %s
=============================
`, cfg.baseURL)

	for _, _pc := range pc {
		pluralizeLink := "link"
		if _pc.Count != 1 {
			pluralizeLink = "links"
		}
		fmt.Printf("Found %d internal %s to %v\n", _pc.Count, pluralizeLink, _pc.Page)
	}
}
