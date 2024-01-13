package ami

import (
	"bufio"
	"context"
	"net"
	"net/textproto"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

// OnConnContext initializes a new AMI client with the provided network connection.
// It returns the initialized AMI client along with a cancellable context for managing its lifecycle.
//
// Parameters:
//   - conn: The network connection used by the AMI client.
//
// Returns:
//   - An initialized AMI client (*AMI).
//   - A cancellable context for managing the AMI client's lifecycle (context.Context).
//
// Example:
//
//	// Creating an AMI client with a network connection
//	conn, err := net.Dial("tcp", "localhost:5038")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	client, ctx := OnConnContext(conn)
//	// Use the client and context for AMI operations.
//	// Make sure to close the connection and cancel the context when done.
//	defer client.Close()
//	defer client.Cancel()
func OnConnContext(conn net.Conn) (*AMI, context.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	client := &AMI{
		Reader: textproto.NewReader(bufio.NewReader(conn)),
		Writer: bufio.NewWriter(conn),
		Conn:   conn,
		Cancel: cancel,
	}
	// Check if the connection is available
	if conn != nil {
		addr := conn.RemoteAddr().String()
		_socket, err := NewAmiSocketContext(ctx, addr)

		if err == nil {
			client.Socket = _socket
			D().Info("OnConnContext, cloning (addr: %v) socket connection succeeded", addr)
		}
	}
	return client, ctx
}

// OnTcpConn opens a network connection to the specified IP address and port using the default TCP network.
//
// Parameters:
//   - ip:   The IP address to connect to.
//   - port: The port number to connect to.
//
// Returns:
//   - The opened network connection (net.Conn).
//   - An error if the connection cannot be established.
//
// Example:
//
//	// Dialing an AMI server at localhost on port 5038
//	conn, err := OnTcpConn("localhost", 5038)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// Use the connection for AMI operations.
//	// Make sure to close the connection when done.
//	defer conn.Close()
func OnTcpConn(ip string, port int) (net.Conn, error) {
	return NewConn(config.AmiNetworkTcpKey, ip, port)
}

// OnUdpConn opens a network connection to the specified IP address and port using the default UDP network.
//
// Parameters:
//   - ip:   The IP address to connect to.
//   - port: The port number to connect to.
//
// Returns:
//   - The opened network connection (net.Conn).
//   - An error if the connection cannot be established.
//
// Example:
//
//	// Dialing an AMI server at localhost on port 5038
//	conn, err := OnUdpConn("localhost", 5038)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// Use the connection for AMI operations.
//	// Make sure to close the connection when done.
//	defer conn.Close()
func OnUdpConn(ip string, port int) (net.Conn, error) {
	return NewConn(config.AmiNetworkUdpKey, ip, port)
}

// NewConn opens a network connection to the specified IP address and port using the specified network type.
//
// Parameters:
//   - network: The network type ("tcp", "udp", etc.).
//   - ip:      The IP address to connect to.
//   - port:    The port number to connect to.
//
// Returns:
//   - The opened network connection (net.Conn).
//   - An error if the connection cannot be established.
//
// Example:
//
//	// Dialing an AMI server at localhost on port 5038 using UDP
//	conn, err := NewConn("udp", "localhost", 5038)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// Use the connection for AMI operations.
//	// Make sure to close the connection when done.
//	defer conn.Close()
func NewConn(network, ip string, port int) (net.Conn, error) {
	if !config.AmiNetworkKeys[network] {
		return nil, AMIErrorNew("AMI: Invalid network")
	}
	if IsStringEmpty(ip) {
		return nil, AMIErrorNew("AMI: IP must be not empty")
	}
	if port <= 0 {
		return nil, AMIErrorNew("AMI: Port must be positive number")
	}
	host, _port, _ := DecodeIp(ip)
	if len(host) > 0 && len(_port) > 0 {
		form := net.JoinHostPort(host, _port)
		D().Info("AMI: (IP decoded) dial connection: %v", form)
		return net.Dial(network, form)
	}
	form := RemoveProtocol(ip, port)
	D().Info("AMI: dial connection: %v", form)
	return net.Dial(network, form)
}
