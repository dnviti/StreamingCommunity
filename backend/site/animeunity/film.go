package animeunity

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type FilmDetails struct {
	Title       string
	Description string
	VideoURL    string // Placeholder for the actual video URL
	MP4VideoURL string // New field for MP4 video URL
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

	// Attempt to find MP4 video URL (this is a very basic example and might need refinement)
	doc.Find("video source").Each(func(i int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if exists && film.MP4VideoURL == "" && (len(src) > 0) {
			if (len(src) > 4) && (src[len(src)-4:] == ".mp4") {
				film.MP4VideoURL = src
			}
		}
	})

	return film, nil
}