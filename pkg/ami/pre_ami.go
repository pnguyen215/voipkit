package ami

import (
	"bufio"
	"context"
	"log"
	"net"
	"net/textproto"
	"strconv"
	"strings"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/fatal"
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/utils"
)

func OpenContext(conn net.Conn) (*AMI, context.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	socket, err := NewAMISocketConn(ctx, conn)

	if err != nil {
		log.Printf("OpenContext, socket has an error occurred: %v", err.Error())
	}

	client := &AMI{
		Reader: textproto.NewReader(bufio.NewReader(conn)),
		Writer: bufio.NewWriter(conn),
		Conn:   conn,
		Cancel: cancel,
		Socket: socket,
	}

	return client, ctx
}

func OpenDial(ip string, port int) (net.Conn, error) {
	return OpenDialWith(config.AmiNetworkTcpKey, ip, port)
}

func OpenDialWith(network, ip string, port int) (net.Conn, error) {

	if !config.AmiNetworkKeys[network] {
		return nil, fatal.AMIErrorNew("AMI: Invalid network")
	}

	if ip == "" {
		return nil, fatal.AMIErrorNew("AMI: IP must be not empty")
	}

	if port <= 0 {
		return nil, fatal.AMIErrorNew("AMI: Port must be positive number")
	}

	host, _port, _ := utils.IPDecode(ip)

	if len(host) > 0 && len(_port) > 0 {
		form := net.JoinHostPort(host, _port)
		log.Printf("AMI: (IP decoded) dial connection = %v", form)
		return net.Dial(network, form)
	}

	form := RemoveProtocol(ip, port)
	log.Printf("AMI: dial connection = %v", form)
	return net.Dial(network, form)
}

func RemoveProtocol(ip string, port int) string {
	if strings.HasPrefix(ip, config.AmiProtocolHttpKey) {
		ip = strings.Replace(ip, config.AmiProtocolHttpKey, "", -1)
	}

	if strings.HasPrefix(ip, config.AmiProtocolHttpsKey) {
		ip = strings.Replace(ip, config.AmiProtocolHttpsKey, "", -1)
	}

	form := net.JoinHostPort(ip, strconv.Itoa(port))
	return form
}
