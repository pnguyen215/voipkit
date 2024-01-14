package example

import (
	"fmt"
	"log"
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
	ami.D().Debug("Asterisk server credentials: %v", c.String())
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

func TestAllEventConsume(t *testing.T) {
	c, err := createConn()
	if err != nil {
		ami.D().Error(err.Error())
		return
	}
	event := ami.NewEventListener()
	event.OpenFullEventsAsyncFunc(c)
}

func TestDialOut(t *testing.T) {
	// adding logic here
}

func TestChanspy(t *testing.T) {

}

func TestGetSIPPeersStatus(t *testing.T) {
	c, err := createConn()
	if err != nil {
		ami.D().Error(err.Error())
		return
	}
	c.Core().AddSession()
	peers, err := c.Core().GetSIPPeersStatus(c.Context())
	if err != nil {
		ami.D().Error(err.Error())
		return
	}
	log.Println(fmt.Sprintf("SIP peer status: %v", ami.JsonString(peers[1])))
}

func TestGetQueueStatuses(t *testing.T) {
	c, err := createConn()
	if err != nil {
		ami.D().Error(err.Error())
		return
	}
	c.Core().AddSession()
	c.Core().Dictionary.SetEnabledForceTranslate(true)
	c.Core().Dictionary.AddKeyLinkTranslator("https://raw.githubusercontent.com/pnguyen215/gear-insights-free/master/ami.dictionaries.json")
	peers, err := c.Core().GetQueueStatuses(c.Context(), "")
	if err != nil {
		ami.D().Error(err.Error())
		return
	}
	log.Println(fmt.Sprintf("SIP queues status: %v", ami.JsonString(peers)))
}

func TestGetQueueSummary(t *testing.T) {
	c, err := createConn()
	if err != nil {
		ami.D().Error(err.Error())
		return
	}
	c.Core().AddSession()
	c.Core().Dictionary.SetEnabledForceTranslate(true)
	c.Core().Dictionary.AddKeyLinkTranslator("https://raw.githubusercontent.com/pnguyen215/gear-insights-free/master/ami.dictionaries.json")
	peers, err := c.Core().GetQueueSummary(c.Context(), "")
	if err != nil {
		ami.D().Error(err.Error())
		return
	}
	log.Println(fmt.Sprintf("SIP queues summary: %v", ami.JsonString(peers)))
}
