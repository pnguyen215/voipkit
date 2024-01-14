package ami

import (
	"context"
	"fmt"
	"log"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

func NewCore() *AMICore {
	c := &AMICore{}
	c.SetEvent(make(chan AmiReply))
	c.SetStop(make(chan struct{}))
	c.SetSocket(NewAmiSocket())
	return c
}

func (c *AMICore) SetSocket(socket *AMISocket) *AMICore {
	c.socket = socket
	return c
}

func (c *AMICore) SetUUID(id string) *AMICore {
	c.UUID = id
	return c
}

func (c *AMICore) SetEvent(event chan AmiReply) *AMICore {
	c.event = event
	return c
}

func (c *AMICore) SetStop(stop chan struct{}) *AMICore {
	c.stop = stop
	return c
}

func (c *AMICore) SetDictionary(dictionary *AMIDictionary) *AMICore {
	c.Dictionary = dictionary
	return c
}

func (c *AMICore) AddSession() *AMICore {
	c.SetUUID(GenUUIDShorten())
	c.socket.SetUUID(c.UUID)
	return c
}

// WithCore
// Creating new instance asterisk server connection
// Firstly, create new instance AMISocket
// Secondly, create new request body to login
func WithCore(ctx context.Context, socket *AMISocket, auth *AMIAuth) (*AMICore, error) {
	uuid, err := GenUUID()
	if err != nil {
		return nil, err
	}
	socket.SetUUID(uuid)
	err = WithAuthenticate(ctx, *socket, auth)
	if err != nil {
		return nil, err
	}
	core := NewCore()
	core.SetSocket(socket)
	core.SetUUID(uuid)
	core.SetDictionary(socket.Dictionary)
	core.wg.Add(1)
	go core.run(ctx)
	return core, nil
}

// run
// Go-func to consume event from asterisk server response
func (c *AMICore) run(ctx context.Context) {
	defer c.wg.Done()
	for {
		select {
		case <-c.stop:
			return
		case <-ctx.Done():
			return
		default:
			event, err := Events(ctx, *c.socket)
			if err != nil {
				log.Printf(config.AmiErrorConsumeEvent, err)
				return
			}
			c.event <- event
		}
	}
}

// Events
// Consume all events will be returned an channel received from asterisk server log.
func (c *AMICore) Events() <-chan AmiReply {
	return c.event
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
func (c *AMICore) GetSIPPeers(ctx context.Context) ([]AmiReply, error) {
	var peers []AmiReply
	response, err := SIPPeers(ctx, *c.socket)
	switch {
	case err != nil:
		return nil, err
	case len(response) == 0:
		return nil, fmt.Errorf(config.AmiErrorNoExtensionConfigured)
	default:
		for _, v := range response {
			peer, err := SIPShowPeer(ctx, *c.socket, v.Get(config.AmiJsonFieldObjectName))
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
func (c *AMICore) GetSIPPeer(ctx context.Context, peer string) (AmiReply, error) {
	return SIPShowPeer(ctx, *c.socket, peer)
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
func (c *AMICore) GetSIPPeersStatus(ctx context.Context) ([]AmiReply, error) {
	var peers []AmiReply
	response, err := SIPPeers(ctx, *c.socket)
	switch {
	case err != nil:
		return nil, err
	case len(response) == 0:
		return nil, fmt.Errorf(config.AmiErrorNoExtensionConfigured)
	default:
		for _, v := range response {
			peer, err := SIPPeerStatus(ctx, *c.socket, v.Get(config.AmiJsonFieldObjectName))
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
func (c *AMICore) GetSIPPeerStatus(ctx context.Context, peer string) (AmiReply, error) {
	return SIPPeerStatusShort(ctx, *c.socket, peer)
}

// HasSIPPeerStatus
func (c *AMICore) HasSIPPeerStatus(ctx context.Context, peer string) (bool, error) {
	return SIPPeerStatusExists(ctx, *c.socket, peer)
}

// GetSIPShowRegistry
// Example:
/*

 */
func (c *AMICore) GetSIPShowRegistry(ctx context.Context) ([]AmiReply, error) {
	return SIPShowRegistry(ctx, *c.socket)
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
func (c *AMICore) GetSIPQualifyPeer(ctx context.Context) ([]AmiReply, error) {
	var peers []AmiReply
	response, err := SIPPeers(ctx, *c.socket)
	switch {
	case err != nil:
		return nil, err
	case len(response) == 0:
		return nil, fmt.Errorf(config.AmiErrorNoExtensionConfigured)
	default:
		for _, v := range response {
			peer, err := SIPQualifyPeer(ctx, *c.socket, v.Get(config.AmiJsonFieldObjectName))
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
	close(c.stop)
	c.wg.Wait()
	return Logoff(ctx, *c.socket)
}

// Ping
func (c *AMICore) Ping(ctx context.Context) error {
	return Ping(ctx, *c.socket)
}

// Command executes an Asterisk CLI Command.
func (c *AMICore) Command(ctx context.Context, cmd string) (AmiReplies, error) {
	return Command(ctx, *c.socket, cmd)
}

// CoreSettings shows PBX core settings (version etc).
func (c *AMICore) GetCoreSettings(ctx context.Context) (AmiReply, error) {
	return CoreSettings(ctx, *c.socket)
}

// CoreStatus shows PBX core status variables.
func (c *AMICore) GetCoreStatus(ctx context.Context) (AmiReply, error) {
	return CoreStatus(ctx, *c.socket)
}

// ListCommands lists available manager commands.
// Returns the action name and synopsis for every action that is available to the user
func (c *AMICore) GetListCommands(ctx context.Context) (AmiReply, error) {
	return ListCommands(ctx, *c.socket)
}

// Challenge generates a challenge for MD5 authentication.
func (c *AMICore) Challenge(ctx context.Context) (AmiReply, error) {
	return Challenge(ctx, *c.socket)
}

// CreateConfig creates an empty file in the configuration directory.
// This action will create an empty file in the configuration directory.
// This action is intended to be used before an UpdateConfig action.
func (c *AMICore) CreateConfig(ctx context.Context, filename string) (AmiReply, error) {
	return CreateConfig(ctx, *c.socket, filename)
}

// DataGet retrieves the data api tree.
func (c *AMICore) DataGet(ctx context.Context, path, search, filter string) (AmiReply, error) {
	return DataGet(ctx, *c.socket, path, search, filter)
}

// EventFlow control Event Flow.
// eventMask: Enable/Disable sending of events to this manager client.
func (c *AMICore) EventFlow(ctx context.Context, eventMask string) (AmiReply, error) {
	return EventFlow(ctx, *c.socket, eventMask)
}

// GetConfig retrieves configuration.
// This action will dump the contents of a configuration file by category and contents or optionally by specified category only.
func (c *AMICore) GetConfig(ctx context.Context, filename, category, filter string) (AmiReply, error) {
	return GetConfig(ctx, *c.socket, filename, category, filter)
}

// GetConfigJson retrieves configuration (JSON format).
// This action will dump the contents of a configuration file by category and contents in JSON format.
// This only makes sense to be used using raw man over the HTTP interface.
func (c *AMICore) GetConfigJson(ctx context.Context, filename, category, filter string) (AmiReply, error) {
	return GetConfigJson(ctx, *c.socket, filename, category, filter)
}

// JabberSend sends a message to a Jabber Client
func (c *AMICore) JabberSend(ctx context.Context, jabber, jid, message string) (AmiReply, error) {
	return JabberSend(ctx, *c.socket, jabber, jid, message)
}

// ListCategories lists categories in configuration file.
// Example:
// filename like: manager.conf, extensions.conf, sip.conf...
func (c *AMICore) ListCategories(ctx context.Context, filename string) (AmiReply, error) {
	return ListCategories(ctx, *c.socket, filename)
}

// ModuleCheck checks if module is loaded.
// Checks if Asterisk module is loaded. Will return Success/Failure. For success returns, the module revision number is included.
func (c *AMICore) ModuleCheck(ctx context.Context, module string) (AmiReply, error) {
	return ModuleCheck(ctx, *c.socket, module)
}

// ModuleLoad module management.
// Loads, unloads or reloads an Asterisk module in a running system.
func (c *AMICore) ModuleLoad(ctx context.Context, module, loadType string) (AmiReply, error) {
	return ModuleLoad(ctx, *c.socket, module, loadType)
}

// Reload Sends a reload event.
func (c *AMICore) Reload(ctx context.Context, module string) (AmiReply, error) {
	return Reload(ctx, *c.socket, module)
}

// ShowDialPlan shows dialplan contexts and extensions
// Be aware that showing the full dialplan may take a lot of capacity.
func (c *AMICore) ShowDialPlan(ctx context.Context, extension, context string) ([]AmiReply, error) {
	return ShowDialPlan(ctx, *c.socket, extension, context)
}

// Filter dynamically add filters for the current manager session.
func (c *AMICore) Filter(ctx context.Context, operation, filter string) (AmiReply, error) {
	return Filter(ctx, *c.socket, operation, filter)
}

// DeviceStateList list the current known device states.
func (c *AMICore) GetDeviceStateList(ctx context.Context) ([]AmiReply, error) {
	return DeviceStateList(ctx, *c.socket)
}

// LoggerRotate reload and rotate the Asterisk logger.
func (c *AMICore) LoggerRotate(ctx context.Context) (AmiReply, error) {
	return LoggerRotate(ctx, *c.socket)
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
func (c *AMICore) UpdateConfig(ctx context.Context, sourceFilename, destinationFilename string, reload bool, actions ...AMIUpdateConfigAction) (AmiReply, error) {
	return UpdateConfig(ctx, *c.socket, sourceFilename, destinationFilename, reload, actions...)
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
func (c *AMICore) GetQueueStatuses(ctx context.Context, queue string) ([]AmiReply, error) {
	return QueueStatuses(ctx, *c.socket, queue)
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
func (c *AMICore) GetQueueSummary(ctx context.Context, queue string) ([]AmiReply, error) {
	return QueueSummary(ctx, *c.socket, queue)
}

// QueueMemberRingInUse
func (c *AMICore) QueueMemberRingInUse(ctx context.Context, _interface, ringInUse, queue string) (AmiReply, error) {
	return QueueMemberRingInUse(ctx, *c.socket, _interface, ringInUse, queue)
}

// QueueStatus
func (c *AMICore) QueueStatus(ctx context.Context, queue, member string) (AmiReply, error) {
	return QueueStatus(ctx, *c.socket, queue, member)
}

// QueueRule
func (c *AMICore) QueueRule(ctx context.Context, rule string) (AmiReply, error) {
	return QueueRule(ctx, *c.socket, rule)
}

// QueueReset
// QueueReset resets queue statistics.
func (c *AMICore) QueueReset(ctx context.Context, queue string) (AmiReply, error) {
	return QueueReset(ctx, *c.socket, queue)
}

// QueueRemove
// QueueRemove removes interface from queue.
func (c *AMICore) QueueRemove(ctx context.Context, queue AMIPayloadQueue) (AmiReply, error) {
	return QueueRemove(ctx, *c.socket, queue)
}

// QueueReload
// QueueReload reloads a queue, queues, or any sub-section of a queue or queues.
func (c *AMICore) QueueReload(ctx context.Context, queue AMIPayloadQueue) (AmiReply, error) {
	return QueueReload(ctx, *c.socket, queue)
}

// QueuePenalty
// QueuePenalty sets the penalty for a queue member.
func (c *AMICore) QueuePenalty(ctx context.Context, queue AMIPayloadQueue) (AmiReply, error) {
	return QueuePenalty(ctx, *c.socket, queue)
}

// QueuePause
// QueuePause makes a queue member temporarily unavailable.
func (c *AMICore) QueuePause(ctx context.Context, queue AMIPayloadQueue) (AmiReply, error) {
	return QueuePause(ctx, *c.socket, queue)
}

// QueueLog
// QueueLog adds custom entry in queue_log.
func (c *AMICore) QueueLog(ctx context.Context, queue AMIPayloadQueue) (AmiReply, error) {
	return QueueLog(ctx, *c.socket, queue)
}

// QueueAdd
// QueueAdd adds interface to queue.
func (c *AMICore) QueueAdd(ctx context.Context, queue AMIPayloadQueue) (AmiReply, error) {
	return QueueAdd(ctx, *c.socket, queue)
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
func (c *AMICore) ExtensionStateList(ctx context.Context) ([]AmiReply, error) {
	return ExtensionStateList(ctx, *c.socket)
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
func (c *AMICore) ExtensionState(ctx context.Context, exten, context string) (AmiReply, error) {
	return ExtensionState(ctx, *c.socket, exten, context)
}

// ExtensionStates
func (c *AMICore) ExtensionStates(ctx context.Context) ([]AmiReply, error) {
	var extensions []AmiReply
	response, err := ExtensionStateList(ctx, *c.socket)
	switch {
	case err != nil:
		return nil, err
	case len(response) == 0:
		return nil, fmt.Errorf(config.AmiErrorNoExtensionsConfigured)
	default:
		for _, v := range response {
			extension, err := ExtensionState(ctx, *c.socket, v.Get(config.AmiJsonFieldExten), v.Get(config.AmiJsonFieldContext))
			if err != nil {
				return nil, err
			}
			extensions = append(extensions, extension)
		}
	}

	return extensions, nil
}

// CoreShowChannels
func (c *AMICore) CoreShowChannels(ctx context.Context) ([]AmiReply, error) {
	return CoreShowChannels(ctx, *c.socket)
}

// AbsoluteTimeout
// Hangup a channel after a certain time. Acknowledges set time with Timeout Set message.
func (c *AMICore) AbsoluteTimeout(ctx context.Context, channel string, timeout int) (AmiReply, error) {
	return AbsoluteTimeout(ctx, *c.socket, channel, timeout)
}

// Hangup
// Hangup hangups channel.
func (c *AMICore) Hangup(ctx context.Context, channel, cause string) (AmiReply, error) {
	return Hangup(ctx, *c.socket, channel, cause)
}

// Originate
func (c *AMICore) Originate(ctx context.Context, originate AMIOriginate) (AmiReply, error) {
	return Originate(ctx, *c.socket, originate)
}

func (c *AMICore) MakeCall(ctx context.Context, originate AMIOriginate) (AmiReply, error) {
	return c.Originate(ctx, originate)
}

// ParkedCalls
func (c *AMICore) ParkedCalls(ctx context.Context) ([]AmiReply, error) {
	return ParkedCalls(ctx, *c.socket)
}

// Park
// Park parks a channel.
func (c *AMICore) Park(ctx context.Context, channel1, channel2 string, timeout int, parkinglot string) (AmiReply, error) {
	return Park(ctx, *c.socket, channel1, channel2, timeout, parkinglot)
}

// Parkinglots
func (c *AMICore) Parkinglots(ctx context.Context) ([]AmiReply, error) {
	return Parkinglots(ctx, *c.socket)
}

// PlayDTMF
// PlayDTMF plays DTMF signal on a specific channel.
func (c *AMICore) PlayDTMF(ctx context.Context, channel, digit string, duration int) (AmiReply, error) {
	return PlayDTMF(ctx, *c.socket, channel, digit, duration)
}

// Redirect
// Redirect redirects (transfer) a call.
func (c *AMICore) Redirect(ctx context.Context, call AMIPayloadCall) (AmiReply, error) {
	return Redirect(ctx, *c.socket, call)
}

// SendText
// SendText sends text message to channel.
func (c *AMICore) SendText(ctx context.Context, channel, message string) (AmiReply, error) {
	return SendText(ctx, *c.socket, channel, message)
}

// SetVar
// SetVar sets a channel variable. Sets a global or local channel variable.
// Note: If a channel name is not provided then the variable is global.
func (c *AMICore) SetVar(ctx context.Context, channel, variable, value string) (AmiReply, error) {
	return SetVar(ctx, *c.socket, channel, variable, value)
}

// GetStatus
// Status lists channel status.
// Will return the status information of each channel along with the value for the specified channel variables.
func (c *AMICore) GetStatus(ctx context.Context, channel, variables string) (AmiReply, error) {
	return Status(ctx, *c.socket, channel, variables)
}

// AOCMessage
// AOCMessage generates an Advice of Charge message on a channel.
func (c *AMICore) AOCMessage(ctx context.Context, aoc AMIPayloadAOC) (AmiReply, error) {
	return AOCMessage(ctx, *c.socket, aoc)
}

// GetVar
// GetVar get a channel variable.
func (c *AMICore) GetVar(ctx context.Context, channel, variable string) (AmiReply, error) {
	return GetVar(ctx, *c.socket, channel, variable)
}

// LocalOptimizeAway
// LocalOptimizeAway optimize away a local channel when possible.
// A local channel created with "/n" will not automatically optimize away.
// Calling this command on the local channel will clear that flag and allow it to optimize away if it's bridged or when it becomes bridged.
func (c *AMICore) LocalOptimizeAway(ctx context.Context, channel string) (AmiReply, error) {
	return LocalOptimizeAway(ctx, *c.socket, channel)
}

// MuteAudio
// MuteAudio mute an audio stream.
func (c *AMICore) MuteAudio(ctx context.Context, channel, direction string, state bool) (AmiReply, error) {
	return MuteAudio(ctx, *c.socket, channel, direction, state)
}

// GetAgents
// Agents lists agents and their status.
func (c *AMICore) GetAgents(ctx context.Context) ([]AmiReply, error) {
	return Agents(ctx, *c.socket)
}

// GetAgentLogoff
// AgentLogoff sets an agent as no longer logged in.
func (c *AMICore) GetAgentLogoff(ctx context.Context, agent string, soft bool) (AmiReply, error) {
	return AgentLogoff(ctx, *c.socket, agent, soft)
}

// AGI
// AGI add an AGI command to execute by Async AGI.
func (c *AMICore) AGI(ctx context.Context, channel, agiCommand, agiCommandID string) (AmiReply, error) {
	return AGI(ctx, *c.socket, channel, agiCommand, agiCommandID)
}

// ControlPlayback
// ControlPlayback control the playback of a file being played to a channel.
func (c *AMICore) ControlPlayback(ctx context.Context, channel string, control config.AGIControl) (AmiReply, error) {
	return ControlPlayback(ctx, *c.socket, channel, control)
}

// VoicemailRefresh
// VoicemailRefresh tell asterisk to poll mailboxes for a change.
func (c *AMICore) VoicemailRefresh(ctx context.Context, context, mailbox string) (AmiReply, error) {
	return VoicemailRefresh(ctx, *c.socket, context, mailbox)
}

// VoicemailUsersList
// VoicemailUsersList list all voicemail user information.
func (c *AMICore) VoicemailUsersList(ctx context.Context) ([]AmiReply, error) {
	return VoicemailUsersList(ctx, *c.socket)
}

// PresenceState
// PresenceState check presence state.
func (c *AMICore) PresenceState(ctx context.Context, provider string) (AmiReply, error) {
	return PresenceState(ctx, *c.socket, provider)
}

// PresenceStateList
// PresenceStateList list the current known presence states.
func (c *AMICore) PresenceStateList(ctx context.Context) ([]AmiReply, error) {
	return PresenceStateList(ctx, *c.socket)
}

// MailboxCount
// MailboxCount checks Mailbox Message Count.
func (c *AMICore) MailboxCount(ctx context.Context, mailbox string) (AmiReply, error) {
	return MailboxCount(ctx, *c.socket, mailbox)
}

// MailboxStatus
// MailboxStatus checks Mailbox Message Count.
func (c *AMICore) MailboxStatus(ctx context.Context, mailbox string) (AmiReply, error) {
	return MailboxStatus(ctx, *c.socket, mailbox)
}

// MWIDelete
// MWIDelete delete selected mailboxes.
func (c *AMICore) MWIDelete(ctx context.Context, mailbox string) (AmiReply, error) {
	return MWIDelete(ctx, *c.socket, mailbox)
}

// MWIGet
// MWIGet get selected mailboxes with message counts.
func (c *AMICore) MWIGet(ctx context.Context, mailbox string) (AmiReply, error) {
	return MWIGet(ctx, *c.socket, mailbox)
}

// MWIUpdate
// MWIUpdate update the mailbox message counts.
func (c *AMICore) MWIUpdate(ctx context.Context, mailbox, oldMessages, newMessages string) (AmiReply, error) {
	return MWIUpdate(ctx, *c.socket, mailbox, oldMessages, newMessages)
}

// MessageSend
// MessageSend send an out of call message to an endpoint.
func (c *AMICore) MessageSend(ctx context.Context, message AMIPayloadMessage) (AmiReply, error) {
	return MessageSend(ctx, *c.socket, message)
}

// KSendSMS
// KSendSMS sends a SMS using KHOMP device.
func (c *AMICore) KSendSMS(ctx context.Context, payload AMIPayloadKhompSMS) (AmiReply, error) {
	return KSendSMS(ctx, *c.socket, payload)
}

// IAXnetstats
// IAXnetstats show IAX channels network statistics.
func (c *AMICore) IAXnetstats(ctx context.Context) ([]AmiReply, error) {
	return IAXnetstats(ctx, *c.socket)
}

// IAXpeerlist
// IAXpeerlist show IAX channels network statistics.
func (c *AMICore) IAXpeerlist(ctx context.Context) ([]AmiReply, error) {
	return IAXpeerlist(ctx, *c.socket)
}

// IAXpeers
// IAXpeers list IAX peers.
func (c *AMICore) IAXpeers(ctx context.Context) ([]AmiReply, error) {
	return IAXpeers(ctx, *c.socket)
}

// IAXregistry
// IAXregistry show IAX registrations.
func (c *AMICore) IAXregistry(ctx context.Context) ([]AmiReply, error) {
	return IAXregistry(ctx, *c.socket)
}

// AddDialplanExtension
// AddDialplanExtension add an extension to the dialplan.
func (c *AMICore) AddDialplanExtension(ctx context.Context, extension AMIPayloadExtension) (AmiReply, error) {
	return AddDialplanExtension(ctx, *c.socket, extension)
}

// RemoveDialplanExtension
// RemoveDialplanExtension remove an extension from the dialplan.
func (c *AMICore) RemoveDialplanExtension(ctx context.Context, extension AMIPayloadExtension) (AmiReply, error) {
	return RemoveDialplanExtension(ctx, *c.socket, extension)
}

// Bridge
// Bridge bridges two channels already in the PBX.
func (c *AMICore) Bridge(ctx context.Context, channel1, channel2 string, tone string) (AmiReply, error) {
	return Bridge(ctx, *c.socket, channel1, channel2, tone)
}

// BlindTransfer
// BlindTransfer blind transfer channel(s) to the given destination.
func (c *AMICore) BlindTransfer(ctx context.Context, channel, context, extension string) (AmiReply, error) {
	return BlindTransfer(ctx, *c.socket, channel, context, extension)
}

// BridgeDestroy
// BridgeDestroy destroy a bridge.
func (c *AMICore) BridgeDestroy(ctx context.Context, bridgeUniqueId string) (AmiReply, error) {
	return BridgeDestroy(ctx, *c.socket, bridgeUniqueId)
}

// BridgeInfo
// BridgeInfo get information about a bridge.
func (c *AMICore) BridgeInfo(ctx context.Context, bridgeUniqueId string) (AmiReply, error) {
	return BridgeInfo(ctx, *c.socket, bridgeUniqueId)
}

// BridgeKick
// BridgeKick kick a channel from a bridge.
func (c *AMICore) BridgeKick(ctx context.Context, bridgeUniqueId, channel string) (AmiReply, error) {
	return BridgeKick(ctx, *c.socket, bridgeUniqueId, channel)
}

// BridgeList
// BridgeList get a list of bridges in the system.
func (c *AMICore) BridgeList(ctx context.Context, bridgeType string) (AmiReply, error) {
	return BridgeList(ctx, *c.socket, bridgeType)
}

// BridgeTechnologyList
// BridgeTechnologyList list available bridging technologies and their statuses.
func (c *AMICore) BridgeTechnologyList(ctx context.Context) ([]AmiReply, error) {
	return BridgeTechnologyList(ctx, *c.socket)
}

// BridgeTechnologySuspend
// BridgeTechnologySuspend suspend a bridging technology.
func (c *AMICore) BridgeTechnologySuspend(ctx context.Context, bridgeTechnology string) (AmiReply, error) {
	return BridgeTechnologySuspend(ctx, *c.socket, bridgeTechnology)
}

// BridgeTechnologyUnsuspend
// BridgeTechnologyUnsuspend unsuspend a bridging technology.
func (c *AMICore) BridgeTechnologyUnsuspend(ctx context.Context, bridgeTechnology string) (AmiReply, error) {
	return BridgeTechnologyUnsuspend(ctx, *c.socket, bridgeTechnology)
}

// DBDel
// DBDel Delete DB entry.
func (c *AMICore) DBDel(ctx context.Context, family, key string) (AmiReply, error) {
	return DBDel(ctx, *c.socket, family, key)
}

// DBDelTree
// DBDelTree delete DB tree.
func (c *AMICore) DBDelTree(ctx context.Context, family, key string) (AmiReply, error) {
	return DBDelTree(ctx, *c.socket, family, key)
}

// DBPut
// DBPut puts DB entry.
func (c *AMICore) DBPut(ctx context.Context, family, key, value string) (AmiReply, error) {
	return DBPut(ctx, *c.socket, family, key, value)
}

// DBGet
// DBGet gets DB Entry.
func (c *AMICore) DBGet(ctx context.Context, family, key string) ([]AmiReply, error) {
	return DBGet(ctx, *c.socket, family, key)
}

// PRIDebugFileSet
// PRIDebugFileSet set the file used for PRI debug message output.
func (c *AMICore) PRIDebugFileSet(ctx context.Context, filename string) (AmiReply, error) {
	return PRIDebugFileSet(ctx, *c.socket, filename)
}

// PRIDebugFileUnset
// PRIDebugFileUnset disables file output for PRI debug messages.
func (c *AMICore) PRIDebugFileUnset(ctx context.Context) (AmiReply, error) {
	return PRIDebugFileUnset(ctx, *c.socket)
}

// PRIDebugSet
// PRIDebugSet set PRI debug levels for a span.
func (c *AMICore) PRIDebugSet(ctx context.Context, span, level string) (AmiReply, error) {
	return PRIDebugSet(ctx, *c.socket, span, level)
}

// PRIShowSpans
// PRIShowSpans show status of PRI spans.
func (c *AMICore) PRIShowSpans(ctx context.Context, span string) ([]AmiReply, error) {
	return PRIShowSpans(ctx, *c.socket, span)
}

// SKINNYDevices
// SKINNYDevices lists SKINNY devices (text format).
// Lists Skinny devices in text format with details on current status.
// Devicelist will follow as separate events,
// followed by a final event called DevicelistComplete.
func (c *AMICore) SKINNYDevices(ctx context.Context) ([]AmiReply, error) {
	return SKINNYDevices(ctx, *c.socket)
}

// SKINNYLines
// SKINNYLines lists SKINNY lines (text format).
// Lists Skinny lines in text format with details on current status.
// Linelist will follow as separate events,
// followed by a final event called LinelistComplete.
func (c *AMICore) SKINNYLines(ctx context.Context) ([]AmiReply, error) {
	return SKINNYLines(ctx, *c.socket)
}

// SKINNYShowDevice
// SKINNYShowDevice show SKINNY device (text format).
// Show one SKINNY device with details on current status.
func (c *AMICore) SKINNYShowDevice(ctx context.Context, device string) (AmiReply, error) {
	return SKINNYShowDevice(ctx, *c.socket, device)
}

// SKINNYShowline
// SKINNYShowline shows SKINNY line (text format).
// Show one SKINNY line with details on current status.
func (c *AMICore) SKINNYShowline(ctx context.Context, line string) (AmiReply, error) {
	return SKINNYShowline(ctx, *c.socket, line)
}

// MeetMeList
// MeetMeList lists all users in a particular MeetMe conference.
// Will follow as separate events, followed by a final event called MeetmeListComplete.
func (c *AMICore) MeetMeList(ctx context.Context, conference string) ([]AmiReply, error) {
	return MeetMeList(ctx, *c.socket, conference)
}

// MeetMeMute
// MeetMeMute mute a Meetme user.
func (c *AMICore) MeetMeMute(ctx context.Context, meetme, userNumber string) (AmiReply, error) {
	return MeetMeMute(ctx, *c.socket, meetme, userNumber)
}

// MeetMeUnMute
// MeetMeUnMute unmute a Meetme user.
func (c *AMICore) MeetMeUnMute(ctx context.Context, meetme, userNumber string) (AmiReply, error) {
	return MeetMeUnMute(ctx, *c.socket, meetme, userNumber)
}

// MeetMeListRooms
// MeetMeListRooms list active conferences.
func (c *AMICore) MeetMeListRooms(ctx context.Context) ([]AmiReply, error) {
	return MeetMeListRooms(ctx, *c.socket)
}

// Monitor
// Monitor monitors a channel.
// This action may be used to record the audio on a specified channel.
func (c *AMICore) Monitor(ctx context.Context, payload AMIPayloadMonitor) (AmiReply, error) {
	return Monitor(ctx, *c.socket, payload)
}

// MonitorWith
func (c *AMICore) MonitorWith(ctx context.Context, channel, file, format string, mix bool) (AmiReply, error) {
	return MonitorWith(ctx, *c.socket, channel, file, format, mix)
}

// ChangeMonitor
// ChangeMonitor changes monitoring filename of a channel.
// This action may be used to change the file started by a previous 'Monitor' action.
func (c *AMICore) ChangeMonitor(ctx context.Context, payload AMIPayloadMonitor) (AmiReply, error) {
	return ChangeMonitor(ctx, *c.socket, payload)
}

// ChangeMonitorWith
// ChangeMonitor changes monitoring filename of a channel.
// This action may be used to change the file started by a previous 'Monitor' action.
func (c *AMICore) ChangeMonitorWith(ctx context.Context, channel, file string) (AmiReply, error) {
	return ChangeMonitorWith(ctx, *c.socket, channel, file)
}

// MixMonitor
// MixMonitor record a call and mix the audio during the recording.
func (c *AMICore) MixMonitor(ctx context.Context, payload AMIPayloadMonitor) (AmiReply, error) {
	return MixMonitor(ctx, *c.socket, payload)
}

// MixMonitorWith
// MixMonitor record a call and mix the audio during the recording.
func (c *AMICore) MixMonitorWith(ctx context.Context, channel, file, options, command string) (AmiReply, error) {
	return MixMonitorWith(ctx, *c.socket, channel, file, options, command)
}

// MixMonitorMute
// MixMonitorMute Mute / unMute a Mixmonitor recording.
// This action may be used to mute a MixMonitor recording.
func (c *AMICore) MixMonitorMute(ctx context.Context, channel, direction string, state bool) (AmiReply, error) {
	return MixMonitorMute(ctx, *c.socket, channel, direction, state)
}

// PauseMonitor
// PauseMonitor pauses monitoring of a channel.
// This action may be used to temporarily stop the recording of a channel.
func (c *AMICore) PauseMonitor(ctx context.Context, channel string) (AmiReply, error) {
	return PauseMonitor(ctx, *c.socket, channel)
}

// UnpauseMonitor
// UnpauseMonitor unpause monitoring of a channel.
// This action may be used to re-enable recording of a channel after calling PauseMonitor.
func (c *AMICore) UnpauseMonitor(ctx context.Context, channel string) (AmiReply, error) {
	return UnpauseMonitor(ctx, *c.socket, channel)
}

// StopMonitor
// StopMonitor stops monitoring a channel.
// This action may be used to end a previously started 'Monitor' action.
func (c *AMICore) StopMonitor(ctx context.Context, channel string) (AmiReply, error) {
	return StopMonitor(ctx, *c.socket, channel)
}

// StopMixMonitor
// StopMixMonitor stop recording a call through MixMonitor, and free the recording's file handle.
func (c *AMICore) StopMixMonitor(ctx context.Context, channel, mixMonitorId string) (AmiReply, error) {
	return StopMixMonitor(ctx, *c.socket, channel, mixMonitorId)
}

// PJSIPNotify
// PJSIPNotify send NOTIFY to either an endpoint, an arbitrary URI, or inside a SIP dialog.
func (c *AMICore) PJSIPNotify(ctx context.Context, endpoint, uri, variable string) (AmiReply, error) {
	return PJSIPNotify(ctx, *c.socket, endpoint, uri, variable)
}

// PJSIPQualify
// PJSIPQualify qualify a chan_pjsip endpoint.
func (c *AMICore) PJSIPQualify(ctx context.Context, endpoint string) (AmiReply, error) {
	return PJSIPQualify(ctx, *c.socket, endpoint)
}

// PJSIPRegister
// PJSIPRegister register an outbound registration.
func (c *AMICore) PJSIPRegister(ctx context.Context, registration string) (AmiReply, error) {
	return PJSIPRegister(ctx, *c.socket, registration)
}

// PJSIPUnregister
// PJSIPUnregister unregister an outbound registration.
func (c *AMICore) PJSIPUnregister(ctx context.Context, registration string) (AmiReply, error) {
	return PJSIPUnregister(ctx, *c.socket, registration)
}

// PJSIPShowEndpoint
// PJSIPShowEndpoint detail listing of an endpoint and its objects.
func (c *AMICore) PJSIPShowEndpoint(ctx context.Context, endpoint string) ([]AmiReply, error) {
	return PJSIPShowEndpoint(ctx, *c.socket, endpoint)
}

// PJSIPShowEndpoints
// PJSIPShowEndpoints list pjsip endpoints.
func (c *AMICore) PJSIPShowEndpoints(ctx context.Context) ([]AmiReply, error) {
	return PJSIPShowEndpoints(ctx, *c.socket)
}

// PJSIPShowRegistrationInboundContactStatuses
func (c *AMICore) PJSIPShowRegistrationInboundContactStatuses(ctx context.Context) ([]AmiReply, error) {
	return PJSIPShowRegistrationInboundContactStatuses(ctx, *c.socket)
}

// PJSIPShowRegistrationsInbound
// PJSIPShowRegistrationsInbound lists PJSIP inbound registrations.
func (c *AMICore) PJSIPShowRegistrationsInbound(ctx context.Context) ([]AmiReply, error) {
	return PJSIPShowRegistrationsInbound(ctx, *c.socket)
}

// PJSIPShowRegistrationsOutbound
// PJSIPShowRegistrationsOutbound lists PJSIP outbound registrations.
func (c *AMICore) PJSIPShowRegistrationsOutbound(ctx context.Context) ([]AmiReply, error) {
	return PJSIPShowRegistrationsOutbound(ctx, *c.socket)
}

// PJSIPShowResourceLists
// PJSIPShowResourceLists displays settings for configured resource lists.
func (c *AMICore) PJSIPShowResourceLists(ctx context.Context) ([]AmiReply, error) {
	return PJSIPShowResourceLists(ctx, *c.socket)
}

// PJSIPShowSubscriptionsInbound
// PJSIPShowSubscriptionsInbound list of inbound subscriptions.
func (c *AMICore) PJSIPShowSubscriptionsInbound(ctx context.Context) ([]AmiReply, error) {
	return PJSIPShowSubscriptionsInbound(ctx, *c.socket)
}

// PJSIPShowSubscriptionsOutbound
// PJSIPShowSubscriptionsOutbound list of outbound subscriptions.
func (c *AMICore) PJSIPShowSubscriptionsOutbound(ctx context.Context) ([]AmiReply, error) {
	return PJSIPShowSubscriptionsOutbound(ctx, *c.socket)
}

// FAXSession
// FAXSession responds with a detailed description of a single FAX session.
func (c *AMICore) FAXSession(ctx context.Context, sessionNumber string) (AmiReply, error) {
	return FAXSession(ctx, *c.socket, sessionNumber)
}

// FAXSessions
// FAXSessions list active FAX sessions.
func (c *AMICore) FAXSessions(ctx context.Context) ([]AmiReply, error) {
	return FAXSessions(ctx, *c.socket)
}

// FAXStats
// FAXStats responds with fax statistics.
func (c *AMICore) FAXStats(ctx context.Context) (AmiReply, error) {
	return FAXStats(ctx, *c.socket)
}

// Atxfer
// Atxfer attended transfer.
func (c *AMICore) Atxfer(ctx context.Context, channel, extension, context string) (AmiReply, error) {
	return Atxfer(ctx, *c.socket, channel, extension, context)
}

// CancelAtxfer
// CancelAtxfer cancel an attended transfer.
func (c *AMICore) CancelAtxfer(ctx context.Context, channel string) (AmiReply, error) {
	return CancelAtxfer(ctx, *c.socket, channel)
}

// DAHDIDialOffhook
// DAHDIDialOffhook dials over DAHDI channel while offhook.
// Generate DTMF control frames to the bridged peer.
func (c *AMICore) DAHDIDialOffhook(ctx context.Context, channel, number string) (AmiReply, error) {
	return DAHDIDialOffhook(ctx, *c.socket, channel, number)
}

// DAHDIDNDoff
// DAHDIDNDoff toggles DAHDI channel Do Not Disturb status OFF.
func (c *AMICore) DAHDIDNDoff(ctx context.Context, channel string) (AmiReply, error) {
	return DAHDIDNDoff(ctx, *c.socket, channel)
}

// DAHDIDNDon
// DAHDIDNDon toggles DAHDI channel Do Not Disturb status ON.
func (c *AMICore) DAHDIDNDon(ctx context.Context, channel string) (AmiReply, error) {
	return DAHDIDNDon(ctx, *c.socket, channel)
}

// DAHDIHangup
// DAHDIHangup hangups DAHDI Channel.
func (c *AMICore) DAHDIHangup(ctx context.Context, channel string) (AmiReply, error) {
	return DAHDIHangup(ctx, *c.socket, channel)
}

// DAHDIRestart
// DAHDIRestart fully Restart DAHDI channels (terminates calls).
func (c *AMICore) DAHDIRestart(ctx context.Context) (AmiReply, error) {
	return DAHDIRestart(ctx, *c.socket)
}

// DAHDIShowChannels
// DAHDIShowChannels show status of DAHDI channels.
func (c *AMICore) DAHDIShowChannels(ctx context.Context, channel string) ([]AmiReply, error) {
	return DAHDIShowChannels(ctx, *c.socket, channel)
}

// DAHDITransfer
// DAHDITransfer transfers DAHDI Channel.
func (c *AMICore) DAHDITransfer(ctx context.Context, channel string) (AmiReply, error) {
	return DAHDITransfer(ctx, *c.socket, channel)
}

// ConfbridgeList
// ConfbridgeList lists all users in a particular ConfBridge conference.
func (c *AMICore) ConfbridgeList(ctx context.Context, conference string) ([]AmiReply, error) {
	return ConfbridgeList(ctx, *c.socket, conference)
}

// ConfbridgeListRooms
// ConfbridgeListRooms lists data about all active conferences.
func (c *AMICore) ConfbridgeListRooms(ctx context.Context) ([]AmiReply, error) {
	return ConfbridgeListRooms(ctx, *c.socket)
}

// ConfbridgeMute
// ConfbridgeMute mutes a specified user in a specified conference.
func (c *AMICore) ConfbridgeMute(ctx context.Context, conference string, channel string) (AmiReply, error) {
	return ConfbridgeMute(ctx, *c.socket, conference, channel)
}

// ConfbridgeUnmute
// ConfbridgeUnmute unmute a specified user in a specified conference.
func (c *AMICore) ConfbridgeUnmute(ctx context.Context, conference string, channel string) (AmiReply, error) {
	return ConfbridgeUnmute(ctx, *c.socket, conference, channel)
}

// ConfbridgeKick
// ConfbridgeKick removes a specified user from a specified conference.
func (c *AMICore) ConfbridgeKick(ctx context.Context, conference string, channel string) (AmiReply, error) {
	return ConfbridgeKick(ctx, *c.socket, conference, channel)
}

// ConfbridgeLock
// ConfbridgeLock locks a specified conference.
func (c *AMICore) ConfbridgeLock(ctx context.Context, conference string, channel string) (AmiReply, error) {
	return ConfbridgeLock(ctx, *c.socket, conference, channel)
}

// ConfbridgeUnlock
// ConfbridgeUnlock unlocks a specified conference.
func (c *AMICore) ConfbridgeUnlock(ctx context.Context, conference string, channel string) (AmiReply, error) {
	return ConfbridgeUnlock(ctx, *c.socket, config.AmiErrorLoginFailed, channel)
}

// ConfbridgeSetSingleVideoSrc
// ConfbridgeSetSingleVideoSrc sets a conference user as the single video source distributed to all other video-capable participants.
func (c *AMICore) ConfbridgeSetSingleVideoSrc(ctx context.Context, conference string, channel string) (AmiReply, error) {
	return ConfbridgeSetSingleVideoSrc(ctx, *c.socket, conference, channel)
}

// ConfbridgeStartRecord
// ConfbridgeStartRecord starts a recording in the context of given conference and creates a file with the name specified by recordFile
func (c *AMICore) ConfbridgeStartRecord(ctx context.Context, conference string, recordFile string) (AmiReply, error) {
	return ConfbridgeStartRecord(ctx, *c.socket, conference, recordFile)
}

// ConfbridgeStopRecord
// ConfbridgeStopRecord stops a recording pertaining to the given conference
func (c *AMICore) ConfbridgeStopRecord(ctx context.Context, conference string) (AmiReply, error) {
	return ConfbridgeStopRecord(ctx, *c.socket, conference)
}

// DialOut
func (c *AMICore) DialOut(ctx context.Context, d AMIDialCall) (AmiReply, bool, error) {
	return DialOut(ctx, *c.socket, d)
}

// DialIn
func (c *AMICore) DialIn(ctx context.Context, d AMIDialCall) (AmiReply, bool, error) {
	return DialIn(ctx, *c.socket, d)
}

// Chanspy
func (c *AMICore) Chanspy(ctx context.Context, ch AMIChanspy) (AmiReplies, error) {
	return Chanspy(ctx, *c.socket, ch)
}
