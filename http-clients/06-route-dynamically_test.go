package httpclient

import (
	"io"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

type Methods map[string]http.Handler

func (h Methods) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer func(r io.ReadCloser) {
		_, _ = io.Copy(ioutil.Discard, r)
		_ = r.Close()
	}(r.Body)

	if handler, ok := h[r.Method]; ok {
		if handler == nil {
			http.Error(w, "Internal server error",
				http.StatusInternalServerError)
		} else {
			handler.ServeHTTP(w, r)
		}
		return
	}

	w.Header().Add("Allow", h.allowedMethods())
	if r.Method != http.MethodOptions {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h Methods) allowedMethods() string {
	a := make([]string, 0, len(h))
	for m := range h {
		a = append(a, m)
	}

	sort.Strings(a)
	return strings.Join(a, ", ")
}
