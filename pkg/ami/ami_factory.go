package ami

import (
	"fmt"
	"net"
)

type AmiFactory interface {
	Connect(host string, port int) (net.Conn, error)
}

// NewTcp creates a new instance of the tcpAmiFactory, implementing the AmiFactory interface for TCP connections.
func NewTcp() *tcpAmiFactory {
	return &tcpAmiFactory{}
}

// NewUdp creates a new instance of the udpAmiFactory, implementing the AmiFactory interface for UDP connections.
func NewUdp() *udpAmiFactory {
	return &udpAmiFactory{}
}

// Connect establishes a TCP connection to the specified host and port.
func (t *tcpAmiFactory) Connect(host string, port int) (net.Conn, error) {
	return OnTcpConn(host, port)
}

// Connect establishes a UDP connection to the specified host and port.
func (u *udpAmiFactory) Connect(host string, port int) (net.Conn, error) {
	return OnUdpConn(host, port)
}

// NewClient creates a new AMI client using the provided AmiFactory and AmiClient configuration.
//
// Parameters:
//   - factory: An AmiFactory implementation responsible for establishing a network connection.
//   - request: The AmiClient configuration containing connection and authentication details.
//
// Returns:
//   - *AMI: A pointer to the initialized AMI client if the process succeeds.
//   - error: An error indicating any issues during the initialization or authentication process.
//
// Example:
//
//	amiClient, err := NewClient(NewTcp(), AmiClient{
//	  host:     "localhost",
//	  port:     5038,
//	  username: "admin",
//	  password: "secret*password",
//	  timeout:  time.Second * 5,
//	})
//
// Note: The NewClient function uses the provided AmiFactory to establish a network connection based on the specified
// transport protocol (TCP or UDP). It then calls the serve function to initialize the AMI client, perform authentication,
// and release resources. If any step in the process fails, it returns an appropriate error. If the initialization and
// authentication are successful, it returns a pointer to the AMI client ready for further interaction with the Asterisk
// server.
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
