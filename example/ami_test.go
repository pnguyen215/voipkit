package example

import (
	"testing"
	"time"

	"github.com/pnguyen215/voipkit/pkg/ami"
)

func createConn() (*ami.AMI, error) {
	c := ami.GetAmiClientSample().
		SetEnabled(true).
		SetPort(5038).
		SetUsername("monast").
		SetPassword("T5Monast").
		SetTimeout(5 * time.Second)
	return ami.NewClient(ami.NewTcp(), *c)
}

func TestAmiClient(t *testing.T) {
	_, err := createConn()
	if err != nil {
		ami.D().Error(err.Error())
		return
	}
	ami.D().Info("Authenticated successfully")
}
