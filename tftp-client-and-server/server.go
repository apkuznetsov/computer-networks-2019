package tftp

import (
	"bytes"
	"errors"
	"log"
	"net"
	"time"
)

type Server struct {
	Payload    []byte        // the payload served for all read requests
	RetriesNum uint8         // the number of times to retry a failed transmission
	AckTimeout time.Duration // the duration to wait for an acknowledgment
}

func (s Server) ListenAndServe(addr string) error {
	conn, err := net.ListenPacket("udp", addr)
	if err != nil {
		return err
	}
	defer func() { _ = conn.Close() }()

	log.Printf("Listening on %s ...\n", conn.LocalAddr())
	return s.Serve(conn)
}

func (s *Server) Serve(conn net.PacketConn) error {
	if conn == nil {
		return errors.New("nil connection")
	}
	if s.Payload == nil {
		return errors.New("payload is required")
	}

	if s.RetriesNum == 0 {
		s.RetriesNum = 10
	}
	if s.AckTimeout == 0 {
		s.AckTimeout = 6 * time.Second
	}

	var readreq ReadReq
	for {
		buf := make([]byte, DatagramSize)

		_, addr, err := conn.ReadFrom(buf)
		if err != nil {
			return err
		}

		err = readreq.UnmarshalBinary(buf)
		if err != nil {
			log.Printf("[%s] bad request: %v", addr, err)
			continue
		}

		go s.handle(addr.String(), readreq)
	}
}

func (s Server) handle(clientAddr string, readReq ReadReq) {
	log.Printf("[%s] requested file: %s", clientAddr, readReq.Filename)

	connWithClient, err := net.Dial("udp", clientAddr)
	if err != nil {
		log.Printf("[%s] dial: %v", clientAddr, err)
		return
	}
	defer func() { _ = connWithClient.Close() }()

	var (
		ackPkt  Ack
		errPkt  Err
		dataPkt = Data{Payload: bytes.NewReader(s.Payload)}
		buf     = make([]byte, DatagramSize)
	)

NEXTPACKET:
	for n := DatagramSize; n == DatagramSize; {
		dataMarshed, err := dataPkt.MarshalBinary()
		if err != nil {
			log.Printf("[%s] preparing dataMarshed packet: %v", clientAddr, err)
			return
		}

	RETRY:
		for i := s.RetriesNum; i > 0; i-- {
			n, err = connWithClient.Write(dataMarshed) // send the data packet
			if err != nil {
				log.Printf("[%s] write: %v", clientAddr, err)
				return
			}

			// wait for the client's ACK packet
			_ = connWithClient.SetReadDeadline(time.Now().Add(s.AckTimeout))

			_, err = connWithClient.Read(buf)
			if err != nil {
				if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
					continue RETRY
				}
				log.Printf("[%s] waiting for ACK: %v", clientAddr, err)
				return
			}

			switch {
			case ackPkt.UnmarshalBinary(buf) == nil:
				if uint16(ackPkt) == dataPkt.Block {
					continue NEXTPACKET // received ACK; send next data packet
				}
			case errPkt.UnmarshalBinary(buf) == nil:
				log.Printf("[%s] received error: %v",
					clientAddr, errPkt.Message)
				return
			default:
				log.Printf("[%s] bad packet", clientAddr)
			}
		}

		log.Printf("[%s] exhausted retries", clientAddr)
		return
	}

	log.Printf("[%s] sent %d blocks", clientAddr, dataPkt.Block)
}
