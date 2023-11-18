
package selectRacer

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

//  net/http/httptest enables users to easily create a mock HTTP server.

func TestRacer(t *testing.T) {
	t.Run("compares speed of servers , returning url of the fastest ones ", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer slowServer.Close()
		defer fastServer.Close()
		fastURL := fastServer.URL
		slowURL := slowServer.URL
		want := fastURL
		got, _ := Racer(fastURL, slowURL)
		if got != want {
			t.Errorf("Got %q wanted %q", got, want)
		}
		slowServer.Close()
		fastServer.Close()
	})
	t.Run("returns an error if server does not respond withing secified time", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)
		

		defer server.Close()

		_, err := configurableRacer(server.URL, server.URL , 20*time.Millisecond)
		if err == nil {
			t.Error("Expected error but did not get one")
		}
		fmt.Println(err)
	})

}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
