package downloader

import (
	"context"
	"fmt"

	"github.com/autobrr/go-qbittorrent"
)

type TorrentDownloader struct {
	Host     string
	Username string
	Password string
}

func NewTorrentDownloader(host, username, password string) *TorrentDownloader {
	return &TorrentDownloader{
		Host:     host,
		Username: username,
		Password: password,
	}
}

func (d *TorrentDownloader) Download(torrentURL string, downloadPath string) error {
	// Create a new qBittorrent client
	cfg := qbittorrent.Config{
		Host:     d.Host,
		Username: d.Username,
		Password: d.Password,
	}
	client := qbittorrent.NewClient(cfg)

	// Login to qBittorrent
	if err := client.Login(context.Background()); err != nil {
		return fmt.Errorf("failed to login to qBittorrent: %w", err)
	}

	// Add the torrent
	_, err := client.Add(context.Background(), qbittorrent.AddOptions{
		URLs:     []string{torrentURL},
		Savepath: downloadPath,
	})
	if err != nil {
		return fmt.Errorf("failed to add torrent: %w", err)
	}

	return nil
}
