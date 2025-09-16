package main

import (
	"streamingcommunity/backend/downloader"
	"streamingcommunity/backend/models"
	"streamingcommunity/backend/player/hdplayer"
	"streamingcommunity/backend/player/supervideo"
	"streamingcommunity/backend/player/sweetpixel"
	"streamingcommunity/backend/player/vixcloud"
	"streamingcommunity/backend/site/altadefinizione"
	"streamingcommunity/backend/site/animeunity"
	"streamingcommunity/backend/site/animeworld"
	"streamingcommunity/backend/site/crunchyroll"
	"streamingcommunity/backend/site/mediasetinfinity"
	"streamingcommunity/backend/site/raiplay"
	"streamingcommunity/backend/site/streamingcommunity"
	"streamingcommunity/backend/site/streamingwatch"
	"streamingcommunity/backend/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	storage := storage.NewJSONStorage("data/videos.json", "data/websites.json", "data/downloads.json")

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/videos", func(c *gin.Context) {
		videos, err := storage.GetVideos()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, videos)
	})

	r.POST("/videos", func(c *gin.Context) {
		var video models.Video
		if err := c.ShouldBindJSON(&video); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		videos, err := storage.GetVideos()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		videos = append(videos, video)
		if err := storage.SaveVideos(videos); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, video)
	})

	r.GET("/websites", func(c *gin.Context) {
		websites, err := storage.GetWebsites()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, websites)
	})

	r.POST("/websites", func(c *gin.Context) {
		var website models.Website
		if err := c.ShouldBindJSON(&website); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		websites, err := storage.GetWebsites()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		websites = append(websites, website)
		if err := storage.SaveWebsites(websites); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, website)
	})

	r.GET("/downloads", func(c *gin.Context) {
		downloads, err := storage.GetDownloads()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, downloads)
	})

	r.POST("/downloads", func(c *gin.Context) {
		var download models.Download
		if err := c.ShouldBindJSON(&download); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		downloads, err := storage.GetDownloads()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		downloads = append(downloads, download)
		if err := storage.SaveDownloads(downloads); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, download)
	})

	r.GET("/altadefinizione/search", func(c *gin.Context) {
		query := c.Query("q")
		if query == "" {
			c.JSON(400, gin.H{"error": "query parameter 'q' is required"})
			return
		}

		results, err := altadefinizione.Search(query)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, results)
	})

	r.GET("/altadefinizione/film", func(c *gin.Context) {
		filmURL := c.Query("url")
		if filmURL == "" {
			c.JSON(400, gin.H{"error": "url parameter is required"})
			return
		}

		filmDetails, err := altadefinizione.GetFilmDetails(filmURL)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, filmDetails)
	})

	r.GET("/altadefinizione/series", func(c *gin.Context) {
		seriesURL := c.Query("url")
		if seriesURL == "" {
			c.JSON(400, gin.H{"error": "url parameter is required"})
			return
		}

		seriesDetails, err := altadefinizione.GetSeriesDetails(seriesURL)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, seriesDetails)
	})

	r.GET("/player/supervideo", func(c *gin.Context) {
		embedURL := c.Query("embed_url")
		if embedURL == "" {
			c.JSON(400, gin.H{"error": "embed_url parameter is required"})
			return
		}

		videoURL, err := supervideo.GetVideoURL(embedURL)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"video_url": videoURL})
	})

	r.GET("/player/vixcloud", func(c *gin.Context) {
		embedURL := c.Query("embed_url")
		if embedURL == "" {
			c.JSON(400, gin.H{"error": "embed_url parameter is required"})
			return
		}

		videoURL, err := vixcloud.GetVideoURL(embedURL)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"video_url": videoURL})
	})

	r.GET("/player/sweetpixel", func(c *gin.Context) {
		embedURL := c.Query("embed_url")
		if embedURL == "" {
			c.JSON(400, gin.H{"error": "embed_url parameter is required"})
			return
		}

		videoURL, err := sweetpixel.GetVideoURL(embedURL)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"video_url": videoURL})
	})

	r.GET("/player/hdplayer", func(c *gin.Context) {
		embedURL := c.Query("embed_url")
		if embedURL == "" {
			c.JSON(400, gin.H{"error": "embed_url parameter is required"})
			return
		}

		videoURL, err := hdplayer.GetVideoURL(embedURL)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"video_url": videoURL})
	})

	r.POST("/download/hls", func(c *gin.Context) {
		var req struct {
			URL    string `json:"url"`
			Output string `json:"output"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		hlsDownloader := downloader.NewHLSDownloader(req.URL, req.Output)
		if err := hlsDownloader.Download(); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "HLS download initiated successfully"})
	})

	r.POST("/download/mp4", func(c *gin.Context) {
		var req struct {
			URL    string `json:"url"`
			Output string `json:"output"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		mp4Downloader := downloader.NewMP4Downloader(req.URL, req.Output)
		if err := mp4Downloader.Download(); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "MP4 download initiated successfully"})
	})

	r.POST("/download/dash", func(c *gin.Context) {
		var req struct {
			URL       string `json:"url"`
			Output    string `json:"output"`
			PSSH      string `json:"pssh"`
			LicenseURL string `json:"license_url"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		dashDownloader := downloader.NewDASHDownloader(req.URL, req.Output)
		// In a real scenario, PSSH and LicenseURL would be passed to the downloader
		// and potentially used by yt-dlp or a custom DRM handling logic.
		// For this placeholder, we just initiate the download.
		if err := dashDownloader.Download(); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "DASH download initiated successfully"})
	})

	r.GET("/animeunity/search", func(c *gin.Context) {
		query := c.Query("q")
		if query == "" {
			c.JSON(400, gin.H{"error": "query parameter 'q' is required"})
			return
		}

		results, err := animeunity.Search(query)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, results)
	})

	r.GET("/animeunity/film", func(c *gin.Context) {
		filmURL := c.Query("url")
		if filmURL == "" {
			c.JSON(400, gin.H{"error": "url parameter is required"})
			return
		}

		filmDetails, err := animeunity.GetFilmDetails(filmURL)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, filmDetails)
	})

	r.GET("/animeunity/series", func(c *gin.Context) {
		seriesURL := c.Query("url")
		if seriesURL == "" {
			c.JSON(400, gin.H{"error": "url parameter is required"})
			return
		}

		seriesDetails, err := animeunity.GetSeriesDetails(seriesURL)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, seriesDetails)
	})

	r.GET("/animeworld/search", func(c *gin.Context) {
		query := c.Query("q")
		if query == "" {
			c.JSON(400, gin.H{"error": "query parameter 'q' is required"})
			return
		}

		results, err := animeworld.Search(query)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, results)
	})

	r.GET("/animeworld/film", func(c *gin.Context) {
		filmURL := c.Query("url")
		if filmURL == "" {
			c.JSON(400, gin.H{"error": "url parameter is required"})
			return
		}

		filmDetails, err := animeworld.GetFilmDetails(filmURL)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, filmDetails)
	})

	r.GET("/animeworld/series", func(c *gin.Context) {
		seriesURL := c.Query("url")
		if seriesURL == "" {
			c.JSON(400, gin.H{"error": "url parameter is required"})
			return
		}

		seriesDetails, err := animeworld.GetSeriesDetails(seriesURL)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, seriesDetails)
	})

	r.GET("/crunchyroll/search", func(c *gin.Context) {
		query := c.Query("q")
		if query == "" {
			c.JSON(400, gin.H{"error": "query parameter 'q' is required"})
			return
		}

		results, err := crunchyroll.Search(query)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, results)
	})

	r.GET("/crunchyroll/film", func(c *gin.Context) {
		filmURL := c.Query("url")
		if filmURL == "" {
			c.JSON(400, gin.H{"error": "url parameter is required"})
			return
		}

		filmDetails, err := crunchyroll.GetFilmDetails(filmURL)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, filmDetails)
	})

	r.GET("/crunchyroll/series", func(c *gin.Context) {
		seriesURL := c.Query("url")
		if seriesURL == "" {
			c.JSON(400, gin.H{"error": "url parameter is required"})
			return
		}

		seriesDetails, err := crunchyroll.GetSeriesDetails(seriesURL)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, seriesDetails)
	})

	r.GET("/crunchyroll/license", func(c *gin.Context) {
		videoID := c.Query("video_id")
		if videoID == "" {
			c.JSON(400, gin.H{"error": "video_id parameter is required"})
			return
		}

		license, err := crunchyroll.AcquireLicense(videoID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"license": license})
	})

	r.GET("/mediasetinfinity/search", func(c *gin.Context) {
		query := c.Query("q")
		if query == "" {
			c.JSON(400, gin.H{"error": "query parameter 'q' is required"})
			return
		}

		results, err := mediasetinfinity.Search(query)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, results)
	})

	r.GET("/mediasetinfinity/film", func(c *gin.Context) {
		filmURL := c.Query("url")
		if filmURL == "" {
			c.JSON(400, gin.H{"error": "url parameter is required"})
			return
		}

		filmDetails, err := mediasetinfinity.GetFilmDetails(filmURL)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, filmDetails)
	})

	r.GET("/mediasetinfinity/series", func(c *gin.Context) {
		seriesURL := c.Query("url")
		if seriesURL == "" {
			c.JSON(400, gin.H{"error": "url parameter is required"})
			return
		}

		seriesDetails, err := mediasetinfinity.GetSeriesDetails(seriesURL)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, seriesDetails)
	})

	r.GET("/raiplay/search", func(c *gin.Context) {
		query := c.Query("q")
		if query == "" {
			c.JSON(400, gin.H{"error": "query parameter 'q' is required"})
			return
		}

		results, err := raiplay.Search(query)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, results)
	})

	r.GET("/raiplay/film", func(c *gin.Context) {
		filmURL := c.Query("url")
		if filmURL == "" {
			c.JSON(400, gin.H{"error": "url parameter is required"})
			return
		}

		filmDetails, err := raiplay.GetFilmDetails(filmURL)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, filmDetails)
	})

	r.GET("/raiplay/series", func(c *gin.Context) {
		seriesURL := c.Query("url")
		if seriesURL == "" {
			c.JSON(400, gin.H{"error": "url parameter is required"})
			return
		}

		seriesDetails, err := raiplay.GetSeriesDetails(seriesURL)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, seriesDetails)
	})

	r.GET("/streamingcommunity/search", func(c *gin.Context) {
		query := c.Query("q")
		if query == "" {
			c.JSON(400, gin.H{"error": "query parameter 'q' is required"})
			return
		}

		results, err := streamingcommunity.Search(query)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, results)
	})

	r.GET("/streamingcommunity/film", func(c *gin.Context) {
		filmURL := c.Query("url")
		if filmURL == "" {
			c.JSON(400, gin.H{"error": "url parameter is required"})
			return
		}

		filmDetails, err := streamingcommunity.GetFilmDetails(filmURL)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, filmDetails)
	})

	r.GET("/streamingcommunity/series", func(c *gin.Context) {
		seriesURL := c.Query("url")
		if seriesURL == "" {
			c.JSON(400, gin.H{"error": "url parameter is required"})
			return
		}

		seriesDetails, err := streamingcommunity.GetSeriesDetails(seriesURL)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, seriesDetails)
	})

	r.GET("/streamingwatch/search", func(c *gin.Context) {
		query := c.Query("q")
		if query == "" {
			c.JSON(400, gin.H{"error": "query parameter 'q' is required"})
			return
		}

		results, err := streamingwatch.Search(query)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, results)
	})

	r.GET("/streamingwatch/film", func(c *gin.Context) {
		filmURL := c.Query("url")
		if filmURL == "" {
			c.JSON(400, gin.H{"error": "url parameter is required"})
			return
		}

		filmDetails, err := streamingwatch.GetFilmDetails(filmURL)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, filmDetails)
	})

	r.GET("/streamingwatch/series", func(c *gin.Context) {
		seriesURL := c.Query("url")
		if seriesURL == "" {
			c.JSON(400, gin.H{"error": "url parameter is required"})
			return
		}

		seriesDetails, err := streamingwatch.GetSeriesDetails(seriesURL)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, seriesDetails)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
