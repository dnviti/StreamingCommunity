# Quickstart

## Backend (Go)

1.  Install Go: [https://golang.org/doc/install](https://golang.org/doc/install)
2.  Navigate to the `backend` directory.
3.  Install dependencies: `go mod tidy`
4.  Run the backend: `go run .`

## Frontend (Svelte)

1.  Install Node.js and npm: [https://nodejs.org/](https://nodejs.org/)
2.  Navigate to the `frontend` directory.
3.  Install dependencies: `npm install`
4.  Run the frontend development server: `npm run dev`

## Build

1.  Build the frontend: `npm run build` in the `frontend` directory.
2.  Copy the contents of `frontend/build` to `backend/web`.
3.  Build the backend: `go build` in the `backend` directory.
4.  Run the application: `./backend`

## Additional Dependencies

*   **Widevine CDM:** For DRM-protected content, you will need the Widevine CDM library (`libwidevinecdm.so` on Linux). This is typically extracted from a Chrome or Firefox installation.
*   **C/C++ Compiler:** A C/C++ compiler (like GCC) is required for Cgo to wrap Python or C++ libraries for DRM handling.