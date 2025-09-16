package streamingwatch

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type FilmDetails struct {
	Title       string
	Description string
	VideoURL    string // Placeholder for the actual video URL
	HLSVideoURL string // New field for HLS video URL
}

func GetFilmDetails(filmURL string) (*FilmDetails, error) {
	resp, err := http.Get(filmURL)
	if err != nil {
		return nil, fmt.Errorf("failed to make request to film URL: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("request to film URL failed with status: %s", resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML for film details: %w", err)
	}

	film := &FilmDetails{}

	// Extract Title
	film.Title = doc.Find("h1.entry-title").Text()

	// Extract Description
	film.Description = doc.Find(".entry-content p").First().Text()

	// Placeholder for VideoURL - this will require more advanced scraping
	videoURL, exists := doc.Find("iframe").Attr("src")
	if exists {
		film.VideoURL = videoURL
	}

	// Attempt to find HLS video URL (similar basic approach as in altadefinizione/film.go)
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		scriptContent := s.Text()
		if (film.HLSVideoURL == "") && (len(scriptContent) > 0) && (len(film.VideoURL) == 0) {
			// Look for .m3u8 in script content
			// This is a very naive approach and needs to be improved for robustness
			// A more robust solution would involve regex or a proper JavaScript parser
			startIndex := -1
			for i := 0; i < len(scriptContent)-5; i++ {
				if scriptContent[i:i+5] == ".m3u8" {
					startIndex = i
					break
				}
			}

			if startIndex != -1 {
				endIndex := startIndex + 5
				// Find the start of the URL (e.g., http or https)
				urlStart := -1
				for i := startIndex; i >= 0; i-- {
					if scriptContent[i:i+4] == "http" {
						urlStart = i
						break
					}
				}
				if urlStart != -1 {
					film.HLSVideoURL = scriptContent[urlStart: endIndex]
				}
			}
		}
	})

	return film, nil
}