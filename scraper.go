package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Scraper struct {
	url string
}

func NewScraper(url string) Scraper {
	return Scraper{url: url}
}

func (s *Scraper) Scrape() (HeadlineInfo, error) {
	res, err := http.Get(s.url)
	if err != nil {
		return HeadlineInfo{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return HeadlineInfo{}, fmt.Errorf("request failed with status code: %d", res.StatusCode)
	}

	// Parse the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return HeadlineInfo{}, err
	}

	currentTime := time.Now()

	// Format the current time as a string
	currentTimeString := currentTime.Format("2006-01-02 15:04:05")

	var headlines []string

	// Find and extract the <a> elements with class "post-card__title-link"
	doc.Find("a.post-card__title-link").Each(func(i int, s *goquery.Selection) {
		// Extract the "title" from each <a> element
		title := s.Text()
		fmt.Printf("Title %d: %s\n", i+1, title)

		title = strings.ReplaceAll(title, "‘", "")
		headlines = append(headlines, title)
	})

	return HeadlineInfo{
		Time:      currentTimeString,
		Headlines: headlines,
	}, nil
}
