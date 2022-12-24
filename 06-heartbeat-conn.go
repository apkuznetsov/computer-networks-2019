package tcp

import (
	"context"
	"fmt"
	"io"
	"time"
)

const defaultPingInterval = 30 * time.Second

func Heartbeat(ctx context.Context, w io.Writer, reset <-chan time.Duration) {
	var interval time.Duration
	select {
	case <-ctx.Done():
		return
	case interval = <-reset: // pulled initial interval off reset channel
	default:
	}

	if interval <= 0 {
		interval = defaultPingInterval
	}
	timer := time.NewTimer(interval)
	defer func() {
		if !timer.Stop() {
			<-timer.C
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		case newInterval := <-reset:
			if !timer.Stop() {
				<-timer.C
			}
			if newInterval > 0 {
				interval = newInterval
			}
		case <-timer.C:
			if _, err := w.Write([]byte("ping")); err != nil {
				// track and act on consecutive timeouts here
				return
			}
		}
		_ = timer.Reset(interval)
	}
}

func HeartbeatConn() {
	r, w := io.Pipe()

	resetTimer := make(chan time.Duration, 1)
	resetTimer <- time.Second // initial ping interval

	ctx, cancelCtx := context.WithCancel(context.Background())
	done := make(chan struct{})

	go func() {
		Heartbeat(ctx, w, resetTimer)
		close(done)
	}()

	funcReceivePing := func(interval time.Duration, r io.Reader) {
		if interval >= 0 {
			fmt.Printf("resetting timer (%s)\n", interval)
			resetTimer <- interval
		}

		startedAt := time.Now()

		buf := make([]byte, 1024)
		n, err := r.Read(buf)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("received %q (%s)\n",
			buf[:n], time.Since(startedAt).Round(100*time.Millisecond))
	}

	// testing
	for i, v := range []int64{0, 200, 300, 0, -1, -1, -1} {
		fmt.Printf("Run %d:\n", i+1)
		funcReceivePing(time.Duration(v)*time.Millisecond, r)
	}

	cancelCtx()
	<-done
}
