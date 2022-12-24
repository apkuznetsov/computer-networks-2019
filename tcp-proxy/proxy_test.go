package tcp_proxy

import (
	"io"
	"net"
	"sync"
	"testing"
)

func TestProxy(t *testing.T) {
	// server listening
	serverListener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	// server accepting
	go func() {
		defer wg.Done()
		for {
			connAccepted, err := serverListener.Accept()
			if err != nil {
				return
			}

			// server proxying
			go func(conn net.Conn) {
				defer func(conn net.Conn) {
					_ = conn.Close()
				}(conn)

				for {
					buf := make([]byte, 1024)
					n, err := conn.Read(buf)
					if err != nil {
						if err != io.EOF {
							t.Error(err)
						}
						return
					}

					switch msg := string(buf[:n]); msg {
					case "ping":
						_, err = conn.Write([]byte("pong"))
					default:
						_, err = conn.Write(buf[:n])
					}

					if err != nil {
						if err != io.EOF {
							t.Error(err)
						}
						return
					}
				}
			}(connAccepted)
		}
	}()

	// proxy listening
	proxyListener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		t.Fatal(err)
	}

	// proxying
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			fromClient, err := proxyListener.Accept()
			if err != nil {
				return
			}

			go func(from net.Conn) {
				defer func(from net.Conn) {
					_ = from.Close()
				}(from)

				toServer, err := net.Dial("tcp",
					serverListener.Addr().String())
				if err != nil {
					t.Error(err)
					return
				}
				defer func(toServer net.Conn) {
					_ = toServer.Close()
				}(toServer)

				err = proxy(from, toServer)
				if err != nil && err != io.EOF {
					t.Error(err)
				}
			}(fromClient)
		}
	}()

	// client connecting to proxy
	client, err := net.Dial("tcp", proxyListener.Addr().String())
	if err != nil {
		t.Fatal(err)
	}

	messages := []struct{ Message, Reply string }{
		{"ping", "pong"},
		{"pong", "pong"},
		{"echo", "echo"},
		{"ping", "pong"},
	}

	// client sending to proxy
	for i, m := range messages {
		// writing
		_, err = client.Write([]byte(m.Message))
		if err != nil {
			t.Fatal(err)
		}

		// reading response
		buf := make([]byte, 1024)
		n, err := client.Read(buf)
		if err != nil {
			t.Fatal(err)
		}

		actual := string(buf[:n])
		t.Logf("%q -> proxy -> %q", m.Message, actual)
		if actual != m.Reply {
			t.Errorf("%d: expected reply: %q; actual: %q",
				i, m.Reply, actual)
		}
	}

	_ = client.Close()
	_ = proxyListener.Close()
	_ = serverListener.Close()

	wg.Wait()
}
