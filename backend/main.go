package main

import (
	"streamingcommunity/backend/models"
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

	r.Run() // listen and serve on 0.0.0.0:8080
}
