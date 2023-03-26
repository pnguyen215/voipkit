package ami

import (
	"context"
	"fmt"
	"log"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
)

func NewCore() *AMICore {
	c := &AMICore{}
	c.SetEvent(make(chan AMIResultRaw))
	c.SetStop(make(chan struct{}))
	c.SetSocket(NewAMISocket())
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

func (c *AMICore) ResetUUID() *AMICore {
	c.SetUUID(GenUUIDShorten())
	c.Socket.SetUUID(c.UUID)
	return c
}

// NewAmiCore
// Creating new instance asterisk server connection
// Firstly, create new instance AMISocket
// Secondly, create new request body to login
func NewAmiCore(ctx context.Context, socket *AMISocket, auth *AMIAuth) (*AMICore, error) {
	uuid, err := GenUUID()

	if err != nil {
		return nil, err
	}

	socket.SetUUID(uuid)
	err = Login(ctx, *socket, auth)

	if err != nil {
		return nil, err
	}

	core := NewCore()
	core.SetSocket(socket)
	core.SetUUID(uuid)
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

// GetSIPPeer
// Example:
/*
{
  "acl": "Y",
  "action_id": "2f412a3b-9270-044a-65be-2e2294552179",
  "address_ip": "(null)",
  "address_port": "0",
  "ama_flags": "Unknown",
  "busy_level": "0",
  "call_group": "",
  "call_limit": "2147483647",
  "caller_id": "\"Ext_8103\" <8103>",
  "chan_object_type": "peer",
  "channel_type": "SIP",
  "cid_calling_pres": "Presentation Allowed, Not Screened",
  "codecs": "(alaw)",
  "context": "from-internal",
  "default_addr_ip": "(null)",
  "default_addr_port": "0",
  "default_username": "8103",
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
  "object_name": "8103",
  "parking_lot": "",
  "pickup_group": "",
  "qualify_freq": "60000 ms",
  "reg_contact": "sip:8103@14.169.112.44:52685;transport=TCP;rinstance=a4b66908dac666a9",
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
  "sip_user_agent": "Z 5.5.14 v2.10.18.6",
  "sip_user_phone": "N",
  "sip_video_support": "N",
  "status": "UNKNOWN",
  "to_host": "",
  "tone_zone": "<Not set>",
  "transfer_mode": "open",
  "voicemail_box": "8103@default"
}
*/
func (c *AMICore) GetSIPPeer(ctx context.Context, peer string) (AMIResultRaw, error) {
	return SIPShowPeer(ctx, *c.Socket, peer)
}

// GetSIPPeersStatus
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
func (c *AMICore) GetSIPPeersStatus(ctx context.Context) ([]AMIResultRaw, error) {
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

// GetSIPPeerStatus
// Example:
/*
{
  "action_id": "2e588d76-d771-fb70-d3bd-fced26dc0bf5",
  "channel_type": "SIP",
  "event": "PeerStatus",
  "peer": "SIP/8104",
  "peer_status": "Reachable",
  "privilege": "System",
  "time": "42"
}
*/
func (c *AMICore) GetSIPPeerStatus(ctx context.Context, peer string) (AMIResultRaw, error) {
	peers, err := SIPPeerStatus(ctx, *c.Socket, peer)
	if err != nil {
		return AMIResultRaw{}, err
	}
	if len(peers) == 0 {
		return AMIResultRaw{}, nil
	}
	return peers[0], nil
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

// Challenge generates a challenge for MD5 authentication.
func (c *AMICore) Challenge(ctx context.Context) (AMIResultRaw, error) {
	return Challenge(ctx, *c.Socket)
}

// CreateConfig creates an empty file in the configuration directory.
// This action will create an empty file in the configuration directory.
// This action is intended to be used before an UpdateConfig action.
func (c *AMICore) CreateConfig(ctx context.Context, filename string) (AMIResultRaw, error) {
	return CreateConfig(ctx, *c.Socket, filename)
}

// DataGet retrieves the data api tree.
func (c *AMICore) DataGet(ctx context.Context, path, search, filter string) (AMIResultRaw, error) {
	return DataGet(ctx, *c.Socket, path, search, filter)
}

// EventFlow control Event Flow.
// eventMask: Enable/Disable sending of events to this manager client.
func (c *AMICore) EventFlow(ctx context.Context, eventMask string) (AMIResultRaw, error) {
	return EventFlow(ctx, *c.Socket, eventMask)
}

// GetConfig retrieves configuration.
// This action will dump the contents of a configuration file by category and contents or optionally by specified category only.
func (c *AMICore) GetConfig(ctx context.Context, filename, category, filter string) (AMIResultRaw, error) {
	return GetConfig(ctx, *c.Socket, filename, category, filter)
}

// GetConfigJson retrieves configuration (JSON format).
// This action will dump the contents of a configuration file by category and contents in JSON format.
// This only makes sense to be used using raw man over the HTTP interface.
func (c *AMICore) GetConfigJson(ctx context.Context, filename, category, filter string) (AMIResultRaw, error) {
	return GetConfigJson(ctx, *c.Socket, filename, category, filter)
}

// JabberSend sends a message to a Jabber Client
func (c *AMICore) JabberSend(ctx context.Context, jabber, jid, message string) (AMIResultRaw, error) {
	return JabberSend(ctx, *c.Socket, jabber, jid, message)
}

// ListCategories lists categories in configuration file.
// Example:
// filename like: manager.conf, extensions.conf, sip.conf...
func (c *AMICore) ListCategories(ctx context.Context, filename string) (AMIResultRaw, error) {
	return ListCategories(ctx, *c.Socket, filename)
}

// ModuleCheck checks if module is loaded.
// Checks if Asterisk module is loaded. Will return Success/Failure. For success returns, the module revision number is included.
func (c *AMICore) ModuleCheck(ctx context.Context, module string) (AMIResultRaw, error) {
	return ModuleCheck(ctx, *c.Socket, module)
}

// ModuleLoad module management.
// Loads, unloads or reloads an Asterisk module in a running system.
func (c *AMICore) ModuleLoad(ctx context.Context, module, loadType string) (AMIResultRaw, error) {
	return ModuleLoad(ctx, *c.Socket, module, loadType)
}

// Reload Sends a reload event.
func (c *AMICore) Reload(ctx context.Context, module string) (AMIResultRaw, error) {
	return Reload(ctx, *c.Socket, module)
}

// ShowDialPlan shows dialplan contexts and extensions
// Be aware that showing the full dialplan may take a lot of capacity.
func (c *AMICore) ShowDialPlan(ctx context.Context, extension, context string) ([]AMIResultRaw, error) {
	return ShowDialPlan(ctx, *c.Socket, extension, context)
}

// Filter dynamically add filters for the current manager session.
func (c *AMICore) Filter(ctx context.Context, operation, filter string) (AMIResultRaw, error) {
	return Filter(ctx, *c.Socket, operation, filter)
}

// DeviceStateList list the current known device states.
func (c *AMICore) GetDeviceStateList(ctx context.Context) ([]AMIResultRaw, error) {
	return DeviceStateList(ctx, *c.Socket)
}

// LoggerRotate reload and rotate the Asterisk logger.
func (c *AMICore) LoggerRotate(ctx context.Context) (AMIResultRaw, error) {
	return LoggerRotate(ctx, *c.Socket)
}

// UpdateConfig Updates a config file.
// Dynamically updates an Asterisk configuration file.
// Example:
/*
Action: UpdateConfig
SrcFilename: voicemail2.conf
DstFilename: voicemail2.conf
Action-000000: Append
Cat-000000: default
Var-000000: 127
Value-000000: >5555, Jason Bourne97, ***@noCia.gov.do
Action-000001: Append
Cat-000001: default
Var-000001: 125
Value-000001: >55555, Jason Bourne76, ***@noCia.gov.do
Action-000002: Append
Cat-000002: default
Var-000002: 122
Value-000002: >5555, Jason Bourne74, ***@noCia.gov.do
Action-000003: Append
Cat-000003: default
Var-000003: 128
Value-000003: >5555, Jason Bourne48, ***@noCia.gov.do
Action-000004: Append
Cat-000004: default
Var-000004: 126
Value-000004: >55555, Jason Bourne18, ***@noCia.gov.do
ActionID: 495446608
*/
func (c *AMICore) UpdateConfig(ctx context.Context, sourceFilename, destinationFilename string, reload bool, actions ...AMIUpdateConfigAction) (AMIResultRaw, error) {
	return UpdateConfig(ctx, *c.Socket, sourceFilename, destinationFilename, reload, actions...)
}
