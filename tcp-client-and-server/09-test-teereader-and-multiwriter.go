package tcp

import (
	"io"
	"log"
	"net"
	"os"
)

type Monitor struct {
	*log.Logger
}

func (m *Monitor) Write(p []byte) (int, error) {
	err := m.Output(2, string(p))
	if err != nil {
		log.Println(err)
	}

	return len(p), nil
}

func TestTeeReaderAndMultiWriter() {
	monitor := &Monitor{Logger: log.New(os.Stdout, "monitor: ", 0)}

	tcpListener, err := net.Listen("tcp", "127.0.0.1:")
	if err != nil {
		monitor.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		defer close(done)

		acceptedConn, err := tcpListener.Accept()
		if err != nil {
			return
		}
		defer func(conn net.Conn) {
			_ = conn.Close()
		}(acceptedConn)

		b := make([]byte, 1024)
		r := io.TeeReader(acceptedConn, monitor)
		n, err := r.Read(b)
		if err != nil && err != io.EOF {
			monitor.Println(err)
			return
		}

		w := io.MultiWriter(acceptedConn, monitor)
		_, err = w.Write(b[:n])
		if err != nil && err != io.EOF {
			monitor.Println(err)
			return
		}
	}()

	clientConn, err := net.Dial("tcp", tcpListener.Addr().String())
	if err != nil {
		monitor.Fatal(err)
	}

	_, err = clientConn.Write([]byte("testing\n"))
	if err != nil {
		monitor.Fatal(err)
	}

	_ = clientConn.Close()
	<-done
}
