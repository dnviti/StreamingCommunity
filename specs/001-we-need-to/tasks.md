# Tasks for Refactor StreamingCommunity

## Phase 1: Core Backend Setup (Go)

*   **T001**: [x] Set up the Go backend project with Gin.
*   **T002**: [x] Implement the `Video` and `Website` data models.
*   **T003**: [x] Implement the JSON file storage for the data models.
*   **T004**: [x] Implement the `/downloads` API endpoints.
*   **T005**: [x] Implement HLS downloader.
*   **T006**: [x] Implement DASH downloader.
*   **T007**: [x] Implement MP4 downloader.
*   **T008**: [x] Implement Torrent downloader (integration with a client like qBittorrent).
*   **T009**: [x] Implement Widevine DRM handling.

## Phase 2: Frontend Setup (Svelte)

*   **T010**: [x] Set up the Svelte frontend project with SvelteKit.
*   **T011**: [x] Create the UI for submitting download requests.
*   **T012**: [x] Create the UI for viewing download status.
*   **T013**: [x] Integrate the frontend with the backend API.

## Phase 3: Website Implementation - Altadefinizione

*   **T014**: [x] Implement search functionality for Altadefinizione.
*   **T015**: [x] Implement scraping logic for films on Altadefinizione.
*   **T016**: [x] Implement scraping logic for series on Altadefinizione.
*   **T017**: [x] Implement Supervideo player handling.
*   **T018**: [x] Integrate Altadefinizione with the HLS downloader.

## Phase 4: Website Implementation - AnimeUnity

*   **T019**: [x] Implement search functionality for AnimeUnity.
*   **T020**: [x] Implement scraping logic for films on AnimeUnity.
*   **T021**: [x] Implement scraping logic for series on AnimeUnity.
*   **T022**: [x] Implement Vixcloud player handling.
*   **T023**: [x] Integrate AnimeUnity with the MP4 downloader.

## Phase 5: Website Implementation - AnimeWorld

*   **T024**: [x] Implement search functionality for AnimeWorld.
*   **T025**: [x] Implement scraping logic for films on AnimeWorld.
*   **T026**: [x] Implement scraping logic for series on AnimeWorld.
*   **T027**: [x] Implement Sweetpixel player handling.
*   **T028**: [x] Integrate AnimeWorld with the MP4 downloader.

## Phase 6: Website Implementation - Crunchyroll

*   **T029**: [x] Implement search functionality for Crunchyroll.
*   **T030**: [x] Implement scraping logic for films on Crunchyroll.
*   **T031**: [x] Implement scraping logic for series on Crunchyroll.
*   **T032**: [x] Implement Crunchyroll license acquisition.
*   **T033**: [x] Integrate Crunchyroll with the DASH downloader and Widevine DRM.

## Phase 7: Website Implementation - Guardaserie

*   **T034**: [x] Implement search functionality for Guardaserie.
*   **T035**: [x] Implement scraping logic for series on Guardaserie.
*   **T036**: [x] Implement Supervideo player handling.
*   **T037**: [x] Integrate Guardaserie with the HLS downloader.

## Phase 8: Website Implementation - Mediaset Infinity

*   **T038**: [x] Implement search functionality for Mediaset Infinity.
*   **T039**: [x] Implement scraping logic for films on Mediaset Infinity.
*   **T040**: [x] Implement scraping logic for series on Mediaset Infinity.
*   **T041**: [x] Implement Mediaset Infinity license acquisition.
*   **T042**: [x] Integrate Mediaset Infinity with the DASH downloader and Widevine DRM.

## Phase 9: Website Implementation - RaiPlay

*   **T043**: [x] Implement search functionality for RaiPlay.
*   **T044**: [x] Implement scraping logic for films on RaiPlay.
*   **T045**: [x] Implement scraping logic for series on RaiPlay.
*   **T046**: [x] Implement RaiPlay license acquisition.
*   **T047**: [x] Integrate RaiPlay with the HLS/DASH downloader and Widevine DRM.

## Phase 10: Website Implementation - StreamingCommunity

*   **T048**: [x] Implement search functionality for StreamingCommunity.
*   **T049**: [x] Implement scraping logic for films on StreamingCommunity.
*   **T050**: [x] Implement scraping logic for series on StreamingCommunity.
*   **T051**: [x] Implement Vixcloud player handling.
*   **T052**: [x] Integrate StreamingCommunity with the HLS downloader.

## Phase 11: Website Implementation - StreamingWatch

*   **T053**: [x] Implement search functionality for StreamingWatch.
*   **T054**: [x] Implement scraping logic for films on StreamingWatch.
*   **T055**: [x] Implement scraping logic for series on StreamingWatch.
*   **T056**: [x] Implement HDPlayer player handling.
*   **T057**: [x] Integrate StreamingWatch with the HLS downloader.

## Phase 12: Testing and Polish

*   **T058**: [x] Write unit and integration tests for the backend.
*   **T059**: [x] Write tests for the frontend.
*   **T060**: [x] Implement build script to copy frontend build to `backend/web`.
*   **T061**: [x] Update `README.md` with new build and run instructions.
