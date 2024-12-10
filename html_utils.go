package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func normalizeURL(rawURL string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	if strings.HasSuffix(u.Path, "/") {
		u.Path = u.Path[:len(u.Path)-1]
	}
	return u.Host + u.Path, nil
}

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("HTTP Get Error: %v", err)
	}
	if !strings.Contains(res.Header["Content-Type"][0], "text/html") {
		return "", fmt.Errorf("Expected content type to be 'text/html', got: %s\n", res.Header["Content-Type"])
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 400 {
		return "", fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		return "", fmt.Errorf("%v", err)
	}
	return string(body), nil
}

func getURLsFromHTML(htmlBody string, rawBaseURL string) ([]string, error) {
	urls := []string{}
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, err
	}
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			for _, a := range n.Attr {
				if a.Key == "href" {
					if strings.HasPrefix(a.Val, "/") {
						a.Val = rawBaseURL + a.Val
					}
					if strings.HasPrefix(a.Val, "#") || strings.HasPrefix(a.Val, "?") {
						a.Val = ""
					}
					urls = append(urls, a.Val)
					break
				}
			}
		}
	}
	return urls, nil
}
