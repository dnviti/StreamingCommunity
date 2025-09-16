package raiplay

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
	DASHVideoURL string // New field for DASH video URL
	PSSH        string // New field for PSSH (DRM)
	LicenseURL  string // New field for License URL (DRM)
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
	film.Title = doc.Find("h1.title").Text()

	// Extract Description
	film.Description = doc.Find(".description").First().Text()

	// Placeholder for VideoURL - this will require more advanced scraping
	// RaiPlay uses a complex player, so this will likely involve yt-dlp or similar
	// For now, we'll just return the filmURL as a placeholder for the video source
	film.VideoURL = filmURL

	// Attempt to find HLS/DASH video URL, PSSH, and LicenseURL
	// This is highly dependent on RaiPlay's specific implementation
	// and will likely require more sophisticated parsing (e.g., JavaScript execution, network analysis)
	// For now, these are placeholders.
	film.HLSVideoURL = ""
	film.DASHVideoURL = ""
	film.PSSH = ""
	film.LicenseURL = ""

	return film, nil
}