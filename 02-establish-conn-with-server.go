package tcp

import (
	"io"
	"log"
	"net"
	"time"
)

func EstablishConnWithServer() {
	// bind tcp listener
	tcpListener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	// server: listen and accept tcp conns
	go func() {
		defer func() { done <- struct{}{} }()

		for {
			// accept
			acceptedConn, err := tcpListener.Accept()
			if err != nil {
				log.Println(err)
				return
			}

			go func(conn net.Conn) {
				defer func() {
					err := conn.Close()
					if err != nil {
						return
					}
					done <- struct{}{}
				}()

				// server: read
				buf := make([]byte, 1024)
				for {
					n, err := conn.Read(buf)
					if err != nil {
						if err != io.EOF {
							log.Println(err)
						}
						return
					}
					log.Printf("received: %q", buf[:n])
				}
			}(acceptedConn)
		}
	}()

	// client: connect
	dialedConn, err := net.Dial("tcp", tcpListener.Addr().String())
	if err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)

	// client: close dial
	err = dialedConn.Close()
	if err != nil {
		return
	}
	<-done

	// server: close tcp listener
	err = tcpListener.Close()
	if err != nil {
		return
	}
	<-done
}
