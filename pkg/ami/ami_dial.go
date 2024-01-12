package ami

import (
	"bufio"
	"context"
	"log"
	"net"
	"net/textproto"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

// OpenContext
func OpenContext(conn net.Conn) (*AMI, context.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	client := &AMI{
		Reader: textproto.NewReader(bufio.NewReader(conn)),
		Writer: bufio.NewWriter(conn),
		Conn:   conn,
		Cancel: cancel,
	}
	// checking conn available
	if conn != nil {
		addr := conn.RemoteAddr().String()
		_socket, err := NewAMISocketWith(ctx, addr)

		if err == nil {
			client.Socket = _socket
			log.Printf("OpenContext, cloning (addr: %v) socket connection succeeded", addr)
		}
	}
	return client, ctx
}

// OpenDial
func OpenDial(ip string, port int) (net.Conn, error) {
	return OpenDialWith(config.AmiNetworkTcpKey, ip, port)
}

// OpenDialWith
func OpenDialWith(network, ip string, port int) (net.Conn, error) {
	if !config.AmiNetworkKeys[network] {
		return nil, AMIErrorNew("AMI: Invalid network")
	}
	if ip == "" {
		return nil, AMIErrorNew("AMI: IP must be not empty")
	}
	if port <= 0 {
		return nil, AMIErrorNew("AMI: Port must be positive number")
	}
	host, _port, _ := DecodeIp(ip)
	if len(host) > 0 && len(_port) > 0 {
		form := net.JoinHostPort(host, _port)
		log.Printf("AMI: (IP decoded) dial connection = %v", form)
		return net.Dial(network, form)
	}
	form := RemoveProtocol(ip, port)
	log.Printf("AMI: dial connection = %v", form)
	return net.Dial(network, form)
}
