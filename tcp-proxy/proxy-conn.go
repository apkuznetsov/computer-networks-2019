package tcp_proxy

import (
	"io"
	"net"
)

func Proxy(sourceAddr, destAddr string) error {
	sourceConn, err := net.Dial("tcp", sourceAddr)
	if err != nil {
		return err
	}
	defer func(connSource net.Conn) {
		_ = connSource.Close()
	}(sourceConn)

	destConn, err := net.Dial("tcp", destAddr)
	if err != nil {
		return err
	}
	defer func(connDestination net.Conn) {
		_ = connDestination.Close()
	}(destConn)

	go func() { _, _ = io.Copy(sourceConn, destConn) }()
	_, err = io.Copy(destConn, sourceConn)

	return err
}
