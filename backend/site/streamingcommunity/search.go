package streamingcommunity

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type SearchResult struct {
	Title string
	URL   string
}

func Search(query string) ([]SearchResult, error) {
	baseURL := "https://streamingcommunity.one"
	searchURL := fmt.Sprintf("%s/?s=%s", baseURL, url.QueryEscape(query))

	resp, err := http.Get(searchURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make search request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("search request failed with status: %s", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	var results []SearchResult
	doc.Find(".movie-card").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".movie-card-title a").Text()
		href, exists := s.Find(".movie-card-title a").Attr("href")
		if exists && title != "" {
			results = append(results, SearchResult{Title: title, URL: href})
		}
	})

	return results, nil
}
