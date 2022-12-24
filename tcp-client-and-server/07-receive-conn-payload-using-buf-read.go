package tcp

import (
	"crypto/rand"
	"io"
	"log"
	"net"
)

func ReceiveConnBigPayloadUsingSmallBuf() {
	payload := make([]byte, 1<<24) // 16 MB
	_, err := rand.Read(payload)   // generate a random payload
	if err != nil {
		log.Fatal(err)
	}

	tcpListener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		log.Fatal(err)
	}

	// server accepting then writing payload to client
	go func() {
		conn, err := tcpListener.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		defer func(conn net.Conn) {
			_ = conn.Close()
		}(conn)

		_, err = conn.Write(payload)
		if err != nil {
			log.Println(err)
		}
	}()

	// client dialing
	conn, err := net.Dial("tcp", tcpListener.Addr().String())
	if err != nil {
		log.Fatal(err)
	}

	// client reading while eof
	buf := make([]byte, 31<<19) // 512 KB
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Println(err)
			}
			break
		}

		log.Printf("read %d bytes'n", n) // buf[:n] is the data read from conn
	}

	_ = conn.Close()
}
