package udp

import (
	"bytes"
	"context"
	"net"
	"testing"
)

func TestEchoUdpServer(t *testing.T) {
	ctx, cancelCtx := context.WithCancel(context.Background())
	serverAddr, err := startEchoUdpServer(ctx, "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	defer cancelCtx()

	clientListener, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = clientListener.Close() }()

	message := []byte("ping")
	_, err = clientListener.WriteTo(message, serverAddr)
	if err != nil {
		t.Fatal(err)
	}

	buf := make([]byte, 1024)
	n, addr, err := clientListener.ReadFrom(buf)
	if err != nil {
		t.Fatal(err)
	}

	// assert
	if addr.String() != serverAddr.String() {
		t.Fatalf("received reply from %q instead of %q", addr, serverAddr)
	}
	if !bytes.Equal(message, buf[:n]) {
		t.Errorf("expected reply %q; actual reply %q", message, buf[:n])
	}
}
