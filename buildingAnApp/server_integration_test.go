package poker_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"hello/buildingAnApp"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	// store := NewInMemoryPlayerStore()
	database, cleanDatabase := poker.CreateTempFile(t, `[]`)
	defer cleanDatabase()
	// store := &FileSystemPlayerStore{database}
	store, err := poker.NewFileSystemPlayerStore(database)
	poker.AssertNoError(t, err)

	server, err := poker.NewPlayerServer(store,dummyGame)

	if err != nil {
		t.Error(err)
	}

	player := "virat"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetScoreRequest(player))
		poker.AssertStatus(t, response.Code, http.StatusOK)

		poker.AssertResponseBody(t, response.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newLeagueRequest())
		poker.AssertStatus(t, response.Code, http.StatusOK)

		got := poker.GetLeagueFromResponse(t, response.Body)
		want := []poker.Player{
			{"virat", 3},
		}
		poker.AssertLeague(t, got, want)
	})
}
