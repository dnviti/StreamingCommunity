# Quickstart Guide: Refactor StreamingCommunity to a Web Application

This guide outlines the steps to quickly set up and verify the core functionalities of the refactored StreamingCommunity web application.

## Prerequisites
- Docker (or Go and Node.js/npm/pnpm installed locally)
- Git

## Setup
1.  Clone the repository:
    ```bash
    git clone [repository_url]
    cd StreamingCommunity
    ```
2.  Build and run the backend (Go):
    ```bash
    # Instructions to build and run the Go backend
    # This will likely involve `go build` and running the executable
    # Ensure the backend serves the frontend assets from backend/web
    ```
3.  Build the frontend (Svelte):
    ```bash
    cd frontend
    pnpm install
    pnpm build
    # The build output should be placed in backend/web as per constraints
    ```
4.  Access the application:
    - Open your web browser and navigate to `http://localhost:[backend_port]`

## Verification Steps

### 1. Search and Download a Movie
- **Action**: On the web interface, use the search bar to search for a known movie title (e.g., "The Matrix").
- **Expected Result**: A list of search results is displayed. Select a movie and initiate the download. The download progress should be visible, and the movie file should appear in your local download directory upon completion.

### 2. Search and Download a TV Show Season
- **Action**: Search for a known TV show title (e.g., "Breaking Bad").
- **Expected Result**: Search results for the TV show are displayed. Select the TV show and initiate the download for a season or all episodes. All episodes should download successfully.

### 3. Search and Download an Anime Series (including DRM content)
- **Action**: Search for a known anime series title (e.g., "Attack on Titan"). For DRM content, search for a Crunchyroll anime title.
- **Expected Result**: Search results for the anime are displayed. Initiate the download. All episodes should download successfully, with DRM content handled correctly.

### 4. Global Search Functionality
- **Action**: Use the global search bar with a broad query (e.g., "action").
- **Expected Result**: A consolidated list of results from all supported websites is displayed.

### 5. Error Handling for Invalid/Unsupported Titles
- **Action**: Search for a non-existent title (e.g., "asdfghjkl").
- **Expected Result**: An informative error message is displayed, indicating that no results were found.
- **Action**: Attempt to download from an unsupported website (if such an option is presented, or by attempting to manipulate a request).
- **Expected Result**: An error message indicating the website is not supported.
