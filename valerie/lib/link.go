package lib

import (
	"fmt"
	"net/http"
	"net/url"
)

type linkCheck struct {
	TheURL     *url.URL
	External   bool
	StatusCode int
	Failed     bool
}

func CheckLinks(links []*linkCheck) {

	for _, link := range links {

		resp, err := http.Get(link.TheURL.String())
		if err != nil {
			fmt.Printf("Failed to check: %s\n", link.TheURL)
			continue
		}

		link.StatusCode = resp.StatusCode

		if resp.StatusCode >= 200 && resp.StatusCode < 400 {
			link.Failed = false
		} else {
			link.Failed = true
		}
	}
}

func PrepLinks(links map[string]struct{}) ([]*linkCheck, []*linkCheck) {

	var internalLinks, externalLinks []*linkCheck

	for link, _ := range links {

		aURL, err := url.Parse(link)
		if err != nil {
			fmt.Printf("Found a bad URL: %s", aURL)
			continue
		}

		if aURL.Host == "" {
			internalLinks = append(internalLinks, &linkCheck{
				TheURL:     aURL,
				External:   false,
				StatusCode: 0,
			})
		} else {
			externalLinks = append(externalLinks, &linkCheck{
				TheURL:     aURL,
				External:   true,
				StatusCode: 0,
			})
		}
	}

	return internalLinks, externalLinks
}
