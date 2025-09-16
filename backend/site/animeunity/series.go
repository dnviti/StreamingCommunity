package animeunity

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type Episode struct {
	Title       string
	URL         string
	MP4VideoURL string // New field for MP4 video URL
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
	doc.Find(".season-item").Each(func(i int, s *goquery.Selection) {
		seasonTitle := s.Find(".season-title").Text()
		season := Season{Title: seasonTitle}

		s.Find(".episode-item").Each(func(j int, ep *goquery.Selection) {
			episodeTitle := ep.Find(".episode-title a").Text()
			episodeURL, exists := ep.Find(".episode-title a").Attr("href")
			if exists && episodeTitle != "" {
				newEpisode := Episode{Title: episodeTitle, URL: episodeURL}

				// Attempt to find MP4 video URL
				ep.Find("video source").Each(func(k int, script *goquery.Selection) {
					src, exists := script.Attr("src")
					if exists && newEpisode.MP4VideoURL == "" && (len(src) > 0) {
						if (len(src) > 4) && (src[len(src)-4:] == ".mp4") {
							newEpisode.MP4VideoURL = src
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