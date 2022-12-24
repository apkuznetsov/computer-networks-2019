package tcp

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
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

func HeartbeatDeadline() {
	tcpListener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		log.Fatal(err)
	}

	begin := time.Now()

	done := make(chan struct{})
	go func() {
		defer func() { close(done) }()

		// server starting listening
		connAccepted, err := tcpListener.Accept()
		if err != nil {
			log.Println(err)
			return
		}

		// server starting heartbeat
		ctx, cancelCtx := context.WithCancel(context.Background())
		defer func() {
			cancelCtx()
			_ = connAccepted.Close()
		}()
		resetTimer := make(chan time.Duration, 1)
		resetTimer <- time.Second
		go Heartbeat(ctx, connAccepted, resetTimer)

		// server reading data
		err = connAccepted.SetDeadline(time.Now().Add(15 * time.Second))
		if err != nil {
			log.Println(err)
			return
		}
		buf := make([]byte, 1024)
		for {
			n, err := connAccepted.Read(buf)
			if err != nil {
				return
			}
			log.Printf("[%s] %s",
				time.Since(begin).Truncate(time.Second), buf[:n])

			// read the data, reset the timer
			resetTimer <- 0
			err = connAccepted.SetDeadline(time.Now().Add(5 * time.Second))
			if err != nil {
				log.Println(err)
				return
			}
		}
	}()

	// client connecting
	conn, err := net.Dial("tcp", tcpListener.Addr().String())
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	// client reading ping
	buf := make([]byte, 1024)
	for i := 0; i < 4; i++ { // read four pings
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("[%s] %s\n", time.Since(begin).Truncate(time.Second), buf[:n])
	}

	// client writing pong
	_, err = conn.Write([]byte("PONG")) // reset the ping timer
	if err != nil {
		log.Fatal(err)
	}

	// client reading ping
	for i := 0; i < 4; i++ { // read four pings
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
		log.Printf("[%s] %s\n", time.Since(begin).Truncate(time.Second), buf[:n])
	}

	<-done
	end := time.Since(begin).Truncate(time.Second)

	log.Printf("[%s] done\n", end)
	if end != 9*time.Second {
		log.Fatalf("expected EOF at 9 seconds; actual %s", end)
	}
}
