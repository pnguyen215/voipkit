package ami

import (
	"fmt"
	"strings"
	"time"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

func NewAmiClient() *AmiClient {
	a := &AmiClient{}
	return a
}

func (a *AmiClient) SetEnabled(value bool) *AmiClient {
	a.enabled = value
	return a
}

func (a *AmiClient) IsEnabled() bool {
	return a.enabled
}

func (a *AmiClient) SetHost(value string) *AmiClient {
	a.host = value
	return a
}

func (a *AmiClient) Host() string {
	return a.host
}

func (a *AmiClient) SetPort(value int) *AmiClient {
	a.port = value
	return a
}

func (a *AmiClient) Port() int {
	return a.port
}

func (a *AmiClient) SetUsername(value string) *AmiClient {
	a.username = value
	return a
}

func (a *AmiClient) Username() string {
	return a.username
}

func (a *AmiClient) SetPassword(value string) *AmiClient {
	a.password = value
	return a
}

func (a *AmiClient) SetTimeout(value time.Duration) *AmiClient {
	a.timeout = value
	return a
}

func (a *AmiClient) SetPrivilege(value string) *AmiClient {
	a.privilege = value
	return a
}

func (a *AmiClient) SetPrivileges(values ...string) *AmiClient {
	v := strings.Join(values, ",")
	a.SetPrivilege(v)
	return a
}

func (a *AmiClient) Timeout() time.Duration {
	return a.timeout
}

func (a *AmiClient) String() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("host=%v;", a.host))
	builder.WriteString(fmt.Sprintf("port=%v;", a.port))
	builder.WriteString(fmt.Sprintf("username=%v;", a.username))
	builder.WriteString(fmt.Sprintf("password=%v;", strings.Repeat("*", 8)))
	builder.WriteString(fmt.Sprintf("privilege=%v;", a.privilege))
	builder.WriteString(fmt.Sprintf("timeout=%v;", a.timeout))
	return builder.String()
}

func GetAmiClientSample() *AmiClient {
	a := NewAmiClient().
		SetEnabled(true).
		SetHost("127.0.0.1").
		SetPort(5038).
		SetUsername("admin").
		SetPassword("password").
		SetPrivilege(config.AmiManagerPerm).
		SetTimeout(10 * time.Second)
	return a
}
