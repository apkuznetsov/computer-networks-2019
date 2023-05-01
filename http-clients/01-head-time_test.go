package httpclient

import (
	"net/http"
	"testing"
	"time"
)

func TestHeadTime(t *testing.T) {
	now := time.Now().Round(time.Second)
	resp, err := http.Head("https://www.time.gov/")
	if err != nil {
		t.Fatal(err)
	}
	_ = resp.Body.Close()

	date := resp.Header.Get("Date")
	if date == "" {
		t.Fatal("no Date header received from time.gov")
	}

	dt, err := time.Parse(time.RFC1123, date)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("time.gov: %s (skew %s)", dt, now.Sub(dt))
}