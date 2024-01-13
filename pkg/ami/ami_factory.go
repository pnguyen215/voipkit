package ami

import (
	"fmt"
	"net"
)

type AmiFactory interface {
	Connect(host string, port int) (net.Conn, error)
}

func NewTcp() *tcpAmiFactory {
	return &tcpAmiFactory{}
}

func NewUdp() *udpAmiFactory {
	return &udpAmiFactory{}
}

func (t *tcpAmiFactory) Connect(host string, port int) (net.Conn, error) {
	return OnTcpConn(host, port)
}

func (u *udpAmiFactory) Connect(host string, port int) (net.Conn, error) {
	return OnUdpConn(host, port)
}

func NewClient(factory AmiFactory, request AmiClient) (*AMI, error) {
	if !request.IsEnabled() {
		return nil, fmt.Errorf("Ami unavailable")
	}
	conn, err := factory.Connect(request.host, request.port)
	if err != nil {
		return nil, err
	}
	return serve(conn, request)
}