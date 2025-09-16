# StreamingCommunity Refactor

This repository contains the refactored StreamingCommunity application, now featuring a Go backend and a SvelteKit frontend.

## Table of Contents
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Backend (Go)](#backend-go)
  - [Frontend (SvelteKit)](#frontend-sveltekit)
- [Running the Application](#running-the-application)
- [Features](#features)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

### Prerequisites

Before you begin, ensure you have the following installed:

*   **Go**: Version 1.21 or higher. [Download & Install Go](https://golang.org/doc/install)
*   **Node.js**: Version 18 or higher. [Download & Install Node.js](https://nodejs.org/en/download/)
*   **pnpm**: A fast, disk-space efficient package manager. [Install pnpm](https://pnpm.io/installation)
*   **yt-dlp**: A command-line program to download videos from YouTube.com and other video sites. [Install yt-dlp](https://github.com/yt-dlp/yt-dlp#installation)
*   **Playwright Browsers**: Required for frontend testing. [Install Playwright Browsers](https://playwright.dev/docs/browsers)

### Backend (Go)

1.  Navigate to the `backend` directory:

    ```bash
    cd backend
    ```

2.  Install Go dependencies:

    ```bash
    go mod tidy
    ```

3.  Build the backend application:

    ```bash
    go build -o streamingcommunity-backend
    ```

### Frontend (SvelteKit)

1.  Navigate to the `frontend` directory:

    ```bash
    cd frontend
    ```

2.  Install pnpm dependencies:

    ```bash
    pnpm install
    ```

3.  Build the frontend application and copy it to the backend's web directory:

    ```bash
    pnpm run build:backend
    ```

## Running the Application

1.  Start the backend server:

    ```bash
    cd backend
    ./streamingcommunity-backend
    ```

2.  In a separate terminal, start the frontend development server:

    ```bash
    cd frontend
    pnpm run dev
    ```

3.  Open your browser and navigate to `http://localhost:5173` (or the address shown in your terminal).

## Features

*   **Go Backend**: Robust and efficient API services.
*   **SvelteKit Frontend**: Modern and reactive user interface.
*   **Multi-site Scraping**: Support for various streaming platforms.
*   **HLS/DASH/MP4/Torrent Downloading**: Comprehensive download capabilities.
*   **DRM Handling (Placeholder)**: Basic structure for Widevine DRM integration.
*   **Unit & Integration Tests**: Ensures code quality and reliability.

## Contributing

Contributions are welcome! Please see the [CONTRIBUTING.md](CONTRIBUTING.md) for details on how to contribute.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
