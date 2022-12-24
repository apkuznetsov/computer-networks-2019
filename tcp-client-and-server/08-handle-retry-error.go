package tcp

import (
	"errors"
	"log"
	"net"
	"time"
)

func HandleRetryError(conn net.Conn) error {
	var (
		err error
		n   int
		i   = 7 // maximum number of retries
	)

	for ; i > 0; i-- {
		n, err = conn.Write([]byte("hello world"))
		if err != nil {
			if nErr, ok := err.(net.Error); ok && nErr.Temporary() {
				log.Println("temporary error:", nErr)
				time.Sleep(10 * time.Second)
				continue
			}
			return err
		}
		break
	}

	if i == 0 {
		return errors.New("temporary write failure threshold exceeded")
	}
	log.Printf("wrote %d bytes to %s\n", n, conn.RemoteAddr())

	return nil
}
