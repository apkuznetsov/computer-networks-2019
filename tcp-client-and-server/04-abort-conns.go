package tcp

import (
	"context"
	"log"
	"net"
	"sync"
	"syscall"
	"time"
)

func AbortConn() {
	ctx, cancelCtx := context.WithCancel(context.Background())
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

	cancelCtx()
	<-done
	if ctx.Err() != context.Canceled {
		log.Fatalf("expected canceled context; actual: %q", ctx.Err())
	}
}

func AbortConns() {
	tcpListener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		log.Fatal(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			return
		}
	}(tcpListener)

	// accept only a single connection
	go func() {
		acceptedConn, err := tcpListener.Accept()
		if err == nil {
			err := acceptedConn.Close()
			if err != nil {
				return
			}
		}
	}()

	funcDial := func(ctx context.Context, address string, response chan int, id int, wg *sync.WaitGroup) {
		defer wg.Done()

		var dialer net.Dialer
		conn, err := dialer.DialContext(ctx, "tcp", address)
		if err != nil {
			return
		}

		err = conn.Close()
		if err != nil {
			return
		}

		log.Printf("dialer #%d waiting\n", id)
		select {
		case <-ctx.Done():
		case response <- id:
		}
	}

	ctxDeadline, cancelCtx := context.WithDeadline(
		context.Background(),
		time.Now().Add(10*time.Second),
	)

	chanRes := make(chan int)
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go funcDial(ctxDeadline, tcpListener.Addr().String(), chanRes, i+1, &wg)
	}

	response := <-chanRes
	cancelCtx()
	wg.Wait()
	close(chanRes)

	if ctxDeadline.Err() != context.Canceled {
		log.Fatalf("expected canceled context; actual: %s",
			ctxDeadline.Err(),
		)
	}
	log.Printf("dialer %d retrieved the resource\n", response)
}
