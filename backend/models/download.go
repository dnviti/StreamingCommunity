package models

type Download struct {
	ID      string `json:"id"`
	VideoID string `json:"video_id"`
	Status  string `json:"status"` // e.g., "pending", "downloading", "completed", "failed"
}
