package selectTutorial

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("tests two servers and returns the faster one",func(t *testing.T) {
		slowServer := makeDelayedSerevr(20 * time.Millisecond)
		fastServer := makeDelayedSerevr(0 * time.Millisecond)
		
		defer slowServer.Close()
		defer fastServer.Close()
		
		slowURL := slowServer.URL
		fastURL := fastServer.URL
	
		want := fastURL
		got,err:= Racer(slowURL, fastURL)

		if(err!=nil){
			t.Error("did not expect an error but got one")
		}
	
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("returns an error if servers take more than specified time to respond",func(t *testing.T) {
		slowServer := makeDelayedSerevr(30 * time.Millisecond)
		fastServer := makeDelayedSerevr(25 * time.Millisecond)
		
		defer slowServer.Close()
		defer fastServer.Close()
		
		_,err:=ConfigurableRacer(slowServer.URL,fastServer.URL,22*time.Millisecond)
		// _,err:=Racer(slowServer.URL,fastServer.URL)

		if(err==nil){
			t.Error("expected an error but did not recieve one")
		}
	})


}

func makeDelayedSerevr(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))

}
