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
	AmiListenerEventDeviceStateChange          = "DeviceStateChange"
	AmiListenerEventNewChannel                 = "Newchannel"       // Raised when a new channel is created.
	AmiListenerEventNewState                   = "Newstate"         // Raised when a channel's state changes.
	AmiListenerEventSuccessfulAuth             = "SuccessfulAuth"   // Raised when a request successfully authenticates with a service
	AmiListenerEventNewExtension               = "Newexten"         // Raised when a channel enters a new context, extension, priority.
	AmiListenerEventNewCallerId                = "NewCallerid"      // Raised when a channel receives new Caller ID information.
	AmiListenerEventNewConnectedLine           = "NewConnectedLine" // Raised when a channel's connected line information is changed.
	AmiListenerEventDialBegin                  = "DialBegin"
	AmiListenerEventUserEvent                  = "UserEvent" // A user defined event raised from the dialplan.
	AmiListenerEventBridgeCreate               = "BridgeCreate"
	AmiListenerEventBridgeEnter                = "BridgeEnter"
	AmiListenerEventHangupRequest              = "HangupRequest" // Raised when a hangup is requested
	AmiListenerEventBridgeLeave                = "BridgeLeave"
	AmiListenerEventBridgeDestroy              = "BridgeDestroy"
	AmiListenerEventHangup                     = "Hangup"            // Raised when a channel is hung up
	AmiListenerEventSoftHangupRequest          = "SoftHangupRequest" // Raised when a soft hangup is requested with a specific cause code.
	AmiListenerEventQueueParams                = "QueueParams"
	AmiListenerEventQueueMember                = "QueueMember"
	AmiListenerEventQueueStatusComplete        = "QueueStatusComplete"
	AmiListenerEventQueueMemberPause           = "QueueMemberPause" // Raised when a member is paused/unpaused in the queue.
	AmiListenerEventLocalBridge                = "LocalBridge"      // Raised when two halves of a Local Channel form a bridge.
	AmiListenerEventDialEnd                    = "DialEnd"
	AmiListenerEventConfBridgeJoin             = "ConfbridgeJoin"
	AmiListenerEventConfBridgeTalking          = "ConfbridgeTalking"
	AmiListenerEventConfBridgeKick             = "ConfbridgeKick"
	AmiListenerEventConfBridgeLeave            = "ConfbridgeLeave"
	AmiListenerEventMessageWaiting             = "MessageWaiting" // Raised when the state of messages in a voicemail mailbox has changed or when a channel has finished interacting with a mailbox
	AmiListenerEventCdr                        = "Cdr"
	AmiListenerEventExtensionStatus            = "ExtensionStatus"
	AmiListenerEventConnect                    = "Connect"
	AmiListenerEventDisconnect                 = "Disconnect"
	AmiListenerEventDial                       = "Dial"
	AmiListenerEventBridge                     = "Bridge"
	AmiListenerEventRename                     = "Rename"            // Raised when the name of a channel is changed.
	AmiListenerEventVarSet                     = "VarSet"            // Raised when a variable local to the gosub stack frame is set due to a subroutine call.
	AmiListenerEventParkedCall                 = "ParkedCall"        // Raised when a channel is parked.
	AmiListenerEventParkedCallGiveUp           = "ParkedCallGiveUp"  // Raised when a channel leaves a parking lot because it hung up without being answered
	AmiListenerEventParkedCallTimeOut          = "ParkedCallTimeOut" // Raised when a channel leaves a parking lot due to reaching the time limit of being parked.
	AmiListenerEventUnParkedCall               = "UnparkedCall"      // Raised when a channel leaves a parking lot because it was retrieved from the parking lot and reconnected.
	AmiListenerEventJoin                       = "Join"
	AmiListenerEventLeave                      = "Leave"
	AmiListenerEventQueueMemberStatus          = "QueueMemberStatus"  // Raised when a Queue member's status has changed.
	AmiListenerEventQueueMemberPenalty         = "QueueMemberPenalty" // Raised when a member's penalty is changed.
	AmiListenerEventQueueMemberAdded           = "QueueMemberAdded"   // Raised when a member is added to the queue.
	AmiListenerEventQueueMemberRemoved         = "QueueMemberRemoved" // Raised when a member is removed from the queue.
	AmiListenerEventAbstractMeetMe             = "AbstractMeetMe"
	AmiListenerEventOriginateResponse          = "OriginateResponse" // Raised in response to an Originate command
	AmiListenerEventOriginate                  = "Originate"
	AmiListenerEventAgents                     = "Agents"        // Response event in a series to the Agents AMI action containing information about a defined agent.
	AmiListenerEventAgentCalled                = "AgentCalled"   // Raised when an queue member is notified of a caller in the queue
	AmiListenerEventAgentConnect               = "AgentConnect"  // Raised when a queue member answers and is bridged to a caller in the queue.
	AmiListenerEventAgentComplete              = "AgentComplete" // Raised when a queue member has finished servicing a caller in the queue.
	AmiListenerEventAgentCallbackLogin         = "AgentCallbackLogin"
	AmiListenerEventAgentCallbackLogoff        = "AgentCallbackLogoff"
	AmiListenerEventAgentLogin                 = "AgentLogin"  // Raised when an Agent has logged in.
	AmiListenerEventAgentLogoff                = "AgentLogoff" // Raised when an Agent has logged off.
	AmiListenerEventHoldedCall                 = "HoldedCall"
	AmiListenerEventPeerStatus                 = "PeerStatus" // Raised when the state of a peer changes.
	AmiListenerEventPeerlistComplete           = "PeerlistComplete"
	AmiListenerEventPeerEntry                  = "PeerEntry"
	AmiListenerEventStatus                     = "Status"            // Raised in response to a Status command.
	AmiListenerEventStatusComplete             = "StatusComplete"    // Raised in response to a Status command.
	AmiListenerEventAgentRingNoAnswer          = "AgentRingNoAnswer" // Raised when a queue member is notified of a caller in the queue and fails to answer.
	AmiListenerEventHold                       = "Hold"              // Raised when a channel goes on hold.
	AmiListenerEventChannelUpdate              = "ChannelUpdate"
	AmiListenerEventDialState                  = "DialState"
	AmiListenerEventInvalidPassword            = "InvalidPassword" // Raised when a request provides an invalid password during an authentication attempt
	AmiListenerEventMusicOnHold                = "MusicOnHold"
	AmiListenerEventPickup                     = "Pickup" // Raised when a call pickup occurs
	AmiListenerEventPriEvent                   = "PriEvent"
	AmiListenerEventQueue                      = "Queue"
	AmiListenerEventAgentsComplete             = "AgentsComplete" // Final response event in a series of events to the Agents AMI action.
	AmiListenerEventUnHold                     = "Unhold"         // Raised when a channel goes off hold.
	AmiListenerEventDbGetResponse              = "DbGetResponse"
	AmiListenerEventCommon                     = "Common"
	AmiListenerEventHangupHandlerPush          = "HangupHandlerPush" // Raised when a hangup handler is added to the handler stack by the CHANNEL() function.
	AmiListenerEventHangupHandlerRun           = "HangupHandlerRun"  // Raised when a hangup handler is about to be called.
	AmiListenerEventAgentDump                  = "AgentDump"         // Raised when a queue member hangs up on a caller in the queue
	AmiListenerEventAGIExecEnd                 = "AGIExecEnd"        // Raised when a received AGI command completes processing.
	AmiListenerEventAGIExecStart               = "AGIExecStart"      // Raised when a received AGI command starts processing.
	AmiListenerEventAlarm                      = "Alarm"             // Raised when an alarm is set on a DAHDI channel.
	AmiListenerEventAlarmClear                 = "AlarmClear"        // Raised when an alarm is cleared on a DAHDI channel.
	AmiListenerEventAOCD                       = "AOC-D"             // Raised when an Advice of Charge message is sent during a call.
	AmiListenerEventAOCE                       = "AOC-E"             // Raised when an Advice of Charge message is sent at the end of a call.
	AmiListenerEventAOCS                       = "AOC-S"             // Raised when an Advice of Charge message is sent at the beginning of a call.
	AmiListenerEventAorDetail                  = "AorDetail"         // Provide details about an Address of Record (AoR) section.
	AmiListenerEventAorList                    = "AorList"           // Provide details about an Address of Record (AoR) section.
	AmiListenerEventAorListComplete            = "AorListComplete"   // Provide final information about an aor list
	AmiListenerEventAsyncAGIEnd                = "AsyncAGIEnd"       // Raised when a channel stops AsyncAGI command processing.
	AmiListenerEventAsyncAGIExec               = "AsyncAGIExec"      // Raised when AsyncAGI completes an AGI command.
	AmiListenerEventAsyncAGIStart              = "AsyncAGIStart"     // Raised when a channel starts AsyncAGI command processing.
	AmiListenerEventAttendedTransfer           = "AttendedTransfer"  // Raised when an attended transfer is complete
	AmiListenerEventAuthDetail                 = "AuthDetail"
	AmiListenerEventAuthList                   = "AuthList"
	AmiListenerEventAuthListComplete           = "AuthListComplete"
	AmiListenerEventAuthMethodNotAllowed       = "AuthMethodNotAllowed"
	AmiListenerEventBlindTransfer              = "BlindTransfer"
	AmiListenerEventBridgeInfoChannel          = "BridgeInfoChannel"
	AmiListenerEventBridgeInfoComplete         = "BridgeInfoComplete"
	AmiListenerEventBridgeMerge                = "BridgeMerge"
	AmiListenerEventBridgeVideoSourceUpdate    = "BridgeVideoSourceUpdate"
	AmiListenerEventCel                        = "CEL"
	AmiListenerEventChallengeResponseFailed    = "ChallengeResponseFailed"
	AmiListenerEventChallengeSent              = "ChallengeSent"
	AmiListenerEventChannelTalkingStart        = "ChannelTalkingStart"
	AmiListenerEventChannelTalkingStop         = "ChannelTalkingStop"
	AmiListenerEventChanSpyStart               = "ChanSpyStart"
	AmiListenerEventChanSpyStop                = "ChanSpyStop"
	AmiListenerEventConfbridgeEnd              = "ConfbridgeEnd"
	AmiListenerEventConfbridgeJoin             = "ConfbridgeJoin"
	AmiListenerEventConfbridgeLeave            = "ConfbridgeLeave"
	AmiListenerEventConfbridgeList             = "ConfbridgeList"
	AmiListenerEventConfbridgeMute             = "ConfbridgeMute"
	AmiListenerEventConfbridgeRecord           = "ConfbridgeRecord"
	AmiListenerEventConfbridgeStart            = "ConfbridgeStart"
	AmiListenerEventConfbridgeStopRecord       = "ConfbridgeStopRecord"
	AmiListenerEventConfbridgeTalking          = "ConfbridgeTalking"
	AmiListenerEventConfbridgeUnMute           = "ConfbridgeUnmute"
	AmiListenerEventContactList                = "ContactList"
	AmiListenerEventContactListComplete        = "ContactListComplete"
	AmiListenerEventContactStatus              = "ContactStatus"
	AmiListenerEventContactStatusDetail        = "ContactStatusDetail"
	AmiListenerEventCoreShowChannel            = "CoreShowChannel"
	AmiListenerEventCoreShowChannelsComplete   = "CoreShowChannelsComplete"
	AmiListenerEventDAHDIChannel               = "DAHDIChannel"
	AmiListenerEventDeviceStateListComplete    = "DeviceStateListComplete"
	AmiListenerEventDNDState                   = "DNDState" // page 317
	AmiListenerEventDTMFBegin                  = "DTMFBegin"
	AmiListenerEventDTMFEnd                    = "DTMFEnd"
	AmiListenerEventEndpointDetail             = "EndpointDetail"
	AmiListenerEventEndpointDetailComplete     = "EndpointDetailComplete"
	AmiListenerEventEndpointList               = "EndpointList"
	AmiListenerEventEndpointListComplete       = "EndpointListComplete"
	AmiListenerEventExtensionStateListComplete = "ExtensionStateListComplete"
	AmiListenerEventFailedACL                  = "FailedACL"
	AmiListenerEventFAXSession                 = "FAXSession"
	AmiListenerEventFAXSessionsComplete        = "FAXSessionsComplete"
	AmiListenerEventFAXSessionsEntry           = "FAXSessionsEntry"
	AmiListenerEventFAXStats                   = "FAXStats"
	AmiListenerEventFAXStatus                  = "FAXStatus"
	AmiListenerEventFullyBooted                = "FullyBooted"               // Raised when all Asterisk initialization procedures have finished
	AmiListenerEventHangupHandlerPop           = "HangupHandlerPop"          // Raised when a hangup handler is removed from the handler stack by the CHANNEL() function.
	AmiListenerEventIdentifyDetail             = "IdentifyDetail"            // Provide details about an identify section.
	AmiListenerEventInvalidAccountID           = "InvalidAccountID"          // Raised when a request fails an authentication check due to an invalid account ID.
	AmiListenerEventInvalidTransport           = "InvalidTransport"          // Raised when a request attempts to use a transport not allowed by the Asterisk service.
	AmiListenerEventLoad                       = "Load"                      // Raised when a module has been loaded in Asterisk.
	AmiListenerEventLoadAverageLimit           = "LoadAverageLimit"          // Raised when a request fails because a configured load average limit has been reached.
	AmiListenerEventLocalOptimizationBegin     = "LocalOptimizationBegin"    // Raised when two halves of a Local Channel begin to optimize themselves out of the media path
	AmiListenerEventLocalOptimizationEnd       = "LocalOptimizationEnd"      // Raised when two halves of a Local Channel have finished optimizing themselves out of the media path.
	AmiListenerEventLogChannel                 = "LogChannel"                // Raised when a logging channel is re-enabled after a reload operation.
	AmiListenerEventMCID                       = "MCID"                      // Published when a malicious call ID request arrives
	AmiListenerEventMeetMeEnd                  = "MeetmeEnd"                 // Raised when a MeetMe conference ends
	AmiListenerEventMeetMeJoin                 = "MeetmeJoin"                // Raised when a user joins a MeetMe conference.
	AmiListenerEventMeetMeLeave                = "MeetmeLeave"               // Raised when a user leaves a MeetMe conference.
	AmiListenerEventMeetMeMute                 = "MeetmeMute"                // Raised when a MeetMe user is muted or unmuted.
	AmiListenerEventMeetMeTalking              = "MeetmeTalking"             // Raised when a MeetMe user begins or ends talking.
	AmiListenerEventMeetMeTalkRequest          = "MeetmeTalkRequest"         // Raised when a MeetMe user has started talking.
	AmiListenerEventMemoryLimit                = "MemoryLimit"               // Raised when a request fails due to an internal memory allocation failure.
	AmiListenerEventMiniVoiceMail              = "MiniVoiceMail"             // Raised when a notification is sent out by a MiniVoiceMail application
	AmiListenerEventMonitorStart               = "MonitorStart"              // Raised when monitoring has started on a channel.
	AmiListenerEventMonitorStop                = "MonitorStop"               // Raised when monitoring has stopped on a channel
	AmiListenerEventMusicOnHoldStart           = "MusicOnHoldStart"          // Raised when music on hold has started on a channel.
	AmiListenerEventMusicOnHoldStop            = "MusicOnHoldStop"           // Raised when music on hold has stopped on a channel.
	AmiListenerEventMWIGet                     = "MWIGet"                    // Raised in response to a MWIGet command.
	AmiListenerEventMWIGetComplete             = "MWIGetComplete"            // Raised in response to a MWIGet command.
	AmiListenerEventNewAccountCode             = "NewAccountCode"            // Raised when a Channel's AccountCode is changed.
	AmiListenerEventParkedCallSwap             = "ParkedCallSwap"            // Raised when a channel takes the place of a previously parked channel
	AmiListenerEventPresenceStateChange        = "PresenceStateChange"       // Raised when a presence state changes
	AmiListenerEventPresenceStateListComplete  = "PresenceStateListComplete" // Indicates the end of the list the current known extension states.
	AmiListenerEventPresenceStatus             = "PresenceStatus"            // Raised when a hint changes due to a presence state change.
	AmiListenerEventQueueCallerAbandon         = "QueueCallerAbandon"        // Raised when a caller abandons the queue.
	AmiListenerEventQueueCallerJoin            = "QueueCallerJoin"           // Raised when a caller joins a Queue.
	AmiListenerEventQueueCallerLeave           = "QueueCallerLeave"          // Raised when a caller leaves a Queue.
	AmiListenerEventQueueMemberRinginuse       = "QueueMemberRinginuse"      // Raised when a member's ringinuse setting is changed.
	AmiListenerEventReceiveFAX                 = "ReceiveFAX"                // Raised when a receive fax operation has completed.
	AmiListenerEventRegistry                   = "Registry"                  // Raised when an outbound registration completes.
	AmiListenerEventReload                     = "Reload"                    // Raised when a module has been reloaded in Asterisk.
	AmiListenerEventRequestBadFormat           = "RequestBadFormat"          // Raised when a request is received with bad formatting.
	AmiListenerEventRequestNotAllowed          = "RequestNotAllowed"         // Raised when a request is not allowed by the service.
	AmiListenerEventRequestNotSupported        = "RequestNotSupported"       // Raised when a request fails due to some aspect of the requested item not being supported by the service.
	AmiListenerEventRTCPReceived               = "RTCPReceived"              // Raised when an RTCP packet is received.
	AmiListenerEventRTCPSent                   = "RTCPSent"                  // Raised when an RTCP packet is sent.
	AmiListenerEventSendFAX                    = "SendFAX"                   // Raised when a send fax operation has completed.
	AmiListenerEventSessionLimit               = "SessionLimit"              // Raised when a request fails due to exceeding the number of allowed concurrent sessions for that service.
	AmiListenerEventSessionTimeout             = "SessionTimeout"            // Raised when a SIP session times out.
	AmiListenerEventShutdown                   = "Shutdown"                  // Raised when Asterisk is shutdown or restarted.
	AmiListenerEventSIPQualifyPeerDone         = "SIPQualifyPeerDone"        // Raised when SIPQualifyPeer has finished qualifying the specified peer
	AmiListenerEventSpanAlarm                  = "SpanAlarm"                 // Raised when an alarm is set on a DAHDI span.
	AmiListenerEventSpanAlarmClear             = "SpanAlarmClear"            // Raised when an alarm is cleared on a DAHDI span.
	AmiListenerEventTransportDetail            = "TransportDetail"           // Provide details about an authentication section.
	AmiListenerEventUnexpectedAddress          = "UnexpectedAddress"         // Raised when a request has a different source address then what is expected for a session already in progress with a service
	AmiListenerEventUnload                     = "Unload"                    // Raised when a module has been unloaded in Asterisk.
)

const (
	AmiClassCommand  = "COMMAND"
	AmiClassSecurity = "SECURITY"
	AmiClassCall     = "CALL"
	AmiClassSystem   = "SYSTEM"
	AmiClassUser     = "USER"
	AmiClassDialPlan = "DIALPLAN"
	AmiClassAgent    = "AGENT"
	AmiClassAgi      = "AGI"
	AmiClassAoc      = "AOC"
)

// page 257
var (
	AmiClassCommands map[string][]string = map[string][]string{
		AmiClassCommand: {
			AmiListenerEventAorDetail,
			AmiListenerEventAorList,
			AmiListenerEventAorListComplete,
		},
	}
	AmiClassSecurities map[string][]string = map[string][]string{
		AmiClassSecurity: {
			AmiListenerEventUnexpectedAddress,
		},
	}
	AmiClassCalls map[string][]string = map[string][]string{
		AmiClassCall: {
			AmiListenerEventUnHold,
			AmiListenerEventUnParkedCall,
			AmiListenerEventAttendedTransfer,
		},
	}
	AmiClassSystems map[string][]string = map[string][]string{
		AmiClassSystem: {
			AmiListenerEventUnload,
			AmiListenerEventAlarm,
			AmiListenerEventAlarmClear,
		},
	}
	AmiClassUsers map[string][]string = map[string][]string{
		AmiClassUser: {
			AmiListenerEventUserEvent,
		},
	}
	AmiClassDialPlans map[string][]string = map[string][]string{
		AmiClassDialPlan: {
			AmiListenerEventVarSet,
		},
	}
	AmiClassAgents map[string][]string = map[string][]string{
		AmiClassAgent: {
			AmiListenerEventAgentCalled,
			AmiListenerEventAgentComplete,
			AmiListenerEventAgentConnect,
			AmiListenerEventAgentDump,
			AmiListenerEventAgentLogin,
			AmiListenerEventAgentLogoff,
			AmiListenerEventAgentRingNoAnswer,
			AmiListenerEventAgents,
			AmiListenerEventAgentsComplete,
		},
	}
	AmiClassAgis map[string][]string = map[string][]string{
		AmiClassAgi: {
			AmiListenerEventAGIExecEnd,
			AmiListenerEventAGIExecStart,
			AmiListenerEventAsyncAGIEnd,
			AmiListenerEventAsyncAGIExec,
			AmiListenerEventAsyncAGIStart,
		},
	}
	AmiClassAocs map[string][]string = map[string][]string{
		AmiClassAoc: {
			AmiListenerEventAOCD,
			AmiListenerEventAOCE,
			AmiListenerEventAOCS,
		},
	}
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
	AmiErrorFieldRequired   string = "%v is required"
	AmiErrorInvalidProtocol string = "Invalid protocol"
	AmiErrorProtocolMessage string = "Protocol must have values: %v"
)

const (
	AmiDigitExtensionRegexDefault    string = "^SIP/\\d{4}"
	AmiDigitExtensionRegexWithDigits string = "^SIP/\\d{%v}"
)

const (
	AmiSIPChannelProtocol  string = "SIP"
	AmiIAXChannelProtocol  string = "IAX"
	AmiZapChannelProtocol  string = "Zap"
	AmiH323ChannelProtocol string = "H323"
)

var (
	AmiChannelProtocols map[string]bool = map[string]bool{
		AmiSIPChannelProtocol:  true,
		AmiIAXChannelProtocol:  true,
		AmiZapChannelProtocol:  true,
		AmiH323ChannelProtocol: true,
	}
)
