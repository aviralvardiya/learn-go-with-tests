package poker_test

import (
	// "encoding/json"
	"fmt"
	"strings"
	"time"

	// "io"
	"net/http"
	"net/http/httptest"

	"hello/buildingAnApp"

	// "reflect"
	"testing"

	"github.com/gorilla/websocket"
)

// type StubPlayerStore struct {
// 	scores   map[string]int
// 	winCalls []string
// 	league   League
// }

// func (s *StubPlayerStore) GetPlayerScore(name string) int {
// 	score := s.scores[name]
// 	return score
// }

// func (s *StubPlayerStore) RecordWin(name string) {
// 	s.winCalls = append(s.winCalls, name)
// }

// func (s *StubPlayerStore) GetLeague() League {
// 	return s.league
// }

func TestGETPlayers(t *testing.T) {
	store := poker.StubPlayerStore{
		map[string]int{
			"virat": 20,
			"rohit": 10,
		},
		nil,
		nil,
	}
	server,err := poker.NewPlayerServer(&store,dummyGame)
	if(err!=nil){
		t.Error(err)
	}

	t.Run("returns virat's score", func(t *testing.T) {
		request := newGetScoreRequest("virat")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns rohit's score", func(t *testing.T) {
		request := newGetScoreRequest("rohit")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusOK)
		poker.AssertResponseBody(t, response.Body.String(), "10")
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusNotFound)
	})
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

// func assertStatus(t testing.TB, got, want int) {
// 	t.Helper()
// 	if got != want {
// 		t.Errorf("did not get correct status, got %d, want %d", got, want)
// 	}
// }

// func assertResponseBody(t testing.TB, got, want string) {
// 	t.Helper()
// 	if got != want {
// 		t.Errorf("response body is wrong, got %q want %q", got, want)
// 	}
// }

func TestStoreWins(t *testing.T) {
	store := poker.StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server,err := poker.NewPlayerServer(&store,dummyGame)
	if(err!=nil){
		t.Error(err)
	}

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "virat"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusAccepted)

		poker.AssertPlayerWin(t, &store, player)
	})
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func TestLeague(t *testing.T) {

	t.Run("it returns the league table as JSON", func(t *testing.T) {
		wantedLeague := []poker.Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		store := poker.StubPlayerStore{nil, nil, wantedLeague}
		server,err := poker.NewPlayerServer(&store,dummyGame)
		if(err!=nil){
			t.Error(err)
		}

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := poker.GetLeagueFromResponse(t, response.Body)
		poker.AssertStatus(t, response.Code, http.StatusOK)
		poker.AssertLeague(t, got, wantedLeague)

		poker.AssertContentType(t, response, poker.JsonContentType)
	})

}

// func getLeagueFromResponse(t testing.TB, body io.Reader) (league []Player) {
// 	t.Helper()
// 	err := json.NewDecoder(body).Decode(&league)

// 	if err != nil {
// 		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
// 	}

// 	return
// }

// func assertLeague(t testing.TB, got, want []Player) {
// 	t.Helper()
// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %v want %v", got, want)
// 	}
// }

// func assertContentType(t testing.TB, response *httptest.ResponseRecorder, want string) {
// 	t.Helper()
// 	if response.Result().Header.Get("content-type") != want {
// 		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
// 	}
// }

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

// func assertPlayerWin(t testing.TB, store *StubPlayerStore, winner string) {
// 	t.Helper()

// 	if len(store.winCalls) != 1 {
// 		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
// 	}

// 	if store.winCalls[0] != winner {
// 		t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], winner)
// 	}
// }

func TestGame(t *testing.T) {
	t.Run("GET /game returns 200", func(t *testing.T) {
		server,err := poker.NewPlayerServer(&poker.StubPlayerStore{},dummyGame)
		if(err!=nil){
			t.Error(err)
		}

		request, _ := newGameRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		poker.AssertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("start a game with 3 players, send some blind alerts down WS and declare Ruth the winner", func(t *testing.T) {
		wantedBlindAlert := "Blind is 100"
		winner := "Ruth"
	
		game := &GameSpy{BlindAlert: []byte(wantedBlindAlert)}
		server := httptest.NewServer(mustMakePlayerServer(t, dummyPlayerStore, game))
		ws := mustDialWS(t, "ws"+strings.TrimPrefix(server.URL, "http")+"/ws")
	
		defer server.Close()
		defer ws.Close()
	
		writeWSMessage(t, ws, "3")
		writeWSMessage(t, ws, winner)

		tenMS:=10*time.Millisecond
	
		time.Sleep(tenMS)
	
		assertGameStartedWith(t, game, 3)
		assertFinishCalledWith(t, game, winner)
		within(t, tenMS, func() { assertWebsocketGotMsg(t, ws, wantedBlindAlert) })
	})

	

	
}

var (
	dummyGame = &GameSpy{}
)

func newGameRequest() (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, "/game", nil)
	return req, err
}

func mustMakePlayerServer(t *testing.T, store poker.PlayerStore, game poker.Game) *poker.PlayerServer {
	server, err := poker.NewPlayerServer(store,game)
	if err != nil {
		t.Fatal("problem creating player server", err)
	}
	return server
}

func mustDialWS(t *testing.T, url string) *websocket.Conn {
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)

	if err != nil {
		t.Fatalf("could not open a ws connection on %s %v", url, err)
	}

	return ws
}

func writeWSMessage(t testing.TB, conn *websocket.Conn, message string) {
	t.Helper()
	if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
		t.Fatalf("could not send message over ws connection %v", err)
	}
}

func within(t testing.TB, d time.Duration, assert func()) {
	t.Helper()

	done := make(chan struct{}, 1)

	go func() {
		assert()
		done <- struct{}{}
	}()

	select {
	case <-time.After(d):
		t.Error("timed out")
	case <-done:
	}
}

func assertWebsocketGotMsg(t *testing.T, ws *websocket.Conn, want string) {
	_, msg, _ := ws.ReadMessage()
	if string(msg) != want {
		t.Errorf(`got "%s", want "%s"`, string(msg), want)
	}
}