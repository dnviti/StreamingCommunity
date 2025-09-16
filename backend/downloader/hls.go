package downloader

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/grafov/m3u8"
)

type HLSDownloader struct {
	URL    string
	Output string
}

func NewHLSDownloader(url, output string) *HLSDownloader {
	return &HLSDownloader{
		URL:    url,
		Output: output,
	}
}

func (d *HLSDownloader) Download() error {
	// Create output directory if it doesn't exist
	if err := os.MkdirAll(filepath.Dir(d.Output), 0755); err != nil {
		return err
	}

	// Get the playlist

	playlist, listType, err := d.getPlaylist()
	if err != nil {
		return err
	}

	var segments []*m3u8.MediaSegment

	// If it's a master playlist, choose a variant
	if listType == m3u8.MASTER {
		masterPl := playlist.(*m3u8.MasterPlaylist)
		// For simplicity, we'll just choose the first variant
		if len(masterPl.Variants) > 0 {
			variant := masterPl.Variants[0]
			d.URL, err = d.resolveURL(variant.URI)
			if err != nil {
				return err
			}
			mediaPl, _, err := d.getPlaylist()
			if err != nil {
				return err
			}
			segments = mediaPl.(*m3u8.MediaPlaylist).Segments
		}
	} else if listType == m3u8.MEDIA {
		mediaPl := playlist.(*m3u8.MediaPlaylist)
		segments = mediaPl.Segments
	}

	// Download segments
	var segmentURIs []string
	for _, segment := range segments {
		if segment != nil {
			segmentURL, err := d.resolveURL(segment.URI)
			if err != nil {
				return err
			}
			segmentURIs = append(segmentURIs, segmentURL)
		}
	}

	// Create a temporary directory for segments
	tempDir, err := os.MkdirTemp("", "hls-segments-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)

	// Download each segment
	for i, segmentURI := range segmentURIs {
		err := d.downloadSegment(segmentURI, filepath.Join(tempDir, fmt.Sprintf("segment-%d.ts", i)))
		if err != nil {
			return err
		}
	}

	// Concatenate segments using ffmpeg
	return d.concatenateSegments(tempDir)
}

func (d *HLSDownloader) getPlaylist() (m3u8.Playlist, m3u8.ListType, error) {
	resp, err := http.Get(d.URL)
	if err != nil {
		return nil, 0, err
	}
	defer resp.Body.Close()

	playlist, listType, err := m3u8.DecodeFrom(resp.Body, true)
	if err != nil {
		return nil, 0, err
	}

	return playlist, listType, nil
}

func (d *HLSDownloader) resolveURL(uri string) (string, error) {
	if strings.HasPrefix(uri, "http") {
		return uri, nil
	}

	baseURL, err := url.Parse(d.URL)
	if err != nil {
		return "", err
	}

	relativeURL, err := url.Parse(uri)
	if err != nil {
		return "", err
	}

	return baseURL.ResolveReference(relativeURL).String(), nil
}

func (d *HLSDownloader) downloadSegment(url, outputPath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func (d *HLSDownloader) concatenateSegments(tempDir string) error {
	// Create a file list for ffmpeg
	fileListPath := filepath.Join(tempDir, "filelist.txt")
	fileList, err := os.Create(fileListPath)
	if err != nil {
		return err
	}
	defer fileList.Close()

	files, err := filepath.Glob(filepath.Join(tempDir, "*.ts"))
	if err != nil {
		return err
	}

	for _, file := range files {
		_, err := fileList.WriteString(fmt.Sprintf("file \"%s\"\n", file))
		if err != nil {
			return err
		}
	}

	// Run ffmpeg to concatenate
	cmd := exec.Command("ffmpeg", "-f", "concat", "-safe", "0", "-i", fileListPath, "-c", "copy", d.Output)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
