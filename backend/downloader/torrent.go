package downloader

import (
	"fmt"
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
	return fmt.Errorf("Torrent download not implemented")
}
