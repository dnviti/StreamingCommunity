package crunchyroll

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Episode struct {
	Title       string
	URL         string
	DASHVideoURL string // New field for DASH video URL
	PSSH        string // New field for PSSH (DRM)
	LicenseURL  string // New field for License URL (DRM)
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
	series.Title = doc.Find("h1.hero-title").Text()

	// Extract Description
	series.Description = doc.Find(".description-text").First().Text()

	// Extract Seasons and Episodes
	doc.Find(".season-dropdown-item").Each(func(i int, s *goquery.Selection) {
		seasonTitle := s.Find(".season-title").Text()
		season := Season{Title: seasonTitle}

		s.Find(".episode-card").Each(func(j int, ep *goquery.Selection) {
			episodeTitle := ep.Find(".episode-title a").Text()
			episodeURL, exists := ep.Find(".episode-title a").Attr("href")
			if exists && episodeTitle != "" {
				newEpisode := Episode{Title: episodeTitle, URL: episodeURL}

				// Attempt to find DASH video URL, PSSH, and LicenseURL
				// This is highly dependent on Crunchyroll's specific implementation
				// and will likely require more sophisticated parsing (e.g., JavaScript execution, network analysis)
				// For now, these are placeholders.
				newEpisode.DASHVideoURL = ""
				newEpisode.PSSH = ""
				newEpisode.LicenseURL = ""

				season.Episodes = append(season.Episodes, newEpisode)
			}
		})
		series.Seasons = append(series.Seasons, season)
	})

	return series, nil
}