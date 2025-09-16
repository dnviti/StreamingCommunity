# Tasks: Refactor StreamingCommunity

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
- **Single project**: `src/`, `tests/` at repository root
- **Web app**: `backend/src/`, `frontend/src/`
- **Mobile**: `api/src/`, `ios/src/` or `android/src/`
- Paths shown below assume single project - adjust based on plan.md structure

## Phase 3.1: Setup
- [ ] T001 Create Go backend project structure (`backend/`)
- [ ] T002 Initialize Go backend with Gin and other dependencies (`backend/go.mod`, `backend/main.go`)
- [ ] T003 [P] Configure Go linting and formatting (`backend/.golangci.yml` or similar)
- [ ] T004 Create SvelteKit frontend project structure (`frontend/`)
- [ ] T005 Initialize SvelteKit frontend with dependencies (`frontend/package.json`, `frontend/svelte.config.js`)
- [ ] T006 [P] Configure SvelteKit linting and formatting (`frontend/.eslintrc.cjs`, `frontend/.prettierrc`)
- [ ] T007 Configure Go testing (`backend/tests/`)
- [ ] T008 Configure SvelteKit testing (Vitest/Playwright) (`frontend/vitest.config.ts`, `frontend/playwright.config.ts`)

## Phase 3.2: Tests First (TDD) ⚠️ MUST COMPLETE BEFORE 3.3
**CRITICAL: These tests MUST be written and MUST FAIL before ANY implementation**
- [ ] T009 [P] Contract test API endpoints from `specs/001-we-need-to/contracts/openapi.yaml` (`backend/tests/api_contract_test.go`)
- [ ] T010 [P] Integration test for `altadefinizione` scraper (`backend/tests/scraper_altadefinizione_test.go`)
- [ ] T011 [P] Integration test for HLS downloader (`backend/tests/downloader_hls_test.go`)
- [ ] T012 [P] E2E test for main navigation in SvelteKit (`frontend/tests/e2e/navigation.test.ts`)
- [ ] T013 [P] E2E test for video display in SvelteKit (`frontend/tests/e2e/video_display.test.ts`)

## Phase 3.3: Core Implementation (ONLY after tests are failing)
- [ ] T014 [P] Implement `Video` model (`backend/models/video.go`)
- [ ] T015 [P] Implement `Website` model (`backend/models/website.go`)
- [ ] T016 Implement JSON storage for `Video` and `Website` (`backend/storage/json_storage.go`)
- [ ] T017 Implement base scraping logic using `colly` (`backend/scrapers/base_scraper.go`)
- [ ] T018 Implement `altadefinizione` scraper (`backend/scrapers/altadefinizione/altadefinizione.go`)
- [ ] T019 Implement `animeunity` scraper (`backend/scrapers/animeunity/animeunity.go`)
- [ ] T020 Implement HLS downloader (`backend/downloader/hls.go`)
- [ ] T021 Implement DASH downloader (`backend/downloader/dash.go`)
- [ ] T022 Implement MP4 downloader (`backend/downloader/mp4.go`)
- [ ] T023 Implement Torrent downloader using `anacrolix/torrent` (`backend/downloader/torrent.go`)
- [ ] T024 Implement Widevine DRM handling (Cgo wrapper or external Python script) (`backend/drm/widevine.go`)
- [ ] T025 Integrate `goja` for JavaScript execution (`backend/utils/js_executor.go`)
- [ ] T026 Implement Gin web server setup (`backend/main.go`)
- [ ] T027 Implement API endpoints based on `openapi.yaml` (`backend/api/video_handlers.go`, `backend/api/website_handlers.go`)
- [ ] T028 Set up basic SvelteKit application layout (`frontend/src/routes/+layout.svelte`)
- [ ] T029 Implement pages for displaying videos (`frontend/src/routes/videos/+page.svelte`)
- [ ] T030 Implement pages for managing websites (`frontend/src/routes/websites/+page.svelte`)
- [ ] T031 Implement search functionality component (`frontend/src/lib/components/Search.svelte`)
- [ ] T032 Implement video playback component (`frontend/src/lib/components/VideoPlayer.svelte`)

## Phase 3.4: Integration
- [ ] T033 Connect SvelteKit frontend to Go backend API (`frontend/src/lib/services/api.ts`)
- [ ] T034 Implement logging for backend (`backend/utils/logger.go`)
- [ ] T035 Implement error handling for backend (`backend/utils/errors.go`)

## Phase 3.5: Polish
- [ ] T036 [P] Unit tests for `Video` model (`backend/models/video_test.go`)
- [ ] T037 [P] Unit tests for `Website` model (`backend/models/website_test.go`)
- [ ] T038 [P] Unit tests for JSON storage (`backend/storage/json_storage_test.go`)
- [ ] T039 Performance optimization for scrapers (`backend/scrapers/`)
- [ ] T040 Performance optimization for downloaders (`backend/downloader/`)
- [ ] T041 [P] Add documentation for API endpoints (`docs/api.md`)
- [ ] T042 [P] Add documentation for frontend components (`frontend/src/lib/components/README.md`)

## Dependencies
- Setup (T001-T008) before Tests (T009-T013)
- Tests (T009-T013) before Core Implementation (T014-T032)
- Models (T014-T015) before Storage (T016) and API handlers (T027)
- Base Scraper (T017) before specific scrapers (T018-T019)
- Core Implementation (T014-T032) before Integration (T033-T035)
- Integration (T033-T035) before Polish (T036-T042)
- T026 (Gin server setup) before T027 (API endpoints)
- T028 (SvelteKit layout) before T029-T032 (SvelteKit pages/components)

## Parallel Example
```
# Launch T003 and T006 together:
Task: "Configure Go linting and formatting (`backend/.golangci.yml` or similar)"
Task: "Configure SvelteKit linting and formatting (`frontend/.eslintrc.cjs`, `frontend/.prettierrc`)"

# Launch T009-T013 together:
Task: "Contract test API endpoints from `specs/001-we-need-to/contracts/openapi.yaml` (`backend/tests/api_contract_test.go`)"
Task: "Integration test for `altadefinizione` scraper (`backend/tests/scraper_altadefinizione_test.go`)"
Task: "Integration test for HLS downloader (`backend/tests/downloader_hls_test.go`)"
Task: "E2E test for main navigation in SvelteKit (`frontend/tests/e2e/navigation.test.ts`)"
Task: "E2E test for video display in SvelteKit (`frontend/tests/e2e/video_display.test.ts`)"

# Launch T014 and T015 together:
Task: "Implement `Video` model (`backend/models/video.go`)"
Task: "Implement `Website` model (`backend/models/website.go`)"

# Launch T036-T038, T041-T042 together:
Task: "Unit tests for `Video` model (`backend/models/video_test.go`)"
Task: "Unit tests for `Website` model (`backend/models/website_test.go`)"
Task: "Unit tests for JSON storage (`backend/storage/json_storage_test.go`)"
Task: "Add documentation for API endpoints (`docs/api.md`)"
Task: "Add documentation for frontend components (`frontend/src/lib/components/README.md`)"
```

## Notes
- [P] tasks = different files, no dependencies
- Verify tests fail before implementing
- Commit after each task
- Avoid: vague tasks, same file conflicts

## Task Generation Rules
*Applied during main() execution*

1. **From Contracts**:
   - Each contract file → contract test task [P]
   - Each endpoint → implementation task
   
2. **From Data Model**:
   - Each entity → model creation task [P]
   - Relationships → service layer tasks
   
3. **From User Stories**:
   - Each story → integration test [P]
   - Quickstart scenarios → validation tasks

4. **Ordering**:
   - Setup → Tests → Models → Services → Endpoints → Polish
   - Dependencies block parallel execution

## Validation Checklist
*GATE: Checked by main() before returning*

- [ ] All contracts have corresponding tests
- [ ] All entities have model tasks
- [ ] All tests come before implementation
- [ ] Parallel tasks truly independent
- [ ] Each task specifies exact file path
- [ ] No task modifies same file as another [P] task
