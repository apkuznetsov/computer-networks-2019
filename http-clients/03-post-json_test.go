package httpclient

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type User struct {
	First string
	Last  string
}

func handlePostUser(t *testing.T) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func(r io.ReadCloser) {
			_, _ = io.Copy(ioutil.Discard, r)
			_ = r.Close()
		}(r.Body)

		if r.Method != http.MethodPost {
			http.Error(w, "", http.StatusMethodNotAllowed)
			return
		}

		var u User
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			t.Error(err)
			http.Error(w, "Decode Failed", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusAccepted)
	}
}

func TestPostUser(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(handlePostUser(t)))
	defer testServer.Close()

	resp, err := http.Get(testServer.URL)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Fatalf("expected status %d; actual status %d",
			http.StatusMethodNotAllowed, resp.StatusCode)
	}

	buf := new(bytes.Buffer)
	u := User{First: "Adam", Last: "Woodbeck"}
	err = json.NewEncoder(buf).Encode(&u)
	if err != nil {
		t.Fatal(err)
	}

	resp, err = http.Post(testServer.URL, "application/json", buf)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusAccepted {
		t.Fatalf("expected status %d; actual status %d",
			http.StatusAccepted, resp.StatusCode)

	}
	_ = resp.Body.Close()
}
