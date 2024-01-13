package example

import (
	"testing"

	"github.com/pnguyen215/voipkit/pkg/ami"
)

func TestLogger(t *testing.T) {
	ami.D().Info("Test information: %v", 1)
}
