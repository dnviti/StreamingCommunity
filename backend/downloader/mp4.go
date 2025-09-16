package downloader

import (
	"io"
	"net/http"
	"os"
)

type MP4Downloader struct {
	URL    string
	Output string
}

func NewMP4Downloader(url, output string) *MP4Downloader {
	return &MP4Downloader{
		URL:    url,
		Output: output,
	}
}

func (d *MP4Downloader) Download() error {
	resp, err := http.Get(d.URL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(d.Output)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
