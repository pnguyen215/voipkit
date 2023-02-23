package ami

import (
	"bufio"
	"context"
	"net"
	"net/textproto"
)

func BuildContext(conn net.Conn) (*AMI, context.Context) {
	ctx, cancel := context.WithCancel(context.Background())

	client := &AMI{
		Reader: textproto.NewReader(bufio.NewReader(conn)),
		Writer: bufio.NewWriter(conn),
		Conn:   conn,
		Cancel: cancel,
	}

	return client, ctx
}
