package guardaserie

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Episode struct {
	Title       string
	URL         string
	HLSVideoURL string // New field for HLS video URL
}

type Season struct {
	Title    string
	Episodes []Episode
}

type SeriesDetails struct {
	Title       string
	Description string
	Seasons     []Season
}

func GetSeriesDetails(seriesURL string) (*SeriesDetails, error) {
	resp, err := http.Get(seriesURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make request to series URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request to series URL failed with status: %s", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML for series details: %w", err)
	}

	series := &SeriesDetails{}

	// Extract Title
	series.Title = doc.Find("h1.entry-title").Text()

	// Extract Description
	series.Description = doc.Find(".entry-content p").First().Text()

	// Extract Seasons and Episodes
	doc.Find(".tv-season").Each(func(i int, s *goquery.Selection) {
		seasonTitle := s.Find(".tv-season-title").Text()
		season := Season{Title: seasonTitle}

		s.Find(".tv-episode").Each(func(j int, ep *goquery.Selection) {
			episodeTitle := ep.Find(".tv-episode-title a").Text()
			episodeURL, exists := ep.Find(".tv-episode-title a").Attr("href")
			if exists && episodeTitle != "" {
				newEpisode := Episode{Title: episodeTitle, URL: episodeURL}

				// Attempt to find HLS video URL (similar basic approach as in altadefinizione/film.go)
				ep.Find("script").Each(func(k int, script *goquery.Selection) {
					scriptContent := script.Text()
					if (newEpisode.HLSVideoURL == "") && (len(scriptContent) > 0) {
						startIndex := -1
						for l := 0; l < len(scriptContent)-5; l++ {
							if scriptContent[l:l+5] == ".m3u8" {
								startIndex = l
								break
							}
						}

						if startIndex != -1 {
							endIndex := startIndex + 5
							urlStart := -1
							for l := startIndex; l >= 0; l-- {
								if scriptContent[l:l+4] == "http" {
									urlStart = l
									break
								}
							}
							if urlStart != -1 {
								newEpisode.HLSVideoURL = scriptContent[urlStart: endIndex]
							}
						}
					}
				})
				season.Episodes = append(season.Episodes, newEpisode)
			}
		})
		series.Seasons = append(series.Seasons, season)
	})

	return series, nil
}