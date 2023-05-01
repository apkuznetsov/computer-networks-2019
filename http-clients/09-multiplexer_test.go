package httpclient

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func drainAndClose(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
			_, _ = io.Copy(ioutil.Discard, r.Body)
			_ = r.Body.Close()
		},
	)
}

func TestSimpleMux(t *testing.T) {
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})
	serveMux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Hello friend.")
	})
	serveMux.HandleFunc("/hello/there/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprint(w, "Why, hello there.")
	})

	mux := drainAndClose(serveMux)

	testCases := []struct {
		path     string
		response string
		code     int
	}{
		{"http://test/", "", http.StatusNoContent},
		{"http://test/hello", "Hello friend.", http.StatusOK},
		{"http://test/hello/there/", "Why, hello there.", http.StatusOK},
	}

	for i, c := range testCases {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, c.path, nil)
		mux.ServeHTTP(w, r)

		resp := w.Result()
		if actual := resp.StatusCode; c.code != actual {
			t.Errorf("%d: expected code %d; actual %d", i, c.code, actual)
		}

		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		_ = resp.Body.Close()
		if actual := string(b); c.response != actual {
			t.Errorf("%d: expected response %q; actual %q", i,
				c.response, actual)
		}
	}
}
