# Data Model

## Video

*   `id`: string (UUID)
*   `title`: string
*   `source_url`: string
*   `status`: string (e.g., "pending", "downloading", "completed", "failed")
*   `website_id`: string (foreign key to Website)
*   `progress`: float
*   `drm_protected`: boolean
*   `license_url`: string (optional)
*   `pssh`: string (optional)
*   `downloader_type`: string (e.g., "hls", "dash", "mp4", "torrent")
*   `season_number`: int (optional)
*   `episode_number`: int (optional)

## Website

*   `id`: string (UUID)
*   `name`: string
*   `url`: string
*   `player_type`: string (e.g., "vixcloud", "supervideo")
*   `search_path`: string
*   `film_path`: string (optional)
*   `series_path`: string (optional)