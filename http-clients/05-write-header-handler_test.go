package httpclient

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWriteHeaderHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "http://test", nil)
	handler := func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Bad request"))
		w.WriteHeader(http.StatusBadRequest)
	}
	handler(w, r)
	t.Logf("Response status: %q", w.Result().Status)

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "http://test", nil)
	handler = func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Bad request"))
	}
	handler(w, r)
	t.Logf("Response status: %q", w.Result().Status)
}
