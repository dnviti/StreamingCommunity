package storage

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"streamingcommunity/backend/models"
)

type JSONStorage struct {
	videoDataPath    string
	websiteDataPath  string
	downloadDataPath string
}

func NewJSONStorage(videoDataPath, websiteDataPath, downloadDataPath string) *JSONStorage {
	return &JSONStorage{
		videoDataPath:    videoDataPath,
		websiteDataPath:  websiteDataPath,
		downloadDataPath: downloadDataPath,
	}
}

func (s *JSONStorage) GetVideos() ([]models.Video, error) {
	return s.readVideos()
}

func (s *JSONStorage) GetWebsites() ([]models.Website, error) {
	return s.readWebsites()
}

func (s *JSONStorage) GetDownloads() ([]models.Download, error) {
	return s.readDownloads()
}

func (s *JSONStorage) SaveVideos(videos []models.Video) error {
	return s.writeVideos(videos)
}

func (s *JSONStorage) SaveWebsites(websites []models.Website) error {
	return s.writeWebsites(websites)
}

func (s *JSONStorage) SaveDownloads(downloads []models.Download) error {
	return s.writeDownloads(downloads)
}

func (s *JSONStorage) readVideos() ([]models.Video, error) {
	if _, err := os.Stat(s.videoDataPath); os.IsNotExist(err) {
		return []models.Video{}, nil
	}

	data, err := ioutil.ReadFile(s.videoDataPath)
	if err != nil {
		return nil, err
	}

	var videos []models.Video
	err = json.Unmarshal(data, &videos)
	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (s *JSONStorage) readWebsites() ([]models.Website, error) {
	if _, err := os.Stat(s.websiteDataPath); os.IsNotExist(err) {
		return []models.Website{}, nil
	}

	data, err := ioutil.ReadFile(s.websiteDataPath)
	if err != nil {
		return nil, err
	}

	var websites []models.Website
	err = json.Unmarshal(data, &websites)
	if err != nil {
		return nil, err
	}

	return websites, nil
}

func (s *JSONStorage) readDownloads() ([]models.Download, error) {
	if _, err := os.Stat(s.downloadDataPath); os.IsNotExist(err) {
		return []models.Download{}, nil
	}

	data, err := ioutil.ReadFile(s.downloadDataPath)
	if err != nil {
		return nil, err
	}

	var downloads []models.Download
	err = json.Unmarshal(data, &downloads)
	if err != nil {
		return nil, err
	}

	return downloads, nil
}

func (s *JSONStorage) writeVideos(videos []models.Video) error {
	data, err := json.MarshalIndent(videos, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(s.videoDataPath, data, 0644)
}

func (s *JSONStorage) writeWebsites(websites []models.Website) error {
	data, err := json.MarshalIndent(websites, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(s.websiteDataPath, data, 0644)
}

func (s *JSONStorage) writeDownloads(downloads []models.Download) error {
	data, err := json.MarshalIndent(downloads, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(s.downloadDataPath, data, 0644)
}