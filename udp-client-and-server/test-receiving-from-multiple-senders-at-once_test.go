package udp

import (
	"bytes"
	"context"
	"net"
	"testing"
)

func TestEchoUdpServerUsingMultipleSenders(t *testing.T) {
	// server
	ctx, cancelCtx := context.WithCancel(context.Background())
	serverAddr, err := startEchoUdpServer(ctx, "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	defer cancelCtx()

	// client
	clientListener, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}
	defer func() { _ = clientListener.Close() }()

	// interloper
	interloperListener, err := net.ListenPacket("udp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	// interloper interrupting client
	interruptMessage := []byte("pardon me")
	n, err := interloperListener.WriteTo(interruptMessage, clientListener.LocalAddr())
	if err != nil {
		t.Fatal(err)
	}
	_ = interloperListener.Close()

	// assert interrupt
	if l := len(interruptMessage); l != n {
		t.Fatalf("wrote %d bytes of %d", n, l)
	}

	// client writing to server
	pingMessage := []byte("ping")
	_, err = clientListener.WriteTo(pingMessage, serverAddr)
	if err != nil {
		t.Fatal(err)
	}

	// client reading
	buf := make([]byte, 1024)
	n, addr, err := clientListener.ReadFrom(buf)
	if err != nil {
		t.Fatal(err)
	}

	// assert that first message is from interloper
	if !bytes.Equal(interruptMessage, buf[:n]) {
		t.Errorf("expected reply %q; actual reply %q", interruptMessage, buf[:n])
	}
	if addr.String() != interloperListener.LocalAddr().String() {
		t.Errorf("expected message from %q; actual sender is %q",
			interloperListener.LocalAddr(), addr)
	}

	// client reading
	n, addr, err = clientListener.ReadFrom(buf)
	if err != nil {
		t.Fatal(err)
	}

	// assert that second message is from server
	if !bytes.Equal(pingMessage, buf[:n]) {
		t.Errorf("expected reply %q; actual reply %q", pingMessage, buf[:n])
	}
	if addr.String() != serverAddr.String() {
		t.Errorf("expected message from %q; actual sender is %q",
			serverAddr, addr)
	}
}
