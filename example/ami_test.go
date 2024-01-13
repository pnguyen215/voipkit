package example

import (
	"testing"

	"github.com/pnguyen215/voipkit/pkg/ami"
)

func TestAmiClient(t *testing.T) {
	c := ami.GetAmiClientSample()
	ami.D().Info("ami client request: %v", c.String())

	ami.NewClient(ami.NewTcp(), *c)
}
