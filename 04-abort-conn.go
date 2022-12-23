package tcp

import (
	"context"
	"log"
	"net"
	"syscall"
	"time"
)

func AbortConn() {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan struct{})

	go func() {
		defer func() { done <- struct{}{} }()

		var dialer net.Dialer
		dialer.Control = func(_, _ string, _ syscall.RawConn) error {
			time.Sleep(time.Second)
			return nil
		}

		conn, err := dialer.DialContext(ctx, "tcp", "10.0.0.1:80")
		if err != nil {
			log.Println(err)
			return
		}

		err = conn.Close()
		if err != nil {
			return
		}
		log.Fatal("connection did not time out")
	}()

	cancel()
	<-done
	if ctx.Err() != context.Canceled {
		log.Fatalf("expected canceled context; actual: %q", ctx.Err())
	}
}
