package httpclient

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func blockIndefinitely(w http.ResponseWriter, r *http.Request) {
	select {}
}

func TestBlockIndefinitelyWithTimeout(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(blockIndefinitely))

	ctx, cancelCtx := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelCtx()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, testServer.URL, nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatal(err)
		}
		return
	}
	_ = resp.Body.Close()
}
