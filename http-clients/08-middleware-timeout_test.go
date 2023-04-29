package httpclient

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestTimeoutMiddleware(t *testing.T) {
	handlerTimeout := http.TimeoutHandler(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusNoContent)
			time.Sleep(time.Minute)
		}),
		time.Second,
		"Timed out while reading response",
	)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://test/", nil)

	handlerTimeout.ServeHTTP(w, r)

	resp := w.Result()
	if resp.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("unexpected status code: %q", resp.Status)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	_ = resp.Body.Close()
	if actual := string(b); actual != "Timed out while reading response" {
		t.Logf("unexpected body: %q", actual)
	}
}
