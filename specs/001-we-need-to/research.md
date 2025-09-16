# Research

## Go Web Framework

*   **Decision**: Use Gin as the Go web framework.
*   **Rationale**: Gin is a popular, high-performance, and well-documented framework for Go. It's a good choice for building a simple and fast backend.
*   **Alternatives considered**: Echo, Fiber.

## Frontend Framework

*   **Decision**: Use SvelteKit for the frontend.
*   **Rationale**: SvelteKit is the official framework for building Svelte applications. It provides a great developer experience with features like routing, server-side rendering, and easy build configuration.
*   **Alternatives considered**: Using Svelte without a framework.

## Web Scraping in Go

*   **Decision**: Use `colly` for web scraping.
*   **Rationale**: `colly` is a powerful and flexible scraping framework for Go that provides a clean API for traversing websites and extracting data. It's a good replacement for Python's BeautifulSoup and Scrapy.
*   **Alternatives considered**: `goquery` (a library similar to jQuery for Go), which is also a good option but `colly` provides a more complete framework.

## HLS and DASH Downloading in Go

*   **Decision**: Research and select a Go library for parsing and downloading HLS and DASH streams.
*   **Rationale**: A dedicated library will handle the complexities of parsing manifests, downloading segments, and merging them. This will save significant development time.
*   **Alternatives considered**: Implementing the logic from scratch, which would be a complex and time-consuming task.

## Widevine DRM in Go

*   **Decision**: This is a critical and challenging area. The primary approach will be to use Cgo to wrap the `pywidevine` library or a similar C++ library. If that proves too difficult, an alternative is to call a Python script as an external process to get the decryption keys.
*   **Rationale**: Native Go libraries for Widevine DRM are not mature or readily available. Wrapping a battle-tested library like `pywidevine` is the most realistic approach to ensure support for DRM-protected content.
*   **Alternatives considered**: A pure Go implementation, which is not feasible at this time due to the complexity of the Widevine CDM.

## Torrent Downloading in Go

*   **Decision**: Use the `anacrolix/torrent` library.
*   **Rationale**: This is a feature-complete torrent client library for Go that is widely used and well-maintained.
*   **Alternatives considered**: Integrating with an external torrent client like qBittorrent via its API. Using a native Go library will keep the application self-contained.

## JavaScript Execution in Go

*   **Decision**: Use `goja` to execute JavaScript code within the Go application.
*   **Rationale**: Some websites use JavaScript to obfuscate video URLs. `goja` is a performant ECMAScript 5.1(+) interpreter written in Go, which will allow us to execute these scripts and extract the necessary data without relying on a headless browser.
*   **Alternatives considered**: `otto` (another JS interpreter), or a headless browser like `chromedp`. `goja` is more performant and has better ES6 support than `otto`, and a headless browser would be too resource-intensive for this application.

## Scale and Scope

*   **Decision**: For the scale/scope, we will assume a small to medium user base (up to 1000 concurrent users) and a moderate amount of data (up to 1GB JSON file).
*   **Rationale**: This is a reasonable starting point for a refactoring project. The architecture can be scaled later if needed.
*   **Alternatives considered**: Planning for a larger scale from the beginning, which would add unnecessary complexity.