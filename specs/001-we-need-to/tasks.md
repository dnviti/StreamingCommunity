# Tasks for Refactor StreamingCommunity

## Phase 1: Core Backend Setup (Go)

*   **T001**: [P] Set up the Go backend project with Gin.
*   **T002**: [P] Implement the `Video` and `Website` data models.
*   **T003**: [P] Implement the JSON file storage for the data models.
*   **T004**: Implement the `/downloads` API endpoints.
*   **T005**: [P] Implement HLS downloader.
*   **T006**: [P] Implement DASH downloader.
*   **T007**: [P] Implement MP4 downloader.
*   **T008**: [P] Implement Torrent downloader (integration with a client like qBittorrent).
*   **T009**: Implement Widevine DRM handling.

## Phase 2: Frontend Setup (Svelte)

*   **T010**: [P] Set up the Svelte frontend project with SvelteKit.
*   **T011**: [P] Create the UI for submitting download requests.
*   **T012**: [P] Create the UI for viewing download status.
*   **T013**: Integrate the frontend with the backend API.

## Phase 3: Website Implementation - Altadefinizione

*   **T014**: [P] Implement search functionality for Altadefinizione.
*   **T015**: [P] Implement scraping logic for films on Altadefinizione.
*   **T016**: [P] Implement scraping logic for series on Altadefinizione.
*   **T017**: [P] Implement Supervideo player handling.
*   **T018**: Integrate Altadefinizione with the HLS downloader.

## Phase 4: Website Implementation - AnimeUnity

*   **T019**: [P] Implement search functionality for AnimeUnity.
*   **T020**: [P] Implement scraping logic for films on AnimeUnity.
*   **T021**: [P] Implement scraping logic for series on AnimeUnity.
*   **T022**: [P] Implement Vixcloud player handling.
*   **T023**: Integrate AnimeUnity with the MP4 downloader.

## Phase 5: Website Implementation - AnimeWorld

*   **T024**: [P] Implement search functionality for AnimeWorld.
*   **T025**: [P] Implement scraping logic for films on AnimeWorld.
*   **T026**: [P] Implement scraping logic for series on AnimeWorld.
*   **T027**: [P] Implement Sweetpixel player handling.
*   **T028**: Integrate AnimeWorld with the MP4 downloader.

## Phase 6: Website Implementation - Crunchyroll

*   **T029**: [P] Implement search functionality for Crunchyroll.
*   **T030**: [P] Implement scraping logic for films on Crunchyroll.
*   **T031**: [P] Implement scraping logic for series on Crunchyroll.
*   **T032**: Implement Crunchyroll license acquisition.
*   **T033**: Integrate Crunchyroll with the DASH downloader and Widevine DRM.

## Phase 7: Website Implementation - Guardaserie

*   **T034**: [P] Implement search functionality for Guardaserie.
*   **T035**: [P] Implement scraping logic for series on Guardaserie.
*   **T036**: [P] Implement Supervideo player handling.
*   **T037**: Integrate Guardaserie with the HLS downloader.

## Phase 8: Website Implementation - Mediaset Infinity

*   **T038**: [P] Implement search functionality for Mediaset Infinity.
*   **T039**: [P] Implement scraping logic for films on Mediaset Infinity.
*   **T040**: [P] Implement scraping logic for series on Mediaset Infinity.
*   **T041**: Implement Mediaset Infinity license acquisition.
*   **T042**: Integrate Mediaset Infinity with the DASH downloader and Widevine DRM.

## Phase 9: Website Implementation - RaiPlay

*   **T043**: [P] Implement search functionality for RaiPlay.
*   **T044**: [P] Implement scraping logic for films on RaiPlay.
*   **T045**: [P] Implement scraping logic for series on RaiPlay.
*   **T046**: Implement RaiPlay license acquisition.
*   **T047**: Integrate RaiPlay with the HLS/DASH downloader and Widevine DRM.

## Phase 10: Website Implementation - StreamingCommunity

*   **T048**: [P] Implement search functionality for StreamingCommunity.
*   **T049**: [P] Implement scraping logic for films on StreamingCommunity.
*   **T050**: [P] Implement scraping logic for series on StreamingCommunity.
*   **T051**: [P] Implement Vixcloud player handling.
*   **T052**: Integrate StreamingCommunity with the HLS downloader.

## Phase 11: Website Implementation - StreamingWatch

*   **T053**: [P] Implement search functionality for StreamingWatch.
*   **T054**: [P] Implement scraping logic for films on StreamingWatch.
*   **T055**: [P] Implement scraping logic for series on StreamingWatch.
*   **T056**: [P] Implement HDPlayer player handling.
*   **T057**: Integrate StreamingWatch with the HLS downloader.

## Phase 12: Testing and Polish

*   **T058**: [P] Write unit and integration tests for the backend.
*   **T059**: [P] Write tests for the frontend.
*   **T060**: Implement build script to copy frontend build to `backend/web`.
*   **T061**: Update `README.md` with new build and run instructions.
