package udp

import (
	"context"
	"fmt"
	"net"
)

func startEchoUdpServer(ctx context.Context, addr string) (net.Addr, error) {
	udpListener, err := net.ListenPacket("udp", addr)
	if err != nil {
		return nil, fmt.Errorf("binding to udp %s: %w", addr, err)
	}

	go func() {
		go func() {
			<-ctx.Done()
			_ = udpListener.Close()
		}()

		buf := make([]byte, 1024)
		for {
			n, clientAddr, err := udpListener.ReadFrom(buf) // client to server
			if err != nil {
				return
			}

			_, err = udpListener.WriteTo(buf[:n], clientAddr) // server to client
			if err != nil {
				return
			}
		}
	}()

	return udpListener.LocalAddr(), nil
}
