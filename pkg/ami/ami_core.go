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

// GetQueueStatuses
// QueueStatuses show status all members in queue.
// Example:
/*
[
  {
    "action_id": "ba1f72f2-4d33-48a9-a395-f2295ec19c2b",
    "calls_taken": "0",
    "event": "QueueMember",
    "in_call": "0",
    "last_call": "0",
    "last_pause": "0",
    "location": "Local/8101@from-queue/n",
    "membership": "dynamic",
    "name": "Ext_8101",
    "paused": "0",
    "paused_reason": "",
    "penalty": "0",
    "queue": "8002",
    "status": "5",
    "status_interface": "hint:8101@ext-local",
    "wrap_uptime": "0"
  }
]
*/
func (c *AMICore) GetQueueStatuses(ctx context.Context, queue string) ([]AMIResultRaw, error) {
	return QueueStatuses(ctx, *c.Socket, queue)
}

// GetQueueSummary
// QueueSummary show queue summary.
// Example:
/*
[
  {
    "action_id": "1dc27f71-1569-3ef7-f256-70674eb33f9e",
    "available": "0",
    "callers": "0",
    "event": "QueueSummary",
    "hold_time": "0",
    "logged_in": "0",
    "longest_hold_time": "0",
    "queue": "8001",
    "talk_time": "0"
  }
]
*/
func (c *AMICore) GetQueueSummary(ctx context.Context, queue string) ([]AMIResultRaw, error) {
	return QueueSummary(ctx, *c.Socket, queue)
}

// QueueMemberRingInUse
func (c *AMICore) QueueMemberRingInUse(ctx context.Context, _interface, ringInUse, queue string) (AMIResultRaw, error) {
	return QueueMemberRingInUse(ctx, *c.Socket, _interface, ringInUse, queue)
}

// QueueStatus
func (c *AMICore) QueueStatus(ctx context.Context, queue, member string) (AMIResultRaw, error) {
	return QueueStatus(ctx, *c.Socket, queue, member)
}

// QueueRule
func (c *AMICore) QueueRule(ctx context.Context, rule string) (AMIResultRaw, error) {
	return QueueRule(ctx, *c.Socket, rule)
}

// QueueReset
// QueueReset resets queue statistics.
func (c *AMICore) QueueReset(ctx context.Context, queue string) (AMIResultRaw, error) {
	return QueueReset(ctx, *c.Socket, queue)
}

// QueueRemove
// QueueRemove removes interface from queue.
func (c *AMICore) QueueRemove(ctx context.Context, queue AMIPayloadQueue) (AMIResultRaw, error) {
	return QueueRemove(ctx, *c.Socket, queue)
}

// QueueReload
// QueueReload reloads a queue, queues, or any sub-section of a queue or queues.
func (c *AMICore) QueueReload(ctx context.Context, queue AMIPayloadQueue) (AMIResultRaw, error) {
	return QueueReload(ctx, *c.Socket, queue)
}

// QueuePenalty
// QueuePenalty sets the penalty for a queue member.
func (c *AMICore) QueuePenalty(ctx context.Context, queue AMIPayloadQueue) (AMIResultRaw, error) {
	return QueuePenalty(ctx, *c.Socket, queue)
}

// QueuePause
// QueuePause makes a queue member temporarily unavailable.
func (c *AMICore) QueuePause(ctx context.Context, queue AMIPayloadQueue) (AMIResultRaw, error) {
	return QueuePause(ctx, *c.Socket, queue)
}

// QueueLog
// QueueLog adds custom entry in queue_log.
func (c *AMICore) QueueLog(ctx context.Context, queue AMIPayloadQueue) (AMIResultRaw, error) {
	return QueueLog(ctx, *c.Socket, queue)
}

// QueueAdd
// QueueAdd adds interface to queue.
func (c *AMICore) QueueAdd(ctx context.Context, queue AMIPayloadQueue) (AMIResultRaw, error) {
	return QueueAdd(ctx, *c.Socket, queue)
}

// GetExtensionStateList
// GetExtensionStateList list the current known extension states.
// Example:
/*
[
  {
    "action_id": "70a7ed51-dd26-154a-4c6e-9d683e6457cf",
    "context": "ext-local",
    "event": "ExtensionStatus",
    "exten": "5001",
    "hint": "SIP/5001&Custom:DND5001,CustomPresence:5001",
    "status": "0",
    "status_text": "Idle"
  }
]
*/
func (c *AMICore) ExtensionStateList(ctx context.Context) ([]AMIResultRaw, error) {
	return ExtensionStateList(ctx, *c.Socket)
}

// ExtensionState
// Example
/*
{
  "action_id": "35560f93-55c9-bd4f-6bce-b7f541819e5e",
  "context": "ext-local",
  "exten": "5001",
  "hint": "SIP/5001&Custom:DND5001,CustomPresence:5001",
  "message": "Extension Status",
  "response": "Success",
  "status": "0",
  "status_text": "Idle"
}
*/
func (c *AMICore) ExtensionState(ctx context.Context, exten, context string) (AMIResultRaw, error) {
	return ExtensionState(ctx, *c.Socket, exten, context)
}

// GetExtensionStates
func (c *AMICore) GetExtensionStates(ctx context.Context) ([]AMIResultRaw, error) {
	var extensions []AMIResultRaw
	response, err := ExtensionStateList(ctx, *c.Socket)
	switch {
	case err != nil:
		return nil, err
	case len(response) == 0:
		return nil, fmt.Errorf(config.AmiErrorNoExtensionsConfigured)
	default:
		for _, v := range response {
			extension, err := ExtensionState(ctx, *c.Socket, v.GetVal(config.AmiJsonFieldExten), v.GetVal(config.AmiJsonFieldContext))
			if err != nil {
				return nil, err
			}
			extensions = append(extensions, extension)
		}
	}

	return extensions, nil
}

// CoreShowChannels
func (c *AMICore) CoreShowChannels(ctx context.Context) ([]AMIResultRaw, error) {
	return CoreShowChannels(ctx, *c.Socket)
}

// AbsoluteTimeout
// Hangup a channel after a certain time. Acknowledges set time with Timeout Set message.
func (c *AMICore) AbsoluteTimeout(ctx context.Context, channel string, timeout int) (AMIResultRaw, error) {
	return AbsoluteTimeout(ctx, *c.Socket, channel, timeout)
}

// Hangup
// Hangup hangups channel.
func (c *AMICore) Hangup(ctx context.Context, channel, cause string) (AMIResultRaw, error) {
	return Hangup(ctx, *c.Socket, channel, cause)
}

// Originate
func (c *AMICore) Originate(ctx context.Context, originate AMIPayloadOriginate) (AMIResultRaw, error) {
	return Originate(ctx, *c.Socket, originate)
}

// ParkedCalls
func (c *AMICore) ParkedCalls(ctx context.Context) ([]AMIResultRaw, error) {
	return ParkedCalls(ctx, *c.Socket)
}

// Park
// Park parks a channel.
func (c *AMICore) Park(ctx context.Context, channel1, channel2 string, timeout int, parkinglot string) (AMIResultRaw, error) {
	return Park(ctx, *c.Socket, channel1, channel2, timeout, parkinglot)
}

// Parkinglots
func (c *AMICore) Parkinglots(ctx context.Context) ([]AMIResultRaw, error) {
	return Parkinglots(ctx, *c.Socket)
}

// PlayDTMF
// PlayDTMF plays DTMF signal on a specific channel.
func (c *AMICore) PlayDTMF(ctx context.Context, channel, digit string, duration int) (AMIResultRaw, error) {
	return PlayDTMF(ctx, *c.Socket, channel, digit, duration)
}

// Redirect
// Redirect redirects (transfer) a call.
func (c *AMICore) Redirect(ctx context.Context, call AMIPayloadCall) (AMIResultRaw, error) {
	return Redirect(ctx, *c.Socket, call)
}

// SendText
// SendText sends text message to channel.
func (c *AMICore) SendText(ctx context.Context, channel, message string) (AMIResultRaw, error) {
	return SendText(ctx, *c.Socket, channel, message)
}

// SetVar
// SetVar sets a channel variable. Sets a global or local channel variable.
// Note: If a channel name is not provided then the variable is global.
func (c *AMICore) SetVar(ctx context.Context, channel, variable, value string) (AMIResultRaw, error) {
	return SetVar(ctx, *c.Socket, channel, variable, value)
}

// GetStatus
// Status lists channel status.
// Will return the status information of each channel along with the value for the specified channel variables.
func (c *AMICore) GetStatus(ctx context.Context, channel, variables string) (AMIResultRaw, error) {
	return Status(ctx, *c.Socket, channel, variables)
}

// AOCMessage
// AOCMessage generates an Advice of Charge message on a channel.
func (c *AMICore) AOCMessage(ctx context.Context, aoc AMIPayloadAOC) (AMIResultRaw, error) {
	return AOCMessage(ctx, *c.Socket, aoc)
}

// GetVar
// GetVar get a channel variable.
func (c *AMICore) GetVar(ctx context.Context, channel, variable string) (AMIResultRaw, error) {
	return GetVar(ctx, *c.Socket, channel, variable)
}

// LocalOptimizeAway
// LocalOptimizeAway optimize away a local channel when possible.
// A local channel created with "/n" will not automatically optimize away.
// Calling this command on the local channel will clear that flag and allow it to optimize away if it's bridged or when it becomes bridged.
func (c *AMICore) LocalOptimizeAway(ctx context.Context, channel string) (AMIResultRaw, error) {
	return LocalOptimizeAway(ctx, *c.Socket, channel)
}

// MuteAudio
// MuteAudio mute an audio stream.
func (c *AMICore) MuteAudio(ctx context.Context, channel, direction string, state bool) (AMIResultRaw, error) {
	return MuteAudio(ctx, *c.Socket, channel, direction, state)
}

// GetAgents
// Agents lists agents and their status.
func (c *AMICore) GetAgents(ctx context.Context) ([]AMIResultRaw, error) {
	return Agents(ctx, *c.Socket)
}

// GetAgentLogoff
// AgentLogoff sets an agent as no longer logged in.
func (c *AMICore) GetAgentLogoff(ctx context.Context, agent string, soft bool) (AMIResultRaw, error) {
	return AgentLogoff(ctx, *c.Socket, agent, soft)
}

// AGI
// AGI add an AGI command to execute by Async AGI.
func (c *AMICore) AGI(ctx context.Context, channel, agiCommand, agiCommandID string) (AMIResultRaw, error) {
	return AGI(ctx, *c.Socket, channel, agiCommand, agiCommandID)
}

// ControlPlayback
// ControlPlayback control the playback of a file being played to a channel.
func (c *AMICore) ControlPlayback(ctx context.Context, channel string, control config.AGIControl) (AMIResultRaw, error) {
	return ControlPlayback(ctx, *c.Socket, channel, control)
}

// VoicemailRefresh
// VoicemailRefresh tell asterisk to poll mailboxes for a change.
func (c *AMICore) VoicemailRefresh(ctx context.Context, context, mailbox string) (AMIResultRaw, error) {
	return VoicemailRefresh(ctx, *c.Socket, context, mailbox)
}

// VoicemailUsersList
// VoicemailUsersList list all voicemail user information.
func (c *AMICore) VoicemailUsersList(ctx context.Context) ([]AMIResultRaw, error) {
	return VoicemailUsersList(ctx, *c.Socket)
}

// PresenceState
// PresenceState check presence state.
func (c *AMICore) PresenceState(ctx context.Context, provider string) (AMIResultRaw, error) {
	return PresenceState(ctx, *c.Socket, provider)
}

// PresenceStateList
// PresenceStateList list the current known presence states.
func (c *AMICore) PresenceStateList(ctx context.Context) ([]AMIResultRaw, error) {
	return PresenceStateList(ctx, *c.Socket)
}

// MailboxCount
// MailboxCount checks Mailbox Message Count.
func (c *AMICore) MailboxCount(ctx context.Context, mailbox string) (AMIResultRaw, error) {
	return MailboxCount(ctx, *c.Socket, mailbox)
}

// MailboxStatus
// MailboxStatus checks Mailbox Message Count.
func (c *AMICore) MailboxStatus(ctx context.Context, mailbox string) (AMIResultRaw, error) {
	return MailboxStatus(ctx, *c.Socket, mailbox)
}

// MWIDelete
// MWIDelete delete selected mailboxes.
func (c *AMICore) MWIDelete(ctx context.Context, mailbox string) (AMIResultRaw, error) {
	return MWIDelete(ctx, *c.Socket, mailbox)
}

// MWIGet
// MWIGet get selected mailboxes with message counts.
func (c *AMICore) MWIGet(ctx context.Context, mailbox string) (AMIResultRaw, error) {
	return MWIGet(ctx, *c.Socket, mailbox)
}

// MWIUpdate
// MWIUpdate update the mailbox message counts.
func (c *AMICore) MWIUpdate(ctx context.Context, mailbox, oldMessages, newMessages string) (AMIResultRaw, error) {
	return MWIUpdate(ctx, *c.Socket, mailbox, oldMessages, newMessages)
}
