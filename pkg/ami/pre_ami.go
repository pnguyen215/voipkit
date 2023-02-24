package ami

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"strings"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/fatal"
)

func OpenContext(conn net.Conn) (*AMI, context.Context) {
	ctx, cancel := context.WithCancel(context.Background())

	client := &AMI{
		Reader: textproto.NewReader(bufio.NewReader(conn)),
		Writer: bufio.NewWriter(conn),
		Conn:   conn,
		Cancel: cancel,
	}

	return client, ctx
}

func OpenDial(host string, port int) (net.Conn, error) {
	return OpenDialWith(config.AmiNetworkTcpKey, host, port)
}

func OpenDialWith(network, host string, port int) (net.Conn, error) {

	if !config.AmiNetworkKeys[network] {
		return nil, fatal.AMIErrorNew("AMI: Invalid network")
	}

	if host == "" {
		return nil, fatal.AMIErrorNew("AMI: Host must be not empty")
	}

	if port <= 0 {
		return nil, fatal.AMIErrorNew("AMI: Port must be positive number")
	}

	if strings.HasPrefix(host, config.AmiProtocolHttpKey) {
		host = strings.Replace(host, config.AmiProtocolHttpKey, "", -1)
	}

	if strings.HasPrefix(host, config.AmiProtocolHttpsKey) {
		host = strings.Replace(host, config.AmiProtocolHttpsKey, "", -1)
	}

	form := fmt.Sprintf("%s:%d", host, port)
	log.Printf("AMI: dial connection = %v", form)
	return net.Dial(network, form)
}
