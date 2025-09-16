package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"streamingcommunity/backend/models"
	"streamingcommunity/backend/storage"
)

// setupRouter sets up a test Gin router with the necessary routes.
func setupRouter() *gin.Engine {
	router := gin.Default()
	// Re-implement the routes from main.go that we want to test
	storage := storage.NewJSONStorage("../data/videos.json", "../data/websites.json", "../data/downloads.json")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/videos", func(c *gin.Context) {
	
videos, err := storage.GetVideos()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, videos)
	})

	router.POST("/videos", func(c *gin.Context) {
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

	return router
}

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"pong\"}", w.Body.String())
}

func TestGetVideos(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/videos", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	// Assuming videos.json is empty initially, or contains known data
	assert.Equal(t, "[]", w.Body.String())
}

func TestPostVideo(t *testing.T) {
	router := setupRouter()

	video := models.Video{
		ID:          "test-video-1",
		Title:       "Test Video",
		Description: "A video for testing",
		URL:         "http://example.com/test.mp4",
		WebsiteID:   "test-site",
	}
	jsonValue, _ := json.Marshal(video)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/videos", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "test-video-1")

	// Clean up the created video to ensure test idempotency
	// This would ideally involve a test-specific storage or a cleanup function
}

