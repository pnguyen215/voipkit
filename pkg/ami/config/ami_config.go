package config

import (
	"time"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/fatal"
)

const (
	AmiActionKey      = "Action"
	AmiEventKey       = "Event"
	AmiResponseKey    = "Response"
	AmiActionIdKey    = "ActionID"
	AmiLoginKey       = "Login"
	AmiCallManagerKey = "Asterisk Call Manager"
)

const (
	AmiStatusSuccessKey = "success"
	AmiStatusFailedKey  = "failed"
)

const (
	AmiUsernameField = "Username"
	AmiSecretField   = "Secret"
	AmiPasswordField = "Password"
)

const (
	AmiPubSubKeyRef = "ami-key"
)

const (
	AmiNetworkTcpKey        = "tcp"
	AmiNetworkUdpKey        = "udp"
	AmiNetworkTcp4Key       = "tcp4"
	AmiNetworkTcp6Key       = "tcp6"
	AmiNetworkUdp4Key       = "udp4"
	AmiNetworkUdp6Key       = "udp6"
	AmiNetworkIpKey         = "ip"
	AmiNetworkIp4Key        = "ip4"
	AmiNetworkIp6Key        = "ip6"
	AmiNetworkUnixKey       = "unix"
	AmiNetworkUnixGramKey   = "unixgram"
	AmiNetworkUnixPacketKey = "unixpacket"
)

const (
	AmiProtocolHttpKey  = "http://"
	AmiProtocolHttpsKey = "https://"
)

const (
	AmiListenerEventDeviceStateChange        = "DeviceStateChange"
	AmiListenerEventNewChannel               = "Newchannel"
	AmiListenerEventNewState                 = "Newstate"
	AmiListenerEventSuccessfulAuth           = "SuccessfulAuth"
	AmiListenerEventNewExtension             = "Newexten"
	AmiListenerEventNewCallerId              = "NewCallerid"
	AmiListenerEventNewConnectedLine         = "NewConnectedLine"
	AmiListenerEventDialBegin                = "DialBegin"
	AmiListenerEventUserEvent                = "UserEvent"
	AmiListenerEventBridgeCreate             = "BridgeCreate"
	AmiListenerEventBridgeEnter              = "BridgeEnter"
	AmiListenerEventHangupRequest            = "HangupRequest"
	AmiListenerEventBridgeLeave              = "BridgeLeave"
	AmiListenerEventBridgeDestroy            = "BridgeDestroy"
	AmiListenerEventHangup                   = "Hangup"
	AmiListenerEventSoftHangupRequest        = "SoftHangupRequest"
	AmiListenerEventQueueParams              = "QueueParams"
	AmiListenerEventQueueMember              = "QueueMember"
	AmiListenerEventQueueStatusComplete      = "QueueStatusComplete"
	AmiListenerEventQueueMemberPause         = "QueueMemberPause"
	AmiListenerEventLocalBridge              = "LocalBridge"
	AmiListenerEventDialEnd                  = "DialEnd"
	AmiListenerEventConfBridgeJoin           = "ConfbridgeJoin"
	AmiListenerEventConfBridgeTalking        = "ConfbridgeTalking"
	AmiListenerEventConfBridgeKick           = "ConfbridgeKick"
	AmiListenerEventConfBridgeLeave          = "ConfbridgeLeave"
	AmiListenerEventMessageWaiting           = "MessageWaiting"
	AmiListenerEventCdr                      = "Cdr"
	AmiListenerEventExtensionStatus          = "ExtensionStatus"
	AmiListenerEventConnect                  = "Connect"
	AmiListenerEventDisconnect               = "Disconnect"
	AmiListenerEventDial                     = "Dial"
	AmiListenerEventBridge                   = "Bridge"
	AmiListenerEventRename                   = "Rename"
	AmiListenerEventVarSet                   = "VarSet"
	AmiListenerEventParkedCall               = "ParkedCall"
	AmiListenerEventParkedCallGiveUp         = "ParkedCallGiveUp"
	AmiListenerEventParkedCallTimeOut        = "ParkedCallTimeOut"
	AmiListenerEventUnParkedCall             = "UnparkedCall"
	AmiListenerEventJoin                     = "Join"
	AmiListenerEventLeave                    = "Leave"
	AmiListenerEventQueueMemberStatus        = "QueueMemberStatus"
	AmiListenerEventQueueMemberPenalty       = "QueueMemberPenalty"
	AmiListenerEventQueueMemberAdded         = "QueueMemberAdded"
	AmiListenerEventQueueMemberRemoved       = "QueueMemberRemoved"
	AmiListenerEventAbstractMeetMe           = "AbstractMeetMe"
	AmiListenerEventOriginateResponse        = "OriginateResponse"
	AmiListenerEventOriginate                = "Originate"
	AmiListenerEventAgents                   = "Agents"
	AmiListenerEventAgentCalled              = "AgentCalled"
	AmiListenerEventAgentConnect             = "AgentConnect"
	AmiListenerEventAgentComplete            = "AgentComplete"
	AmiListenerEventAgentCallbackLogin       = "AgentCallbackLogin"
	AmiListenerEventAgentCallbackLogoff      = "AgentCallbackLogoff"
	AmiListenerEventAgentLogin               = "AgentLogin"
	AmiListenerEventAgentLogoff              = "AgentLogoff"
	AmiListenerEventHoldedCall               = "HoldedCall"
	AmiListenerEventPeerStatus               = "PeerStatus"
	AmiListenerEventPeerlistComplete         = "PeerlistComplete"
	AmiListenerEventPeerEntry                = "PeerEntry"
	AmiListenerEventStatus                   = "Status"
	AmiListenerEventStatusComplete           = "StatusComplete"
	AmiListenerEventAgentRingNoAnswer        = "AgentRingNoAnswer"
	AmiListenerEventHold                     = "Hold"
	AmiListenerEventChannelUpdate            = "ChannelUpdate"
	AmiListenerEventDialState                = "DialState"
	AmiListenerEventInvalidPassword          = "InvalidPassword"
	AmiListenerEventMusicOnHold              = "MusicOnHold"
	AmiListenerEventPickup                   = "Pickup"
	AmiListenerEventPriEvent                 = "PriEvent"
	AmiListenerEventQueue                    = "Queue"
	AmiListenerEventAgentsComplete           = "AgentsComplete"
	AmiListenerEventUnHold                   = "Unhold"
	AmiListenerEventDbGetResponse            = "DbGetResponse"
	AmiListenerEventCommon                   = "Common"
	AmiListenerEventHangupHandlerPush        = "HangupHandlerPush"
	AmiListenerEventHangupHandlerRun         = "HangupHandlerRun"
	AmiListenerEventAgentDump                = "AgentDump"
	AmiListenerEventAGIExecEnd               = "AGIExecEnd"
	AmiListenerEventAGIExecStart             = "AGIExecStart"
	AmiListenerEventAlarm                    = "Alarm"
	AmiListenerEventAlarmClear               = "AlarmClear"
	AmiListenerEventAOCD                     = "AOC-D"
	AmiListenerEventAOCE                     = "AOC-E"
	AmiListenerEventAOCS                     = "AOC-S"
	AmiListenerEventAorDetail                = "AorDetail"
	AmiListenerEventAorList                  = "AorList"
	AmiListenerEventAorListComplete          = "AorListComplete"
	AmiListenerEventAsyncAGIEnd              = "AsyncAGIEnd"
	AmiListenerEventAsyncAGIExec             = "AsyncAGIExec"
	AmiListenerEventAsyncAGIStart            = "AsyncAGIStart"
	AmiListenerEventAttendedTransfer         = "AttendedTransfer"
	AmiListenerEventAuthDetail               = "AuthDetail"
	AmiListenerEventAuthList                 = "AuthList"
	AmiListenerEventAuthListComplete         = "AuthListComplete"
	AmiListenerEventAuthMethodNotAllowed     = "AuthMethodNotAllowed"
	AmiListenerEventBlindTransfer            = "BlindTransfer"
	AmiListenerEventBridgeInfoChannel        = "BridgeInfoChannel"
	AmiListenerEventBridgeInfoComplete       = "BridgeInfoComplete"
	AmiListenerEventBridgeMerge              = "BridgeMerge"
	AmiListenerEventBridgeVideoSourceUpdate  = "BridgeVideoSourceUpdate"
	AmiListenerEventCel                      = "CEL"
	AmiListenerEventChallengeResponseFailed  = "ChallengeResponseFailed"
	AmiListenerEventChallengeSent            = "ChallengeSent"
	AmiListenerEventChannelTalkingStart      = "ChannelTalkingStart"
	AmiListenerEventChannelTalkingStop       = "ChannelTalkingStop"
	AmiListenerEventChanSpyStart             = "ChanSpyStart"
	AmiListenerEventChanSpyStop              = "ChanSpyStop"
	AmiListenerEventConfbridgeEnd            = "ConfbridgeEnd"
	AmiListenerEventConfbridgeJoin           = "ConfbridgeJoin"
	AmiListenerEventConfbridgeLeave          = "ConfbridgeLeave"
	AmiListenerEventConfbridgeList           = "ConfbridgeList"
	AmiListenerEventConfbridgeMute           = "ConfbridgeMute"
	AmiListenerEventConfbridgeRecord         = "ConfbridgeRecord"
	AmiListenerEventConfbridgeStart          = "ConfbridgeStart"
	AmiListenerEventConfbridgeStopRecord     = "ConfbridgeStopRecord"
	AmiListenerEventConfbridgeTalking        = "ConfbridgeTalking"
	AmiListenerEventConfbridgeUnMute         = "ConfbridgeUnmute"
	AmiListenerEventContactList              = "ContactList"
	AmiListenerEventContactListComplete      = "ContactListComplete"
	AmiListenerEventContactStatus            = "ContactStatus"
	AmiListenerEventContactStatusDetail      = "ContactStatusDetail"
	AmiListenerEventCoreShowChannel          = "CoreShowChannel"
	AmiListenerEventCoreShowChannelsComplete = "CoreShowChannelsComplete"
	AmiListenerEventDAHDIChannel             = "DAHDIChannel"
	AmiListenerEventDeviceStateListComplete  = "DeviceStateListComplete"
	AmiListenerEventDNDState                 = "DNDState" // page 317
)

var (
	AmiNetworkKeys map[string]bool = map[string]bool{
		AmiNetworkTcpKey:        true,
		AmiNetworkUdpKey:        true,
		AmiNetworkTcp4Key:       true,
		AmiNetworkTcp6Key:       true,
		AmiNetworkUdp4Key:       true,
		AmiNetworkUdp6Key:       true,
		AmiNetworkIpKey:         true,
		AmiNetworkIp4Key:        true,
		AmiNetworkIp6Key:        true,
		AmiNetworkUnixKey:       true,
		AmiNetworkUnixGramKey:   true,
		AmiNetworkUnixPacketKey: true,
	}
	AmiProtocolKeys map[string]bool = map[string]bool{
		AmiProtocolHttpKey:  true,
		AmiProtocolHttpsKey: true,
	}
)

var (
	// ErrorAsteriskConnTimeout error on connection timeout
	ErrorAsteriskConnTimeout = fatal.AMIErrorNew("Asterisk Server connection timeout")

	// ErrorAsteriskInvalidPrompt invalid prompt received from AMI server
	ErrorAsteriskInvalidPrompt = fatal.AMIErrorNew("Asterisk Server invalid prompt command line")

	// ErrorAsteriskNetwork networking errors
	ErrorAsteriskNetwork = fatal.AMIErrorNew("Network error")

	// ErrorAsteriskLogin AMI server login failed
	ErrorAsteriskLogin = fatal.AMIErrorNew("Asterisk Server login failed")

	// Error EOF
	ErrorEOF = "EOF"

	// Error I/O
	ErrorIO          = "io: read/write on closed pipe"
	ErrorLoginFailed = "Failed login"
)

const (
	NetworkTimeoutAfterSeconds = time.Second * 3 // default is 3 seconds
)

const (
	AmiCliCommand string = "Command"
)

const (
	AmiCliErrorFieldRequired string = "%v is required"
)
