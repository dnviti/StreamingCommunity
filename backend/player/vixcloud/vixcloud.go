package vixcloud

import (
	"fmt"
	"os/exec"
)

// GetVideoURL attempts to extract the direct video URL from a Vixcloud embed URL.
// This is a placeholder and assumes yt-dlp can handle the URL.
func GetVideoURL(embedURL string) (string, error) {
	// Use yt-dlp to extract the direct video URL
	cmd := exec.Command("yt-dlp", "--get-url", embedURL)
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get video URL using yt-dlp: %w", err)
	}

	return string(out), nil
}
