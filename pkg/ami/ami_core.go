package ami

import (
	"context"
	"fmt"
	"log"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
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
	return SIPPeerStatusShort(ctx, *c.Socket, peer)
}

// HasSIPPeerStatus
func (c *AMICore) HasSIPPeerStatus(ctx context.Context, peer string) (bool, error) {
	return HasSIPPeerStatus(ctx, *c.Socket, peer)
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

// ExtensionStates
func (c *AMICore) ExtensionStates(ctx context.Context) ([]AMIResultRaw, error) {
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

func (c *AMICore) MakeCall(ctx context.Context, originate AMIPayloadOriginate) (AMIResultRaw, error) {
	return c.Originate(ctx, originate)
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

// MessageSend
// MessageSend send an out of call message to an endpoint.
func (c *AMICore) MessageSend(ctx context.Context, message AMIPayloadMessage) (AMIResultRaw, error) {
	return MessageSend(ctx, *c.Socket, message)
}

// KSendSMS
// KSendSMS sends a SMS using KHOMP device.
func (c *AMICore) KSendSMS(ctx context.Context, payload AMIPayloadKhompSMS) (AMIResultRaw, error) {
	return KSendSMS(ctx, *c.Socket, payload)
}

// IAXnetstats
// IAXnetstats show IAX channels network statistics.
func (c *AMICore) IAXnetstats(ctx context.Context) ([]AMIResultRaw, error) {
	return IAXnetstats(ctx, *c.Socket)
}

// IAXpeerlist
// IAXpeerlist show IAX channels network statistics.
func (c *AMICore) IAXpeerlist(ctx context.Context) ([]AMIResultRaw, error) {
	return IAXpeerlist(ctx, *c.Socket)
}

// IAXpeers
// IAXpeers list IAX peers.
func (c *AMICore) IAXpeers(ctx context.Context) ([]AMIResultRaw, error) {
	return IAXpeers(ctx, *c.Socket)
}

// IAXregistry
// IAXregistry show IAX registrations.
func (c *AMICore) IAXregistry(ctx context.Context) ([]AMIResultRaw, error) {
	return IAXregistry(ctx, *c.Socket)
}

// AddDialplanExtension
// AddDialplanExtension add an extension to the dialplan.
func (c *AMICore) AddDialplanExtension(ctx context.Context, extension AMIPayloadExtension) (AMIResultRaw, error) {
	return AddDialplanExtension(ctx, *c.Socket, extension)
}

// RemoveDialplanExtension
// RemoveDialplanExtension remove an extension from the dialplan.
func (c *AMICore) RemoveDialplanExtension(ctx context.Context, extension AMIPayloadExtension) (AMIResultRaw, error) {
	return RemoveDialplanExtension(ctx, *c.Socket, extension)
}

// Bridge
// Bridge bridges two channels already in the PBX.
func (c *AMICore) Bridge(ctx context.Context, channel1, channel2 string, tone string) (AMIResultRaw, error) {
	return Bridge(ctx, *c.Socket, channel1, channel2, tone)
}

// BlindTransfer
// BlindTransfer blind transfer channel(s) to the given destination.
func (c *AMICore) BlindTransfer(ctx context.Context, channel, context, extension string) (AMIResultRaw, error) {
	return BlindTransfer(ctx, *c.Socket, channel, context, extension)
}

// BridgeDestroy
// BridgeDestroy destroy a bridge.
func (c *AMICore) BridgeDestroy(ctx context.Context, bridgeUniqueId string) (AMIResultRaw, error) {
	return BridgeDestroy(ctx, *c.Socket, bridgeUniqueId)
}

// BridgeInfo
// BridgeInfo get information about a bridge.
func (c *AMICore) BridgeInfo(ctx context.Context, bridgeUniqueId string) (AMIResultRaw, error) {
	return BridgeInfo(ctx, *c.Socket, bridgeUniqueId)
}

// BridgeKick
// BridgeKick kick a channel from a bridge.
func (c *AMICore) BridgeKick(ctx context.Context, bridgeUniqueId, channel string) (AMIResultRaw, error) {
	return BridgeKick(ctx, *c.Socket, bridgeUniqueId, channel)
}

// BridgeList
// BridgeList get a list of bridges in the system.
func (c *AMICore) BridgeList(ctx context.Context, bridgeType string) (AMIResultRaw, error) {
	return BridgeList(ctx, *c.Socket, bridgeType)
}

// BridgeTechnologyList
// BridgeTechnologyList list available bridging technologies and their statuses.
func (c *AMICore) BridgeTechnologyList(ctx context.Context) ([]AMIResultRaw, error) {
	return BridgeTechnologyList(ctx, *c.Socket)
}

// BridgeTechnologySuspend
// BridgeTechnologySuspend suspend a bridging technology.
func (c *AMICore) BridgeTechnologySuspend(ctx context.Context, bridgeTechnology string) (AMIResultRaw, error) {
	return BridgeTechnologySuspend(ctx, *c.Socket, bridgeTechnology)
}

// BridgeTechnologyUnsuspend
// BridgeTechnologyUnsuspend unsuspend a bridging technology.
func (c *AMICore) BridgeTechnologyUnsuspend(ctx context.Context, bridgeTechnology string) (AMIResultRaw, error) {
	return BridgeTechnologyUnsuspend(ctx, *c.Socket, bridgeTechnology)
}

// DBDel
// DBDel Delete DB entry.
func (c *AMICore) DBDel(ctx context.Context, family, key string) (AMIResultRaw, error) {
	return DBDel(ctx, *c.Socket, family, key)
}

// DBDelTree
// DBDelTree delete DB tree.
func (c *AMICore) DBDelTree(ctx context.Context, family, key string) (AMIResultRaw, error) {
	return DBDelTree(ctx, *c.Socket, family, key)
}

// DBPut
// DBPut puts DB entry.
func (c *AMICore) DBPut(ctx context.Context, family, key, value string) (AMIResultRaw, error) {
	return DBPut(ctx, *c.Socket, family, key, value)
}

// DBGet
// DBGet gets DB Entry.
func (c *AMICore) DBGet(ctx context.Context, family, key string) ([]AMIResultRaw, error) {
	return DBGet(ctx, *c.Socket, family, key)
}

// PRIDebugFileSet
// PRIDebugFileSet set the file used for PRI debug message output.
func (c *AMICore) PRIDebugFileSet(ctx context.Context, filename string) (AMIResultRaw, error) {
	return PRIDebugFileSet(ctx, *c.Socket, filename)
}

// PRIDebugFileUnset
// PRIDebugFileUnset disables file output for PRI debug messages.
func (c *AMICore) PRIDebugFileUnset(ctx context.Context) (AMIResultRaw, error) {
	return PRIDebugFileUnset(ctx, *c.Socket)
}

// PRIDebugSet
// PRIDebugSet set PRI debug levels for a span.
func (c *AMICore) PRIDebugSet(ctx context.Context, span, level string) (AMIResultRaw, error) {
	return PRIDebugSet(ctx, *c.Socket, span, level)
}

// PRIShowSpans
// PRIShowSpans show status of PRI spans.
func (c *AMICore) PRIShowSpans(ctx context.Context, span string) ([]AMIResultRaw, error) {
	return PRIShowSpans(ctx, *c.Socket, span)
}

// SKINNYDevices
// SKINNYDevices lists SKINNY devices (text format).
// Lists Skinny devices in text format with details on current status.
// Devicelist will follow as separate events,
// followed by a final event called DevicelistComplete.
func (c *AMICore) SKINNYDevices(ctx context.Context) ([]AMIResultRaw, error) {
	return SKINNYDevices(ctx, *c.Socket)
}

// SKINNYLines
// SKINNYLines lists SKINNY lines (text format).
// Lists Skinny lines in text format with details on current status.
// Linelist will follow as separate events,
// followed by a final event called LinelistComplete.
func (c *AMICore) SKINNYLines(ctx context.Context) ([]AMIResultRaw, error) {
	return SKINNYLines(ctx, *c.Socket)
}

// SKINNYShowDevice
// SKINNYShowDevice show SKINNY device (text format).
// Show one SKINNY device with details on current status.
func (c *AMICore) SKINNYShowDevice(ctx context.Context, device string) (AMIResultRaw, error) {
	return SKINNYShowDevice(ctx, *c.Socket, device)
}

// SKINNYShowline
// SKINNYShowline shows SKINNY line (text format).
// Show one SKINNY line with details on current status.
func (c *AMICore) SKINNYShowline(ctx context.Context, line string) (AMIResultRaw, error) {
	return SKINNYShowline(ctx, *c.Socket, line)
}

// MeetMeList
// MeetMeList lists all users in a particular MeetMe conference.
// Will follow as separate events, followed by a final event called MeetmeListComplete.
func (c *AMICore) MeetMeList(ctx context.Context, conference string) ([]AMIResultRaw, error) {
	return MeetMeList(ctx, *c.Socket, conference)
}

// MeetMeMute
// MeetMeMute mute a Meetme user.
func (c *AMICore) MeetMeMute(ctx context.Context, meetme, userNumber string) (AMIResultRaw, error) {
	return MeetMeMute(ctx, *c.Socket, meetme, userNumber)
}

// MeetMeUnMute
// MeetMeUnMute unmute a Meetme user.
func (c *AMICore) MeetMeUnMute(ctx context.Context, meetme, userNumber string) (AMIResultRaw, error) {
	return MeetMeUnMute(ctx, *c.Socket, meetme, userNumber)
}

// MeetMeListRooms
// MeetMeListRooms list active conferences.
func (c *AMICore) MeetMeListRooms(ctx context.Context) ([]AMIResultRaw, error) {
	return MeetMeListRooms(ctx, *c.Socket)
}

// Monitor
// Monitor monitors a channel.
// This action may be used to record the audio on a specified channel.
func (c *AMICore) Monitor(ctx context.Context, payload AMIPayloadMonitor) (AMIResultRaw, error) {
	return Monitor(ctx, *c.Socket, payload)
}

// MonitorWith
func (c *AMICore) MonitorWith(ctx context.Context, channel, file, format string, mix bool) (AMIResultRaw, error) {
	return MonitorWith(ctx, *c.Socket, channel, file, format, mix)
}

// ChangeMonitor
// ChangeMonitor changes monitoring filename of a channel.
// This action may be used to change the file started by a previous 'Monitor' action.
func (c *AMICore) ChangeMonitor(ctx context.Context, payload AMIPayloadMonitor) (AMIResultRaw, error) {
	return ChangeMonitor(ctx, *c.Socket, payload)
}

// ChangeMonitorWith
// ChangeMonitor changes monitoring filename of a channel.
// This action may be used to change the file started by a previous 'Monitor' action.
func (c *AMICore) ChangeMonitorWith(ctx context.Context, channel, file string) (AMIResultRaw, error) {
	return ChangeMonitorWith(ctx, *c.Socket, channel, file)
}

// MixMonitor
// MixMonitor record a call and mix the audio during the recording.
func (c *AMICore) MixMonitor(ctx context.Context, payload AMIPayloadMonitor) (AMIResultRaw, error) {
	return MixMonitor(ctx, *c.Socket, payload)
}

// MixMonitorWith
// MixMonitor record a call and mix the audio during the recording.
func (c *AMICore) MixMonitorWith(ctx context.Context, channel, file, options, command string) (AMIResultRaw, error) {
	return MixMonitorWith(ctx, *c.Socket, channel, file, options, command)
}

// MixMonitorMute
// MixMonitorMute Mute / unMute a Mixmonitor recording.
// This action may be used to mute a MixMonitor recording.
func (c *AMICore) MixMonitorMute(ctx context.Context, channel, direction string, state bool) (AMIResultRaw, error) {
	return MixMonitorMute(ctx, *c.Socket, channel, direction, state)
}

// PauseMonitor
// PauseMonitor pauses monitoring of a channel.
// This action may be used to temporarily stop the recording of a channel.
func (c *AMICore) PauseMonitor(ctx context.Context, channel string) (AMIResultRaw, error) {
	return PauseMonitor(ctx, *c.Socket, channel)
}

// UnpauseMonitor
// UnpauseMonitor unpause monitoring of a channel.
// This action may be used to re-enable recording of a channel after calling PauseMonitor.
func (c *AMICore) UnpauseMonitor(ctx context.Context, channel string) (AMIResultRaw, error) {
	return UnpauseMonitor(ctx, *c.Socket, channel)
}

// StopMonitor
// StopMonitor stops monitoring a channel.
// This action may be used to end a previously started 'Monitor' action.
func (c *AMICore) StopMonitor(ctx context.Context, channel string) (AMIResultRaw, error) {
	return StopMonitor(ctx, *c.Socket, channel)
}

// StopMixMonitor
// StopMixMonitor stop recording a call through MixMonitor, and free the recording's file handle.
func (c *AMICore) StopMixMonitor(ctx context.Context, channel, mixMonitorId string) (AMIResultRaw, error) {
	return StopMixMonitor(ctx, *c.Socket, channel, mixMonitorId)
}

// PJSIPNotify
// PJSIPNotify send NOTIFY to either an endpoint, an arbitrary URI, or inside a SIP dialog.
func (c *AMICore) PJSIPNotify(ctx context.Context, endpoint, uri, variable string) (AMIResultRaw, error) {
	return PJSIPNotify(ctx, *c.Socket, endpoint, uri, variable)
}

// PJSIPQualify
// PJSIPQualify qualify a chan_pjsip endpoint.
func (c *AMICore) PJSIPQualify(ctx context.Context, endpoint string) (AMIResultRaw, error) {
	return PJSIPQualify(ctx, *c.Socket, endpoint)
}

// PJSIPRegister
// PJSIPRegister register an outbound registration.
func (c *AMICore) PJSIPRegister(ctx context.Context, registration string) (AMIResultRaw, error) {
	return PJSIPRegister(ctx, *c.Socket, registration)
}

// PJSIPUnregister
// PJSIPUnregister unregister an outbound registration.
func (c *AMICore) PJSIPUnregister(ctx context.Context, registration string) (AMIResultRaw, error) {
	return PJSIPUnregister(ctx, *c.Socket, registration)
}

// PJSIPShowEndpoint
// PJSIPShowEndpoint detail listing of an endpoint and its objects.
func (c *AMICore) PJSIPShowEndpoint(ctx context.Context, endpoint string) ([]AMIResultRaw, error) {
	return PJSIPShowEndpoint(ctx, *c.Socket, endpoint)
}

// PJSIPShowEndpoints
// PJSIPShowEndpoints list pjsip endpoints.
func (c *AMICore) PJSIPShowEndpoints(ctx context.Context) ([]AMIResultRaw, error) {
	return PJSIPShowEndpoints(ctx, *c.Socket)
}

// PJSIPShowRegistrationInboundContactStatuses
func (c *AMICore) PJSIPShowRegistrationInboundContactStatuses(ctx context.Context) ([]AMIResultRaw, error) {
	return PJSIPShowRegistrationInboundContactStatuses(ctx, *c.Socket)
}

// PJSIPShowRegistrationsInbound
// PJSIPShowRegistrationsInbound lists PJSIP inbound registrations.
func (c *AMICore) PJSIPShowRegistrationsInbound(ctx context.Context) ([]AMIResultRaw, error) {
	return PJSIPShowRegistrationsInbound(ctx, *c.Socket)
}

// PJSIPShowRegistrationsOutbound
// PJSIPShowRegistrationsOutbound lists PJSIP outbound registrations.
func (c *AMICore) PJSIPShowRegistrationsOutbound(ctx context.Context) ([]AMIResultRaw, error) {
	return PJSIPShowRegistrationsOutbound(ctx, *c.Socket)
}

// PJSIPShowResourceLists
// PJSIPShowResourceLists displays settings for configured resource lists.
func (c *AMICore) PJSIPShowResourceLists(ctx context.Context) ([]AMIResultRaw, error) {
	return PJSIPShowResourceLists(ctx, *c.Socket)
}

// PJSIPShowSubscriptionsInbound
// PJSIPShowSubscriptionsInbound list of inbound subscriptions.
func (c *AMICore) PJSIPShowSubscriptionsInbound(ctx context.Context) ([]AMIResultRaw, error) {
	return PJSIPShowSubscriptionsInbound(ctx, *c.Socket)
}

// PJSIPShowSubscriptionsOutbound
// PJSIPShowSubscriptionsOutbound list of outbound subscriptions.
func (c *AMICore) PJSIPShowSubscriptionsOutbound(ctx context.Context) ([]AMIResultRaw, error) {
	return PJSIPShowSubscriptionsOutbound(ctx, *c.Socket)
}

// FAXSession
// FAXSession responds with a detailed description of a single FAX session.
func (c *AMICore) FAXSession(ctx context.Context, sessionNumber string) (AMIResultRaw, error) {
	return FAXSession(ctx, *c.Socket, sessionNumber)
}

// FAXSessions
// FAXSessions list active FAX sessions.
func (c *AMICore) FAXSessions(ctx context.Context) ([]AMIResultRaw, error) {
	return FAXSessions(ctx, *c.Socket)
}

// FAXStats
// FAXStats responds with fax statistics.
func (c *AMICore) FAXStats(ctx context.Context) (AMIResultRaw, error) {
	return FAXStats(ctx, *c.Socket)
}

// Atxfer
// Atxfer attended transfer.
func (c *AMICore) Atxfer(ctx context.Context, channel, extension, context string) (AMIResultRaw, error) {
	return Atxfer(ctx, *c.Socket, channel, extension, context)
}

// CancelAtxfer
// CancelAtxfer cancel an attended transfer.
func (c *AMICore) CancelAtxfer(ctx context.Context, channel string) (AMIResultRaw, error) {
	return CancelAtxfer(ctx, *c.Socket, channel)
}

// DAHDIDialOffhook
// DAHDIDialOffhook dials over DAHDI channel while offhook.
// Generate DTMF control frames to the bridged peer.
func (c *AMICore) DAHDIDialOffhook(ctx context.Context, channel, number string) (AMIResultRaw, error) {
	return DAHDIDialOffhook(ctx, *c.Socket, channel, number)
}

// DAHDIDNDoff
// DAHDIDNDoff toggles DAHDI channel Do Not Disturb status OFF.
func (c *AMICore) DAHDIDNDoff(ctx context.Context, channel string) (AMIResultRaw, error) {
	return DAHDIDNDoff(ctx, *c.Socket, channel)
}

// DAHDIDNDon
// DAHDIDNDon toggles DAHDI channel Do Not Disturb status ON.
func (c *AMICore) DAHDIDNDon(ctx context.Context, channel string) (AMIResultRaw, error) {
	return DAHDIDNDon(ctx, *c.Socket, channel)
}

// DAHDIHangup
// DAHDIHangup hangups DAHDI Channel.
func (c *AMICore) DAHDIHangup(ctx context.Context, channel string) (AMIResultRaw, error) {
	return DAHDIHangup(ctx, *c.Socket, channel)
}

// DAHDIRestart
// DAHDIRestart fully Restart DAHDI channels (terminates calls).
func (c *AMICore) DAHDIRestart(ctx context.Context) (AMIResultRaw, error) {
	return DAHDIRestart(ctx, *c.Socket)
}

// DAHDIShowChannels
// DAHDIShowChannels show status of DAHDI channels.
func (c *AMICore) DAHDIShowChannels(ctx context.Context, channel string) ([]AMIResultRaw, error) {
	return DAHDIShowChannels(ctx, *c.Socket, channel)
}

// DAHDITransfer
// DAHDITransfer transfers DAHDI Channel.
func (c *AMICore) DAHDITransfer(ctx context.Context, channel string) (AMIResultRaw, error) {
	return DAHDITransfer(ctx, *c.Socket, channel)
}

// ConfbridgeList
// ConfbridgeList lists all users in a particular ConfBridge conference.
func (c *AMICore) ConfbridgeList(ctx context.Context, conference string) ([]AMIResultRaw, error) {
	return ConfbridgeList(ctx, *c.Socket, conference)
}

// ConfbridgeListRooms
// ConfbridgeListRooms lists data about all active conferences.
func (c *AMICore) ConfbridgeListRooms(ctx context.Context) ([]AMIResultRaw, error) {
	return ConfbridgeListRooms(ctx, *c.Socket)
}

// ConfbridgeMute
// ConfbridgeMute mutes a specified user in a specified conference.
func (c *AMICore) ConfbridgeMute(ctx context.Context, conference string, channel string) (AMIResultRaw, error) {
	return ConfbridgeMute(ctx, *c.Socket, conference, channel)
}

// ConfbridgeUnmute
// ConfbridgeUnmute unmute a specified user in a specified conference.
func (c *AMICore) ConfbridgeUnmute(ctx context.Context, conference string, channel string) (AMIResultRaw, error) {
	return ConfbridgeUnmute(ctx, *c.Socket, conference, channel)
}

// ConfbridgeKick
// ConfbridgeKick removes a specified user from a specified conference.
func (c *AMICore) ConfbridgeKick(ctx context.Context, conference string, channel string) (AMIResultRaw, error) {
	return ConfbridgeKick(ctx, *c.Socket, conference, channel)
}

// ConfbridgeLock
// ConfbridgeLock locks a specified conference.
func (c *AMICore) ConfbridgeLock(ctx context.Context, conference string, channel string) (AMIResultRaw, error) {
	return ConfbridgeLock(ctx, *c.Socket, conference, channel)
}

// ConfbridgeUnlock
// ConfbridgeUnlock unlocks a specified conference.
func (c *AMICore) ConfbridgeUnlock(ctx context.Context, conference string, channel string) (AMIResultRaw, error) {
	return ConfbridgeUnlock(ctx, *c.Socket, config.AmiErrorLoginFailed, channel)
}

// ConfbridgeSetSingleVideoSrc
// ConfbridgeSetSingleVideoSrc sets a conference user as the single video source distributed to all other video-capable participants.
func (c *AMICore) ConfbridgeSetSingleVideoSrc(ctx context.Context, conference string, channel string) (AMIResultRaw, error) {
	return ConfbridgeSetSingleVideoSrc(ctx, *c.Socket, conference, channel)
}

// ConfbridgeStartRecord
// ConfbridgeStartRecord starts a recording in the context of given conference and creates a file with the name specified by recordFile
func (c *AMICore) ConfbridgeStartRecord(ctx context.Context, conference string, recordFile string) (AMIResultRaw, error) {
	return ConfbridgeStartRecord(ctx, *c.Socket, conference, recordFile)
}

// ConfbridgeStopRecord
// ConfbridgeStopRecord stops a recording pertaining to the given conference
func (c *AMICore) ConfbridgeStopRecord(ctx context.Context, conference string) (AMIResultRaw, error) {
	return ConfbridgeStopRecord(ctx, *c.Socket, conference)
}

// MakeOutboundCall
func (c *AMICore) MakeOutboundCall(ctx context.Context, d AMIOriginateDirection) (AMIResultRaw, bool, error) {
	return MakeOutboundCall(ctx, *c.Socket, d)
}

// MakeInternalCall
func (c *AMICore) MakeInternalCall(ctx context.Context, d AMIOriginateDirection) (AMIResultRaw, bool, error) {
	return MakeInternalCall(ctx, *c.Socket, d)
}

// Chanspy
func (c *AMICore) Chanspy(ctx context.Context, ch AMIPayloadChanspy) (AMIResultRawLevel, error) {
	return Chanspy(ctx, *c.Socket, ch)
}
