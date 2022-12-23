package tcp

import (
	"context"
	"net"
	"syscall"
	"testing"
	"time"
)

func MockConnWithTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	d := net.Dialer{
		Control: func(_, addr string, _ syscall.RawConn) error {
			return &net.DNSError{
				Err:         "connection timed out",
				Name:        addr,
				Server:      "127.0.0.1",
				IsTimeout:   true,
				IsTemporary: true,
			}
		},
		Timeout: timeout,
	}
	return d.Dial(network, address)
}

func TimeOutConn(t *testing.T) {
	conn, err := MockConnWithTimeout("tcp", "10.0.0.1:http", 35*time.Second)
	if err == nil {
		err := conn.Close()
		if err != nil {
			return
		}
		t.Fatal("connection did not time out")
	}

	netErr, ok := err.(net.Error)
	if !ok {
		t.Fatal(err)
	}
	if !netErr.Timeout() {
		t.Fatal("error is not a timeout")
	}
}

func TimeOutConnUsingContext(t *testing.T) {
	deadlineTime := time.Now().Add(5 * time.Second)
	deadlineContext, cancel := context.WithDeadline(context.Background(), deadlineTime)
	defer cancel()

	var dialer net.Dialer
	dialer.Control = func(_, _ string, _ syscall.RawConn) error {
		time.Sleep(5*time.Second + time.Millisecond)
		return nil
	}

	conn, err := dialer.DialContext(deadlineContext, "tcp", "10.0.0.0:80")
	if err == nil {
		err := conn.Close()
		if err != nil {
			return
		}
		t.Fatal("connection did not time out")
	}

	netErr, ok := err.(net.Error)
	if !ok {
		t.Error(err)
	} else {
		if !netErr.Timeout() {
			t.Errorf("error is not a timeout: %v", err)
		}
	}

	if deadlineContext.Err() != context.DeadlineExceeded {
		t.Errorf("expected deadline exceeded; actual: %v", deadlineContext.Err())
	}
}
