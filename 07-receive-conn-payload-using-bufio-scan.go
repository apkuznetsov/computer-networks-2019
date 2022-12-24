package tcp

import (
	"bufio"
	"log"
	"net"
	"reflect"
)

const payload = "The bigger the interface, the weaker the abstraction."

func ReceiveConnPayloadUsingBufioScan() {
	tcpListener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		log.Fatal(err)
	}

	// server accepting then writing payload to client
	go func() {
		connAccepted, err := tcpListener.Accept()
		if err != nil {
			log.Println(err)
			return
		}
		defer func(conn net.Conn) {
			_ = conn.Close()
		}(connAccepted)

		_, err = connAccepted.Write([]byte(payload))
		if err != nil {
			log.Println(err)
		}
	}()

	conn, err := net.Dial("tcp", tcpListener.Addr().String())
	if err != nil {
		log.Fatal(err)
	}
	defer func(conn net.Conn) {
		_ = conn.Close()
	}(conn)

	scanner := bufio.NewScanner(conn)
	scanner.Split(bufio.ScanWords)
	var words []string
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	err = scanner.Err()
	if err != nil {
		log.Println(err)
	}

	expected := []string{"The", "bigger", "the", "interface,", "the",
		"weaker", "the", "abstraction."}
	if !reflect.DeepEqual(words, expected) {
		log.Fatal("inaccurate scanned word list")
	}

	log.Printf("Scanned words: %#v\n", words)
}
