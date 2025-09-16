package downloader

import (
	"os"
	"os/exec"
)

type DASHDownloader struct {
	URL    string
	Output string
}

func NewDASHDownloader(url, output string) *DASHDownloader {
	return &DASHDownloader{
		URL:    url,
		Output: output,
	}
}

func (d *DASHDownloader) Download() error {
	cmd := exec.Command("yt-dlp", "-o", d.Output, d.URL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}