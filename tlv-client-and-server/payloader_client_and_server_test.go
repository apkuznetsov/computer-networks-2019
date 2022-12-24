package tlv

import (
	"bytes"
	"encoding/binary"
	"net"
	"reflect"
	"testing"
)

func TestPayloadClientAndServer(t *testing.T) {
	b1 := Binary("Clear is better than clever.")
	b2 := Binary("Don't panic.")
	s1 := String("Errors are values.")
	payloads := []Payloader{&b1, &s1, &b2}

	// server running
	tcpListener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	// server accepting
	go func() {
		connAccepted, err := tcpListener.Accept()
		if err != nil {
			t.Error(err)
			return
		}
		defer func(connAccepted net.Conn) {
			_ = connAccepted.Close()
		}(connAccepted)

		// server writing payload to client
		for _, p := range payloads {
			_, err = p.WriteTo(connAccepted)
			if err != nil {
				t.Error(err)
				break
			}
		}
	}()

	// client connecting
	clientConn, err := net.Dial("tcp", tcpListener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}
	defer func(clientConn net.Conn) {
		_ = clientConn.Close()
	}(clientConn)

	// client reading payloads
	for i := 0; i < len(payloads); i++ {
		actual, err := decode(clientConn)
		if err != nil {
			t.Fatal(err)
		}

		if expected := payloads[i]; !reflect.DeepEqual(expected, actual) {
			t.Errorf("value mismatch: %v != %v", expected, actual)
			continue
		}
		t.Logf("[%T] %[1]q", actual)
	}
}

func TestMaxPayloadSize(t *testing.T) {
	buf := new(bytes.Buffer)
	err := buf.WriteByte(BinaryType)
	if err != nil {
		t.Fatal(err)
	}

	err = binary.Write(buf, binary.BigEndian, uint32(1<<30)) // 1 GB
	if err != nil {
		t.Fatal(err)
	}

	var b Binary
	_, err = b.ReadFrom(buf)
	if err != ErrMaxPayloadSize {
		t.Fatalf("expected ErrMaxPayloadSize; actual: %v", err)
	}
}
