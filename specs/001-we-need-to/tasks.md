# Tasks: Refactor StreamingCommunity to a Web Application

**Input**: Design documents from `/specs/001-we-need-to/`
**Prerequisites**: plan.md (required), research.md, data-model.md, contracts/

## Execution Flow (main)
```
1. Load plan.md from feature directory
   → If not found: ERROR "No implementation plan found"
   → Extract: tech stack, libraries, structure
2. Load optional design documents:
   → data-model.md: Extract entities → model tasks
   → contracts/: Each file → contract test task
   → research.md: Extract decisions → setup tasks
3. Generate tasks by category:
   → Setup: project init, dependencies, linting
   → Tests: contract tests, integration tests
   → Core: models, services, CLI commands
   → Integration: DB, middleware, logging
   → Polish: unit tests, performance, docs
4. Apply task rules:
   → Different files = mark [P] for parallel
   → Same file = sequential (no [P])
   → Tests before implementation (TDD)
5. Number tasks sequentially (T001, T002...)
6. Generate dependency graph
7. Create parallel execution examples
8. Validate task completeness:
   → All contracts have tests?
   → All entities have models?
   → All endpoints implemented?
9. Return: SUCCESS (tasks ready for execution)
```

## Format: `[ID] [P?] Description`
- **[P]**: Can run in parallel (different files, no dependencies)
- Include exact file paths in descriptions

## Path Conventions
- **Web app**: `backend/src/`, `frontend/src/`

## Phase 3.1: Setup

-   [ ] T001 Initialize Go backend project in `backend/`
-   [ ] T002 Initialize Svelte frontend project in `frontend/`
-   [ ] T003 Configure Go dependencies (e.g., for JSON handling, HTTP server) in `backend/go.mod`
-   [ ] T004 Configure Svelte dependencies (e.g., Tailwind, Playwright) in `frontend/package.json`
-   [ ] T005 [P] Configure Go linting and formatting tools in `backend/`
-   [ ] T006 [P] Configure Svelte linting and formatting tools in `frontend/`
-   [ ] T007 Configure frontend build process to output to `backend/web`

## Phase 3.2: Tests First (TDD) ⚠️ MUST COMPLETE BEFORE 3.3

**CRITICAL: These tests MUST be written and MUST FAIL before ANY implementation**

-   [ ] T008 [P] Contract test for `GET /search` in `backend/tests/api_test.go`
-   [ ] T009 [P] Contract test for `POST /download` in `backend/tests/api_test.go`
-   [ ] T010 [P] Contract test for `GET /download/{downloadId}/status` in `backend/tests/api_test.go`
-   [ ] T011 [P] Contract test for `GET /websites` in `backend/tests/api_test.go`
-   [ ] T012 [P] Integration test: Search and Download a Movie (frontend Playwright test) in `frontend/e2e/download_movie.test.ts`
-   [ ] T013 [P] Integration test: Search and Download a TV Show Season (frontend Playwright test) in `frontend/e2e/download_tv_show.test.ts`
-   [ ] T014 [P] Integration test: Search and Download an Anime Series (frontend Playwright test) in `frontend/e2e/download_anime.test.ts`
-   [ ] T015 [P] Integration test: Global Search Functionality (frontend Playwright test) in `frontend/e2e/global_search.test.ts`
-   [ ] T016 [P] Integration test: Error Handling for Invalid/Unsupported Titles (frontend Playwright test) in `frontend/e2e/error_handling.test.ts`

## Phase 3.3: Core Implementation (ONLY after tests are failing)

-   [ ] T017 [P] Implement `Video` model in `backend/models/video.go`
-   [ ] T018 [P] Implement `Website` model in `backend/models/website.go`
-   [ ] T019 Implement `Website` data loading and management from `websites.json` in `backend/storage/json_storage.go`
-   [ ] T020 Implement `Video` data loading and management from `videos.json` in `backend/storage/json_storage.go`
-   [ ] T021 Implement `Download` data loading and management from `downloads.json` in `backend/storage/json_storage.go`
-   [ ] T022 Implement `GET /websites` API endpoint in `backend/main.go`
-   [ ] T023 Implement `GET /search` API endpoint in `backend/main.go`
-   [ ] T024 Implement `POST /download` API endpoint in `backend/main.go`
-   [ ] T025 Implement `GET /download/{downloadId}/status` API endpoint in `backend/main.go`
-   [ ] T026 Implement core download logic (HLS, DASH, MP4, Torrent) in `backend/downloader/`
-   [ ] T027 Implement Widevine DRM handling in `backend/drm/widevine.go`
-   [ ] T028 Implement frontend UI for search functionality in `frontend/src/routes/`
-   [ ] T029 Implement frontend UI for displaying search results in `frontend/src/lib/components/`
-   [ ] T030 Implement frontend UI for initiating downloads in `frontend/src/lib/components/`
-   [ ] T031 Implement frontend UI for displaying download status and progress in `frontend/src/lib/components/`
-   [ ] T032 Implement frontend UI for error messages in `frontend/src/lib/components/`
-   [ ] T033 Implement frontend routing and page structure in `frontend/src/routes/`

## Phase 3.4: Integration

-   [ ] T034 Integrate Go backend with existing `StreamingCommunity` download logic (e.g., `StreamingCommunity/Lib/Downloader/`)
-   [ ] T035 Implement error handling and logging for backend API endpoints
-   [ ] T036 Implement CORS for backend API
-   [ ] T037 Connect frontend to backend API endpoints

## Phase 3.5: Polish

-   [ ] T038 [P] Unit tests for `Video` model in `backend/models/video_test.go`
-   [ ] T039 [P] Unit tests for `Website` model in `backend/models/website_test.go`
-   [ ] T040 [P] Unit tests for download logic in `backend/downloader/`
-   [ ] T041 [P] Unit tests for DRM handling in `backend/drm/`
-   [ ] T042 Performance testing for API endpoints
-   [ ] T043 Update `README.md` with setup and usage instructions
-   [ ] T044 Clean up and refactor existing `StreamingCommunity` CLI code to be headless

## Dependencies
-   Setup tasks (T001-T007) before all other tasks.
-   Tests (T008-T016) before Core Implementation (T017-T033).
-   Model implementation (T017-T018) before storage and API endpoints that use them (T019-T025).
-   Download logic (T026-T027) before `POST /download` endpoint (T024).
-   Backend API endpoints (T022-T025) before frontend UI implementation (T028-T033).
-   Integration tasks (T034-T037) after Core Implementation.
-   Polish tasks (T038-T044) after all other tasks.

## Parallel Example
```
# Launch T005-T006 together:
Task: "Configure Go linting and formatting tools in backend/"
Task: "Configure Svelte linting and formatting tools in frontend/"

# Launch T008-T011 together:
Task: "Contract test for GET /search in backend/tests/api_test.go"
Task: "Contract test for POST /download in backend/tests/api_test.go"
Task: "Contract test for GET /download/{downloadId}/status in backend/tests/api_test.go"
Task: "Contract test for GET /websites in backend/tests/api_test.go"

# Launch T012-T016 together:
Task: "Integration test: Search and Download a Movie (frontend Playwright test) in frontend/e2e/download_movie.test.ts"
Task: "Integration test: Search and Download a TV Show Season (frontend Playwright test) in frontend/e2e/download_tv_show.test.ts"
Task: "Integration test: Search and Download an Anime Series (frontend Playwright test) in frontend/e2e/download_anime.test.ts"
Task: "Integration test: Global Search Functionality (frontend Playwright test) in frontend/e2e/global_search.test.ts"
Task: "Integration test: Error Handling for Invalid/Unsupported Titles (frontend Playwright test) in frontend/e2e/error_handling.test.ts"

# Launch T017-T018 together:
Task: "Implement Video model in backend/models/video.go"
Task: "Implement Website model in backend/models/website.go"

# Launch T038-T041 together:
Task: "Unit tests for Video model in backend/models/video_test.go"
Task: "Unit tests for Website model in backend/models/website_test.go"
Task: "Unit tests for download logic in backend/downloader/"
Task: "Unit tests for DRM handling in backend/drm/"
```

## Notes
- [P] tasks = different files, no dependencies
- Verify tests fail before implementing
- Commit after each task
- Avoid: vague tasks, same file conflicts

## Validation Checklist
*GATE: Checked by main() before returning*

- [X] All contracts have corresponding tests
- [X] All entities have model tasks
- [X] All tests come before implementation
- [X] Parallel tasks truly independent
- [X] Each task specifies exact file path
- [X] No task modifies same file as another [P] task