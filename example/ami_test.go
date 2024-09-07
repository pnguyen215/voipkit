package example

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/pnguyen215/voipkit/pkg/ami"
)

func createConn() (*ami.AMI, error) {
	c := ami.GetAmiClientSample().
		SetEnabled(true).
		SetPort(5038).
		SetUsername("admin01").
		SetPassword("c71e6acdf318703ec004d9f4d9e17046a673980e").
		SetTimeout(10 * time.Second)
	ami.D().Debug("Asterisk server credentials: %v", c.String())
	return ami.NewClient(ami.NewTcp(), *c)
}

func TestAmiClient(t *testing.T) {
	_, err := createConn()
	if err != nil {
		ami.D().Error(err.Error())
		// t.Fatal("Failed to create AMI connection:", err)
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
	/*
		{
		"action_id": "5ce39195-ae27-2f31-493d-570552478757",
		"channel_type": "SIP",
		"event": "PeerStatus",
		"peer": "SIP/1004",
		"peer_status": "Unknown",
		"privilege": "System"
		}
	*/
	for _, v := range peers {
		ami.D().Info("Status: %v| Privilege: %v| Peer: %v",
			v.Get("peer_status"),
			v.Get("privilege"),
			v.Get("peer"),
		)
	}
}

func TestGetSIPPeer(t *testing.T) {
	c, err := createConn()
	if err != nil {
		ami.D().Error(err.Error())
		return
	}
	c.Core().AddSession()
	c.Core().Dictionary.SetEnabledForceTranslate(true)
	c.Core().Dictionary.AddKeyLinkTranslator("https://raw.githubusercontent.com/pnguyen215/gear-insights-free/master/ami.dictionaries.json")
	peer, err := c.Core().GetSIPPeer(c.Context(), "1004")
	if err != nil {
		ami.D().Error(err.Error())
		return
	}
	/*
		{
		"acl": "Y",
		"action_id": "d6e2496e-8cdf-a178-11f3-5d840a051c38",
		"address_ip": "(null)",
		"address_port": "0",
		"ama_flags": "Unknown",
		"busy_level": "0",
		"call_group": "",
		"call_limit": "2147483647",
		"caller_id": "Ext_1004 <1004>",
		"chan_object_type": "peer",
		"channel_type": "SIP",
		"cid_calling_pres": "Presentation Allowed, Not Screened",
		"codecs": "(alaw)",
		"context": "from-internal",
		"default_addr_ip": "(null)",
		"default_addr_port": "0",
		"default_username": "1004",
		"description": "",
		"dynamic": "Y",
		"lang": "en",
		"last_message_sent": "-1",
		"max_call_br": "384 kbps",
		"max_forwards": "0",
		"md5_secret_exist": "N",
		"moh_suggest": "",
		"named_call_group": "",
		"named_pick_up_group": "",
		"object_name": "1004",
		"parking_lot": "",
		"pickup_group": "",
		"qualify_freq": "60000 ms",
		"reg_contact": "sip:1004@113.161.81.85:58745;ob",
		"reg_expire": "-1 seconds",
		"remote_secret_exist": "N",
		"response": "Success",
		"secret_exist": "Y",
		"sip_auth_in_secure": "no",
		"sip_can_reinvite": "N",
		"sip_comedia": "Y",
		"sip_direct_media": "N",
		"sip_dtmf_mode": "rfc2833",
		"sip_encryption": "N",
		"sip_forcer_port": "Y",
		"sip_promisc_redir": "N",
		"sip_rtcp_mux": "N",
		"sip_rtp_engine": "asterisk",
		"sip_sess_expires": "1800",
		"sip_sess_min": "90",
		"sip_sess_refresh": "uas",
		"sip_sess_timers": "Accept",
		"sip_text_support": "N",
		"sip_time_38_ec": "Unknown",
		"sip_time_38_max_dt_grm": "4294967295",
		"sip_time_38_support": "N",
		"sip_use_reason_header": "N",
		"sip_user_agent": "MicroSIP/3.21.3",
		"sip_user_phone": "N",
		"sip_video_support": "N",
		"status": "UNKNOWN",
		"to_host": "",
		"tone_zone": "<Not set>",
		"transfer_mode": "open",
		"voicemail_box": ""
		}
	*/
	ami.D().Info("Caller.ID: %v| CID: %v| Context: %v| qualify_freq: %v",
		peer.Get("caller_id"),
		peer.Get("cid_calling_pres"),
		peer.Get("context"),
		peer.Get("qualify_freq"),
	)
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
	queues, err := c.Core().GetQueueStatuses(c.Context(), "")
	if err != nil {
		ami.D().Error(err.Error())
		return
	}
	/*
		{
		"action_id": "ce7e68b1-b9c5-209a-46e1-aacadf38d9b1",
		"calls_taken": "0",
		"event": "QueueMember",
		"in_call": "0",
		"last_call": "0",
		"last_pause": "0",
		"location": "Local/8102@from-queue/n",
		"login_time": "1701427389",
		"membership": "dynamic",
		"name": "Ext_8102",
		"paused": "0",
		"paused_reason": "",
		"penalty": "0",
		"queue": "8002",
		"status": "5",
		"status_interface": "hint:8102@ext-local",
		"wrap_uptime": "0"
		}
	*/
	for _, v := range queues {
		ami.D().Info("Location: %v| Login Time: %v| Name: %v| Queue: %v ",
			v.Get("location"),
			v.Get("login_time"),
			v.Get("name"),
			v.Get("queue"),
		)
	}
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
	queues, err := c.Core().GetQueueSummary(c.Context(), "")
	if err != nil {
		ami.D().Error(err.Error())
		return
	}
	/*
		[
			{
				"action_id": "6aa0c877-193a-1be2-d401-d833e73272b1",
				"available": "0",
				"callers": "0",
				"event": "QueueSummary",
				"hold_time": "0",
				"logged_in": "0",
				"longest_hold_time": "0",
				"queue": "default",
				"talk_time": "0"
			}
		]
	*/
	for _, v := range queues {
		ami.D().Info("Queue: %v| Hold Time: %v| Logged In: %v| Longest Hold Time: %v| Talk Time: %v| Available: %v",
			v.Get("queue"),
			v.Get("hold_time"),
			v.Get("logged_in"),
			v.Get("longest_hold_time"),
			v.Get("talk_time"),
			v.Get("available"),
		)
	}
}

func TestExtensionStateList(t *testing.T) {
	c, err := createConn()
	if err != nil {
		ami.D().Error(err.Error())
		return
	}
	c.Core().AddSession()
	c.Core().Dictionary.SetEnabledForceTranslate(true)
	c.Core().Dictionary.AddKeyLinkTranslator("https://raw.githubusercontent.com/pnguyen215/gear-insights-free/master/ami.dictionaries.json")
	extensions, err := c.Core().ExtensionStateList(c.Context())
	if err != nil {
		ami.D().Error(err.Error())
		return
	}
	/*
	  {
	    "action_id": "7443bd11-e2c6-71bf-8d19-cc3cb49b71ec",
	    "context": "ext-local",
	    "event": "ExtensionStatus",
	    "exten": "1010",
	    "hint": "SIP/1010\u0026Custom:DND1010,CustomPresence:1010",
	    "status": "0",
	    "status_text": "Idle"
	  }
	*/
	for _, v := range extensions {
		ami.D().Info("Extension: %v| Status: %v| Status Text: %v| Context: %v",
			v.Get("exten"),
			v.Get("status"),
			v.Get("status_text"),
			v.Get("context"),
		)
	}
}

func TestExtensionState(t *testing.T) {
	c, err := createConn()
	if err != nil {
		ami.D().Error(err.Error())
		return
	}
	c.Core().AddSession()
	c.Core().Dictionary.SetEnabledForceTranslate(true)
	c.Core().Dictionary.AddKeyLinkTranslator("https://raw.githubusercontent.com/pnguyen215/gear-insights-free/master/ami.dictionaries.json")
	extensions, err := c.Core().ExtensionStates(c.Context())
	if err != nil {
		ami.D().Error(err.Error())
		return
	}
	/*
	  {
	    "action_id": "7443bd11-e2c6-71bf-8d19-cc3cb49b71ec",
	    "context": "ext-local",
	    "event": "ExtensionStatus",
	    "exten": "1010",
	    "hint": "SIP/1010\u0026Custom:DND1010,CustomPresence:1010",
	    "status": "0",
	    "status_text": "Idle"
	  }
	*/
	for _, v := range extensions {
		ami.D().Info("Extension: %v| Status: %v| Status Text: %v| Context: %v",
			v.Get("exten"),
			v.Get("status"),
			v.Get("status_text"),
			v.Get("context"),
		)
	}
}

func TestCommand(t *testing.T) {
	c, err := createConn()
	if err != nil {
		ami.D().Error(err.Error())
		return
	}
	c.Socket().SetDebugMode(true)
	c.Core().AddSession()
	c.Core().Dictionary.SetEnabledForceTranslate(true)
	c.Core().Dictionary.AddKeyLinkTranslator("https://raw.githubusercontent.com/pnguyen215/gear-insights-free/master/ami.dictionaries.json")
	response, err := c.Core().Command(c.Context(), "sip show users")
	if err != nil {
		ami.D().Error(err.Error())
		return
	}
	fmt.Println(ami.JsonString(response))
}

func TestAddSIPExtension(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	c, err := createConn()
	if err != nil {
		ami.D().Error(err.Error())
		// t.Fatal("Failed to create AMI connection:", err)
		return
	}
	c.Core().AddSession()
	exten := "9001"
	context := "from-internal"
	application := "Dial"
	appData := "SIP/${EXTEN}"
	payload := ami.NewAMIPayloadExtension().
		SetExtension(exten).
		SetContext(context).
		SetApplication(application).
		SetApplicationData(appData)
	reply, err := c.Core().AddDialplanExtension(ctx, *payload)
	if err != nil {
		t.Fatal("Failed to add SIP extension:", err)
	}
	fmt.Println("Add SIP Extension Response:", reply)
}
