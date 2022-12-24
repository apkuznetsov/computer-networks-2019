package tcp

import (
	"log"
	"net"
)

func BindListenAndAcceptConns() {
	// bind tcp listener
	tcpListener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = tcpListener.Close() }()
	log.Printf("bound to %q\n", tcpListener.Addr())

	// listen and accept tcp conns
	for {
		acceptedTcpConn, err := tcpListener.Accept()
		if err != nil {
			log.Println(err)
		} else {
			go func(conn net.Conn) {
				defer func(c net.Conn) {
					err := c.Close()
					if err != nil {
						log.Println(err)
					}
				}(conn)
			}(acceptedTcpConn)
		}
	}
}
