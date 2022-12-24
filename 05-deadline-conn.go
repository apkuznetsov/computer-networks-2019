package tcp

import (
	"io"
	"log"
	"net"
	"time"
)

func DeadlineConn() {
	tcpListener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		log.Fatal(err)
	}

	// server: accept and read
	done := make(chan struct{})
	go func() {
		connAccepted, err := tcpListener.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		defer func() {
			_ = connAccepted.Close()
			close(done)
		}()

		err = connAccepted.SetDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			log.Fatal(err)
			return
		}

		buf := make([]byte, 1)
		_, err = connAccepted.Read(buf) // blocked until remote node sends data
		errNet, ok := err.(net.Error)
		if !ok || !errNet.Timeout() {
			log.Printf("expected timeout error; actual: %v\n", err)
		}

		// server: read all bytes
		done <- struct{}{}

		// server: new deadline, waiting deadline
		err = connAccepted.SetDeadline(time.Now().Add(5 * time.Second))
		if err != nil {
			log.Println(err)
			return
		}
		_, err = connAccepted.Read(buf)
		if err != nil {
			log.Println(err)
		}
	}()

	// client: connect to server
	conn, err := net.Dial("tcp", tcpListener.Addr().String())
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	// client: waiting server read all bytes
	<-done

	_, err = conn.Write([]byte("1"))
	if err != nil {
		log.Fatal(err)
	}

	// client: getting deadline
	buf := make([]byte, 1)
	_, err = conn.Read(buf) // blocked until remote node sends data
	if err != io.EOF {
		log.Printf("expected server termination; actual: %v\n", err)
	} else {
		log.Println(err)
	}
}
