# Data Model: Refactor StreamingCommunity to a Web Application

## Entities

### Video
- **Description**: Represents a downloadable video, which can be a movie, a TV show episode, or an anime episode.
- **Attributes**:
    - `id`: Unique identifier (string)
    - `title`: Title of the video (string)
    - `source_url`: URL of the video on the source website (string)
    - `download_status`: Current status of the download (e.g., "pending", "downloading", "completed", "failed") (string)
    - `type`: Type of video (e.g., "movie", "tv_show", "anime") (string)
    - `website_id`: Foreign key referencing the Website entity (string)
    - `drm_protected`: Boolean indicating if the content is DRM protected (boolean)
    - `player_supported`: List of supported players for this video (list of strings)

### Website
- **Description**: Represents a supported source for videos.
- **Attributes**:
    - `id`: Unique identifier (string)
    - `name`: Name of the website (string)
    - `base_url`: Base URL of the website (string)
    - `configuration`: JSON object for any specific configuration needed for downloading from this website (JSON object)
    - `supported_protocols`: List of supported download protocols (e.g., HLS, DASH, MP4, Torrent) (list of strings)
    - `supports_drm`: Boolean indicating if the website supports DRM protected content (boolean)

## Relationships
- A `Video` belongs to one `Website`.
- A `Website` can have many `Videos`.
