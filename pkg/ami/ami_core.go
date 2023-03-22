package ami

import (
	"context"
	"fmt"
	"log"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
)

func NewCore() *AMICore {
	c := &AMICore{}
	return c
}

func (c *AMICore) SetSocket(socket *AMISocket) *AMICore {
	c.Socket = socket
	return c
}

func (c *AMICore) SetUUID(id string) *AMICore {
	c.UUID = id
	return c
}

func (c *AMICore) SetEvent(event chan AMIResultRaw) *AMICore {
	c.Event = event
	return c
}

func (c *AMICore) SetStop(stop chan struct{}) *AMICore {
	c.Stop = stop
	return c
}

func (c *AMICore) SetDictionary(dictionary *AMIDictionary) *AMICore {
	c.Dictionary = dictionary
	return c
}

// NewAmiCore
// Creating new instance asterisk server connection
// Firstly, create new instance AMISocket
// Secondly, create new request body to login
func NewAmiCore(ctx context.Context, socket AMISocket, auth *AMIAuth) (*AMICore, error) {
	uuid, err := GenUUID()

	if err != nil {
		return nil, err
	}

	socket.SetUUID(uuid)
	err = Login(ctx, socket, auth)

	if err != nil {
		return nil, err
	}

	core := NewCore()
	core.SetSocket(&socket)
	core.SetUUID(uuid)
	core.SetEvent(make(chan AMIResultRaw))
	core.SetStop(make(chan struct{}))
	core.SetDictionary(socket.Dictionary)

	core.Wg.Add(1)
	go core.Run(ctx)
	return core, nil
}

// Run
// Go-func to consume event from asterisk server response
func (c *AMICore) Run(ctx context.Context) {
	defer c.Wg.Done()
	for {
		select {
		case <-c.Stop:
			return
		case <-ctx.Done():
			return
		default:
			event, err := Events(ctx, *c.Socket)
			if err != nil {
				log.Printf(config.AmiErrorConsumeEvent, err)
				return
			}
			c.Event <- event
		}
	}
}

// Events
// Consume all events will be returned an channel received from asterisk server log.
func (c *AMICore) Events() <-chan AMIResultRaw {
	return c.Event
}

// GetSIPPeers
// GetSIPPeers fetch the list of SIP peers present on asterisk.
// Example:
/*
{
    "Address-IP": "14.238.106.54",
    "Address-Port": "5060",
    "Busy-level": "0",
    "CID-CallingPres": "Presentation Allowed, Not Screened",
    "Call-limit": "2147483647",
    "Codecs": "(alaw)",
    "Default-Username": "5002",
    "Default-addr-IP": "(null)",
    "Default-addr-port": "0",
    "LastMsgsSent": "-1",
    "MD5SecretExist": "N",
    "MaxCallBR": "384 kbps",
    "Maxforwards": "0",
    "Named Callgroup": "",
    "Named Pickupgroup": "",
    "Parkinglot": "",
    "QualifyFreq": "60000 ms",
    "Reg-Contact": "sip:5002@14.238.106.54:5060",
    "RegExpire": "3579 seconds",
    "RemoteSecretExist": "N",
    "SIP-AuthInsecure": "no",
    "SIP-CanReinvite": "N",
    "SIP-Comedia": "Y",
    "SIP-DTMFmode": "rfc2833",
    "SIP-DirectMedia": "N",
    "SIP-Encryption": "N",
    "SIP-Forcerport": "Y",
    "SIP-PromiscRedir": "N",
    "SIP-RTCP-Mux": "N",
    "SIP-RTP-Engine": "asterisk",
    "SIP-Sess-Expires": "1800",
    "SIP-Sess-Min": "90",
    "SIP-Sess-Refresh": "uas",
    "SIP-Sess-Timers": "Accept",
    "SIP-T.38EC": "Unknown",
    "SIP-T.38MaxDtgrm": "4294967295",
    "SIP-T.38Support": "N",
    "SIP-TextSupport": "N",
    "SIP-Use-Reason-Header": "N",
    "SIP-UserPhone": "N",
    "SIP-Useragent": "ATCOM A1x-2.6.5.d0809 80828708CF75",
    "SIP-VideoSupport": "N",
    "SecretExist": "Y",
    "ToHost": "",
    "TransferMode": "open",
    "VoiceMailbox": "",
    "acl": "Y",
    "action_id": "bf70fb4d-1c7b-5b1d-e63e-e8a1ab189469",
    "ama_flags": "Unknown",
    "call_group": "",
    "caller_id": "\"TA_5002\" <5002>",
    "chan_object_type": "peer",
    "channel_type": "SIP",
    "context": "from-internal",
    "description": "",
    "dynamic": "Y",
    "lang": "en",
    "moh_suggest": "",
    "object_name": "5002",
    "pickup_group": "",
    "response": "Success",
    "status": "OK (41 ms)",
    "tone_zone": "<Not set>"
  }
*/
func (c *AMICore) GetSIPPeers(ctx context.Context) ([]AMIResultRaw, error) {
	var peers []AMIResultRaw
	response, err := SIPPeers(ctx, *c.Socket)
	switch {
	case err != nil:
		return nil, err
	case len(response) == 0:
		return nil, fmt.Errorf(config.AmiErrorNoExtensionConfigured)
	default:
		for _, v := range response {
			peer, err := SIPShowPeer(ctx, *c.Socket, v.GetVal(config.AmiJsonFieldObjectName))
			if err != nil {
				return nil, err
			}
			peers = append(peers, peer)
		}
	}

	return peers, nil
}

// GetSIPPeerStatus
// Example:
/*
{
    "action_id": "b6514f81-6047-6b2c-0097-6701bfe6ecd4",
    "channel_type": "SIP",
    "event": "PeerStatus",
    "peer": "SIP/5001",
    "peer_status": "Reachable",
    "privilege": "System",
    "time": "28"
  }
*/
func (c *AMICore) GetSIPPeerStatus(ctx context.Context) ([]AMIResultRaw, error) {
	var peers []AMIResultRaw
	response, err := SIPPeers(ctx, *c.Socket)
	switch {
	case err != nil:
		return nil, err
	case len(response) == 0:
		return nil, fmt.Errorf(config.AmiErrorNoExtensionConfigured)
	default:
		for _, v := range response {
			peer, err := SIPPeerStatus(ctx, *c.Socket, v.GetVal(config.AmiJsonFieldObjectName))
			if err != nil {
				return nil, err
			}
			peers = append(peers, peer...)
		}
	}

	return peers, nil
}

// GetSIPShowRegistry
// Example:
/*

 */
func (c *AMICore) GetSIPShowRegistry(ctx context.Context) ([]AMIResultRaw, error) {
	return SIPShowRegistry(ctx, *c.Socket)
}

// GetSIPQualifyPeer
// Example
/*
{
    "action_id": "2d4ddf88-0eed-dc3d-3ab8-0bba9d603328",
    "message": "SIP peer found - will qualify",
    "response": "Success"
  },
  {
    "action_id": "2d4ddf88-0eed-dc3d-3ab8-0bba9d603328",
    "event": "SIPQualifyPeerDone",
    "peer": "5001",
    "privilege": "call,all"
  },
*/
func (c *AMICore) GetSIPQualifyPeer(ctx context.Context) ([]AMIResultRaw, error) {
	var peers []AMIResultRaw
	response, err := SIPPeers(ctx, *c.Socket)
	switch {
	case err != nil:
		return nil, err
	case len(response) == 0:
		return nil, fmt.Errorf(config.AmiErrorNoExtensionConfigured)
	default:
		for _, v := range response {
			peer, err := SIPQualifyPeer(ctx, *c.Socket, v.GetVal(config.AmiJsonFieldObjectName))
			if err != nil {
				return nil, err
			}
			peers = append(peers, peer)
		}
	}

	return peers, nil
}

// Logoff
// Logoff closes the current session with AMI.
func (c *AMICore) Logoff(ctx context.Context) error {
	close(c.Stop)
	c.Wg.Wait()
	return Logoff(ctx, *c.Socket)
}

// Ping
func (c *AMICore) Ping(ctx context.Context) error {
	return Ping(ctx, *c.Socket)
}

// Command executes an Asterisk CLI Command.
func (c *AMICore) Command(ctx context.Context, cmd string) (AMIResultRawLevel, error) {
	return Command(ctx, *c.Socket, cmd)
}

// CoreSettings shows PBX core settings (version etc).
func (c *AMICore) GetCoreSettings(ctx context.Context) (AMIResultRaw, error) {
	return CoreSettings(ctx, *c.Socket)
}

// CoreStatus shows PBX core status variables.
func (c *AMICore) GetCoreStatus(ctx context.Context) (AMIResultRaw, error) {
	return CoreStatus(ctx, *c.Socket)
}

// ListCommands lists available manager commands.
// Returns the action name and synopsis for every action that is available to the user
func (c *AMICore) GetListCommands(ctx context.Context) (AMIResultRaw, error) {
	return ListCommands(ctx, *c.Socket)
}
