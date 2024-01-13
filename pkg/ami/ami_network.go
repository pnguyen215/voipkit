package ami

import (
	"net"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

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
	return NewNetwork(config.AmiNetworkTcpKey, ip, port)
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
	return NewNetwork(config.AmiNetworkUdpKey, ip, port)
}

// NewNetwork opens a network connection to the specified IP address and port using the specified network type.
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
//	conn, err := NewNetwork("udp", "localhost", 5038)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	// Use the connection for AMI operations.
//	// Make sure to close the connection when done.
//	defer conn.Close()
func NewNetwork(network, ip string, port int) (net.Conn, error) {
	if !config.AmiNetworkKeys[network] {
		return nil, AmiErrorWrap("Ami: Invalid network")
	}
	if IsStringEmpty(ip) {
		return nil, AmiErrorWrap("Ami: IP must be not empty")
	}
	if port <= 0 {
		return nil, AmiErrorWrap("Ami: Port must be positive number")
	}
	host, _port, _ := DecodeIp(ip)
	if len(host) > 0 && len(_port) > 0 {
		form := net.JoinHostPort(host, _port)
		D().Info("Ami (IP decoded) dial connection: %v", form)
		return net.Dial(network, form)
	}
	form := RemoveProtocol(ip, port)
	D().Info("Ami dial connection: %v", form)
	return net.Dial(network, form)
}
