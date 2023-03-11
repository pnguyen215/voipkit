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
	AmiListenerEventDeviceStateChange          = "DeviceStateChange" // Raised when a device state changes. This differs from the ExtensionStatus event because this event is raised for all device state changes, not only for changes that affect dialplan hints.
	AmiListenerEventNewChannel                 = "Newchannel"        // Raised when a new channel is created.
	AmiListenerEventNewState                   = "Newstate"          // Raised when a channel's state changes.
	AmiListenerEventSuccessfulAuth             = "SuccessfulAuth"    // Raised when a request successfully authenticates with a service
	AmiListenerEventNewExtension               = "NewExten"          // Raised when a channel enters a new context, extension, priority.
	AmiListenerEventNewCallerId                = "NewCallerid"       // Raised when a channel receives new Caller ID information.
	AmiListenerEventNewConnectedLine           = "NewConnectedLine"  // Raised when a channel's connected line information is changed.
	AmiListenerEventDialBegin                  = "DialBegin"         // Raised when a dial action has started
	AmiListenerEventUserEvent                  = "UserEvent"         // A user defined event raised from the dialplan.
	AmiListenerEventBridgeCreate               = "BridgeCreate"      // Raised when a bridge is created
	AmiListenerEventBridgeEnter                = "BridgeEnter"       // Raised when a channel enters a bridge.
	AmiListenerEventHangupRequest              = "HangupRequest"     // Raised when a hangup is requested
	AmiListenerEventBridgeLeave                = "BridgeLeave"       // Raised when a channel leaves a bridge.
	AmiListenerEventBridgeDestroy              = "BridgeDestroy"     // Raised when a bridge is destroyed.
	AmiListenerEventHangup                     = "Hangup"            // Raised when a channel is hung up
	AmiListenerEventSoftHangupRequest          = "SoftHangupRequest" // Raised when a soft hangup is requested with a specific cause code.
	AmiListenerEventQueueParams                = "QueueParams"
	AmiListenerEventQueueMember                = "QueueMember"
	AmiListenerEventQueueStatusComplete        = "QueueStatusComplete"
	AmiListenerEventQueueMemberPause           = "QueueMemberPause"  // Raised when a member is paused/unpaused in the queue.
	AmiListenerEventLocalBridge                = "LocalBridge"       // Raised when two halves of a Local Channel form a bridge.
	AmiListenerEventDialEnd                    = "DialEnd"           // Raised when a dial action has completed.
	AmiListenerEventConfBridgeJoin             = "ConfbridgeJoin"    // Raised when a channel joins a Confbridge conference.
	AmiListenerEventConfBridgeTalking          = "ConfbridgeTalking" // Raised when a confbridge participant begins or ends talking.
	AmiListenerEventConfBridgeKick             = "ConfbridgeKick"
	AmiListenerEventConfBridgeLeave            = "ConfbridgeLeave" // Raised when a channel leaves a Confbridge conference.
	AmiListenerEventMessageWaiting             = "MessageWaiting"  // Raised when the state of messages in a voicemail mailbox has changed or when a channel has finished interacting with a mailbox
	AmiListenerEventCdr                        = "Cdr"             // Raised when a CDR is generated. The Cdr event is only raised when the cdr_manager backend is loaded and registered with the CDR engine.
	AmiListenerEventExtensionStatus            = "ExtensionStatus" // Raised when a hint changes due to a device state change.
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
	AmiListenerEventDialState                  = "DialState"       // Raised when dial status has changed.
	AmiListenerEventInvalidPassword            = "InvalidPassword" // Raised when a request provides an invalid password during an authentication attempt
	AmiListenerEventMusicOnHold                = "MusicOnHold"
	AmiListenerEventPickup                     = "Pickup" // Raised when a call pickup occurs
	AmiListenerEventPriEvent                   = "PriEvent"
	AmiListenerEventQueue                      = "Queue"
	AmiListenerEventAgentsComplete             = "AgentsComplete" // Final response event in a series of events to the Agents AMI action.
	AmiListenerEventUnHold                     = "Unhold"         // Raised when a channel goes off hold.
	AmiListenerEventDbGetResponse              = "DbGetResponse"
	AmiListenerEventCommon                     = "Common"
	AmiListenerEventHangupHandlerPush          = "HangupHandlerPush"          // Raised when a hangup handler is added to the handler stack by the CHANNEL() function.
	AmiListenerEventHangupHandlerRun           = "HangupHandlerRun"           // Raised when a hangup handler is about to be called.
	AmiListenerEventAgentDump                  = "AgentDump"                  // Raised when a queue member hangs up on a caller in the queue
	AmiListenerEventAGIExecEnd                 = "AGIExecEnd"                 // Raised when a received AGI command completes processing.
	AmiListenerEventAGIExecStart               = "AGIExecStart"               // Raised when a received AGI command starts processing.
	AmiListenerEventAlarm                      = "Alarm"                      // Raised when an alarm is set on a DAHDI channel.
	AmiListenerEventAlarmClear                 = "AlarmClear"                 // Raised when an alarm is cleared on a DAHDI channel.
	AmiListenerEventAOCD                       = "AOC-D"                      // Raised when an Advice of Charge message is sent during a call.
	AmiListenerEventAOCE                       = "AOC-E"                      // Raised when an Advice of Charge message is sent at the end of a call.
	AmiListenerEventAOCS                       = "AOC-S"                      // Raised when an Advice of Charge message is sent at the beginning of a call.
	AmiListenerEventAorDetail                  = "AorDetail"                  // Provide details about an Address of Record (AoR) section.
	AmiListenerEventAorList                    = "AorList"                    // Provide details about an Address of Record (AoR) section.
	AmiListenerEventAorListComplete            = "AorListComplete"            // Provide final information about an aor list
	AmiListenerEventAsyncAGIEnd                = "AsyncAGIEnd"                // Raised when a channel stops AsyncAGI command processing.
	AmiListenerEventAsyncAGIExec               = "AsyncAGIExec"               // Raised when AsyncAGI completes an AGI command.
	AmiListenerEventAsyncAGIStart              = "AsyncAGIStart"              // Raised when a channel starts AsyncAGI command processing.
	AmiListenerEventAttendedTransfer           = "AttendedTransfer"           // Raised when an attended transfer is complete
	AmiListenerEventAuthDetail                 = "AuthDetail"                 // Provide details about an authentication section.
	AmiListenerEventAuthList                   = "AuthList"                   // Provide details about an Address of Record (Auth) section.
	AmiListenerEventAuthListComplete           = "AuthListComplete"           // Provide final information about an auth list.
	AmiListenerEventAuthMethodNotAllowed       = "AuthMethodNotAllowed"       // Raised when a request used an authentication method not allowed by the service
	AmiListenerEventBlindTransfer              = "BlindTransfer"              // Raised when a blind transfer is complete.
	AmiListenerEventBridgeInfoChannel          = "BridgeInfoChannel"          // Information about a channel in a bridge.
	AmiListenerEventBridgeInfoComplete         = "BridgeInfoComplete"         // Information about a bridge.
	AmiListenerEventBridgeMerge                = "BridgeMerge"                // Raised when two bridges are merged
	AmiListenerEventBridgeVideoSourceUpdate    = "BridgeVideoSourceUpdate"    // Raised when the channel that is the source of video in a bridge changes
	AmiListenerEventCel                        = "CEL"                        // Raised when a Channel Event Log is generated for a channel.
	AmiListenerEventChallengeResponseFailed    = "ChallengeResponseFailed"    // Raised when a request's attempt to authenticate has been challenged, and the request failed the authentication challenge
	AmiListenerEventChallengeSent              = "ChallengeSent"              // Raised when an Asterisk service sends an authentication challenge to a request.
	AmiListenerEventChannelTalkingStart        = "ChannelTalkingStart"        // Raised when talking is detected on a channel.
	AmiListenerEventChannelTalkingStop         = "ChannelTalkingStop"         // Raised when talking is no longer detected on a channel.
	AmiListenerEventChanSpyStart               = "ChanSpyStart"               // Raised when one channel begins spying on another channel.
	AmiListenerEventChanSpyStop                = "ChanSpyStop"                // Raised when a channel has stopped spying.
	AmiListenerEventConfbridgeEnd              = "ConfbridgeEnd"              // Raised when a conference ends.
	AmiListenerEventConfbridgeJoin             = "ConfbridgeJoin"             // Raised when a channel joins a Confbridge conference.
	AmiListenerEventConfbridgeLeave            = "ConfbridgeLeave"            // Raised when a channel leaves a Confbridge conference.
	AmiListenerEventConfbridgeList             = "ConfbridgeList"             // Raised as part of the ConfbridgeList action response list.
	AmiListenerEventConfbridgeMute             = "ConfbridgeMute"             // Raised when a Confbridge participant mutes.
	AmiListenerEventConfbridgeRecord           = "ConfbridgeRecord"           // Raised when a conference starts recording.
	AmiListenerEventConfbridgeStart            = "ConfbridgeStart"            // Raised when a conference starts.
	AmiListenerEventConfbridgeStopRecord       = "ConfbridgeStopRecord"       // Raised when a conference that was recording stops recording.
	AmiListenerEventConfbridgeTalking          = "ConfbridgeTalking"          // Raised when a confbridge participant begins or ends talking.
	AmiListenerEventConfbridgeUnMute           = "ConfbridgeUnmute"           // Raised when a confbridge participant unmutes.
	AmiListenerEventContactList                = "ContactList"                // Provide details about a contact section.
	AmiListenerEventContactListComplete        = "ContactListComplete"        // Provide final information about a contact list.
	AmiListenerEventContactStatus              = "ContactStatus"              // Raised when the state of a contact changes
	AmiListenerEventContactStatusDetail        = "ContactStatusDetail"        // Provide details about a contact's status.
	AmiListenerEventCoreShowChannel            = "CoreShowChannel"            // Raised in response to a CoreShowChannels command.
	AmiListenerEventCoreShowChannelsComplete   = "CoreShowChannelsComplete"   // Raised at the end of the CoreShowChannel list produced by the CoreShowChannels command.
	AmiListenerEventDAHDIChannel               = "DAHDIChannel"               // Raised when a DAHDI channel is created or an underlying technology is associated with a DAHDI channel
	AmiListenerEventDeviceStateListComplete    = "DeviceStateListComplete"    // Indicates the end of the list the current known extension states.
	AmiListenerEventDNDState                   = "DNDState"                   // page 317, Raised when the Do Not Disturb state is changed on a DAHDI channel.
	AmiListenerEventDTMFBegin                  = "DTMFBegin"                  // Raised when a DTMF digit has started on a channel
	AmiListenerEventDTMFEnd                    = "DTMFEnd"                    // Raised when a DTMF digit has ended on a channel.
	AmiListenerEventEndpointDetail             = "EndpointDetail"             // Provide details about an endpoint section.
	AmiListenerEventEndpointDetailComplete     = "EndpointDetailComplete"     // Provide final information about endpoint details
	AmiListenerEventEndpointList               = "EndpointList"               // Provide details about a contact's status.
	AmiListenerEventEndpointListComplete       = "EndpointListComplete"       // Provide final information about an endpoint list
	AmiListenerEventExtensionStateListComplete = "ExtensionStateListComplete" // Indicates the end of the list the current known extension states.
	AmiListenerEventFailedACL                  = "FailedACL"                  // Raised when a request violates an ACL check.
	AmiListenerEventFAXSession                 = "FAXSession"                 // Raised in response to FAXSession manager command
	AmiListenerEventFAXSessionsComplete        = "FAXSessionsComplete"        // Raised when all FAXSession events are completed for a FAXSessions command
	AmiListenerEventFAXSessionsEntry           = "FAXSessionsEntry"           // A single list item for the FAXSessions AMI command
	AmiListenerEventFAXStats                   = "FAXStats"                   // Raised in response to FAXStats manager command
	AmiListenerEventFAXStatus                  = "FAXStatus"                  // Raised periodically during a fax transmission.
	AmiListenerEventFullyBooted                = "FullyBooted"                // Raised when all Asterisk initialization procedures have finished
	AmiListenerEventHangupHandlerPop           = "HangupHandlerPop"           // Raised when a hangup handler is removed from the handler stack by the CHANNEL() function.
	AmiListenerEventIdentifyDetail             = "IdentifyDetail"             // Provide details about an identify section.
	AmiListenerEventInvalidAccountID           = "InvalidAccountID"           // Raised when a request fails an authentication check due to an invalid account ID.
	AmiListenerEventInvalidTransport           = "InvalidTransport"           // Raised when a request attempts to use a transport not allowed by the Asterisk service.
	AmiListenerEventLoad                       = "Load"                       // Raised when a module has been loaded in Asterisk.
	AmiListenerEventLoadAverageLimit           = "LoadAverageLimit"           // Raised when a request fails because a configured load average limit has been reached.
	AmiListenerEventLocalOptimizationBegin     = "LocalOptimizationBegin"     // Raised when two halves of a Local Channel begin to optimize themselves out of the media path
	AmiListenerEventLocalOptimizationEnd       = "LocalOptimizationEnd"       // Raised when two halves of a Local Channel have finished optimizing themselves out of the media path.
	AmiListenerEventLogChannel                 = "LogChannel"                 // Raised when a logging channel is re-enabled after a reload operation.
	AmiListenerEventMCID                       = "MCID"                       // Published when a malicious call ID request arrives
	AmiListenerEventMeetMeEnd                  = "MeetmeEnd"                  // Raised when a MeetMe conference ends
	AmiListenerEventMeetMeJoin                 = "MeetmeJoin"                 // Raised when a user joins a MeetMe conference.
	AmiListenerEventMeetMeLeave                = "MeetmeLeave"                // Raised when a user leaves a MeetMe conference.
	AmiListenerEventMeetMeMute                 = "MeetmeMute"                 // Raised when a MeetMe user is muted or unmuted.
	AmiListenerEventMeetMeTalking              = "MeetmeTalking"              // Raised when a MeetMe user begins or ends talking.
	AmiListenerEventMeetMeTalkRequest          = "MeetmeTalkRequest"          // Raised when a MeetMe user has started talking.
	AmiListenerEventMemoryLimit                = "MemoryLimit"                // Raised when a request fails due to an internal memory allocation failure.
	AmiListenerEventMiniVoiceMail              = "MiniVoiceMail"              // Raised when a notification is sent out by a MiniVoiceMail application
	AmiListenerEventMonitorStart               = "MonitorStart"               // Raised when monitoring has started on a channel.
	AmiListenerEventMonitorStop                = "MonitorStop"                // Raised when monitoring has stopped on a channel
	AmiListenerEventMusicOnHoldStart           = "MusicOnHoldStart"           // Raised when music on hold has started on a channel.
	AmiListenerEventMusicOnHoldStop            = "MusicOnHoldStop"            // Raised when music on hold has stopped on a channel.
	AmiListenerEventMWIGet                     = "MWIGet"                     // Raised in response to a MWIGet command.
	AmiListenerEventMWIGetComplete             = "MWIGetComplete"             // Raised in response to a MWIGet command.
	AmiListenerEventNewAccountCode             = "NewAccountCode"             // Raised when a Channel's AccountCode is changed.
	AmiListenerEventParkedCallSwap             = "ParkedCallSwap"             // Raised when a channel takes the place of a previously parked channel
	AmiListenerEventPresenceStateChange        = "PresenceStateChange"        // Raised when a presence state changes
	AmiListenerEventPresenceStateListComplete  = "PresenceStateListComplete"  // Indicates the end of the list the current known extension states.
	AmiListenerEventPresenceStatus             = "PresenceStatus"             // Raised when a hint changes due to a presence state change.
	AmiListenerEventQueueCallerAbandon         = "QueueCallerAbandon"         // Raised when a caller abandons the queue.
	AmiListenerEventQueueCallerJoin            = "QueueCallerJoin"            // Raised when a caller joins a Queue.
	AmiListenerEventQueueCallerLeave           = "QueueCallerLeave"           // Raised when a caller leaves a Queue.
	AmiListenerEventQueueMemberRinginuse       = "QueueMemberRinginuse"       // Raised when a member's ringinuse setting is changed.
	AmiListenerEventReceiveFAX                 = "ReceiveFAX"                 // Raised when a receive fax operation has completed.
	AmiListenerEventRegistry                   = "Registry"                   // Raised when an outbound registration completes.
	AmiListenerEventReload                     = "Reload"                     // Raised when a module has been reloaded in Asterisk.
	AmiListenerEventRequestBadFormat           = "RequestBadFormat"           // Raised when a request is received with bad formatting.
	AmiListenerEventRequestNotAllowed          = "RequestNotAllowed"          // Raised when a request is not allowed by the service.
	AmiListenerEventRequestNotSupported        = "RequestNotSupported"        // Raised when a request fails due to some aspect of the requested item not being supported by the service.
	AmiListenerEventRTCPReceived               = "RTCPReceived"               // Raised when an RTCP packet is received.
	AmiListenerEventRTCPSent                   = "RTCPSent"                   // Raised when an RTCP packet is sent.
	AmiListenerEventSendFAX                    = "SendFAX"                    // Raised when a send fax operation has completed.
	AmiListenerEventSessionLimit               = "SessionLimit"               // Raised when a request fails due to exceeding the number of allowed concurrent sessions for that service.
	AmiListenerEventSessionTimeout             = "SessionTimeout"             // Raised when a SIP session times out.
	AmiListenerEventShutdown                   = "Shutdown"                   // Raised when Asterisk is shutdown or restarted.
	AmiListenerEventSIPQualifyPeerDone         = "SIPQualifyPeerDone"         // Raised when SIPQualifyPeer has finished qualifying the specified peer
	AmiListenerEventSpanAlarm                  = "SpanAlarm"                  // Raised when an alarm is set on a DAHDI span.
	AmiListenerEventSpanAlarmClear             = "SpanAlarmClear"             // Raised when an alarm is cleared on a DAHDI span.
	AmiListenerEventTransportDetail            = "TransportDetail"            // Provide details about an authentication section.
	AmiListenerEventUnexpectedAddress          = "UnexpectedAddress"          // Raised when a request has a different source address then what is expected for a session already in progress with a service
	AmiListenerEventUnload                     = "Unload"                     // Raised when a module has been unloaded in Asterisk.
)

const (
	AmiClassCommand                = "COMMAND"
	AmiClassSecurity               = "SECURITY"
	AmiClassCall                   = "CALL"
	AmiClassSystem                 = "SYSTEM"
	AmiClassUser                   = "USER"
	AmiClassDialPlan               = "DIALPLAN"
	AmiClassAgent                  = "AGENT"
	AmiClassAgi                    = "AGI"
	AmiClassAoc                    = "AOC"
	AmiClassCallDetailRecord       = "CDR"
	AmiClassChannelEventLogging    = "CEL"
	AmiClass                       = "CLASS"
	AmiClassReporting              = "REPORTING"
	AmiClassDualToneMultiFrequency = "DTMF"
)

// page 263
var (
	AmiClassCommands map[string][]string = map[string][]string{
		AmiClassCommand: {
			AmiListenerEventAorDetail,
			AmiListenerEventAorList,
			AmiListenerEventAorListComplete,
			AmiListenerEventAuthDetail,
			AmiListenerEventAuthList,
			AmiListenerEventAuthListComplete,
			AmiListenerEventBridgeInfoChannel,
			AmiListenerEventBridgeInfoComplete,
			AmiListenerEventContactList,
			AmiListenerEventContactListComplete,
			AmiListenerEventContactStatusDetail,
			AmiListenerEventDeviceStateListComplete,
			AmiListenerEventEndpointDetail,
			AmiListenerEventEndpointDetailComplete,
			AmiListenerEventEndpointList,
			AmiListenerEventEndpointListComplete,
			AmiListenerEventExtensionStateListComplete,
			AmiListenerEventIdentifyDetail,
			AmiListenerEventTransportDetail,
		},
	}
	AmiClassSecurities map[string][]string = map[string][]string{
		AmiClassSecurity: {
			AmiListenerEventUnexpectedAddress,
			AmiListenerEventChallengeResponseFailed,
			AmiListenerEventChallengeSent,
			AmiListenerEventFailedACL,
			AmiListenerEventInvalidAccountID,
			AmiListenerEventInvalidPassword,
			AmiListenerEventInvalidTransport,
			AmiListenerEventLoadAverageLimit,
			AmiListenerEventMemoryLimit,
			AmiListenerEventRequestBadFormat,
			AmiListenerEventRequestNotAllowed,
			AmiListenerEventRequestNotSupported,
			AmiListenerEventSessionLimit,
			AmiListenerEventSuccessfulAuth,
		},
	}
	AmiClassCalls map[string][]string = map[string][]string{
		AmiClassCall: {
			AmiListenerEventUnHold,
			AmiListenerEventUnParkedCall,
			AmiListenerEventAttendedTransfer,
			AmiListenerEventBlindTransfer,
			AmiListenerEventBridgeCreate,
			AmiListenerEventBridgeDestroy,
			AmiListenerEventBridgeEnter,
			AmiListenerEventBridgeLeave,
			AmiListenerEventBridgeMerge,
			AmiListenerEventBridgeVideoSourceUpdate,
			AmiListenerEventChanSpyStart,
			AmiListenerEventChanSpyStop,
			AmiListenerEventConfbridgeEnd,
			AmiListenerEventConfBridgeJoin,
			AmiListenerEventConfBridgeLeave,
			AmiListenerEventConfbridgeMute,
			AmiListenerEventConfbridgeRecord,
			AmiListenerEventConfbridgeStart,
			AmiListenerEventConfbridgeStopRecord,
			AmiListenerEventConfBridgeTalking,
			AmiListenerEventConfbridgeUnMute,
			AmiListenerEventCoreShowChannel,
			AmiListenerEventCoreShowChannelsComplete,
			AmiListenerEventDAHDIChannel,
			AmiListenerEventDeviceStateChange,
			AmiListenerEventDialBegin,
			AmiListenerEventDialEnd,
			AmiListenerEventDialState,
			AmiListenerEventExtensionStatus,
			AmiListenerEventFAXSessionsComplete,
			AmiListenerEventFAXStatus,
			AmiListenerEventHangup,
			AmiListenerEventHangupRequest,
			AmiListenerEventHold,
			AmiListenerEventLocalBridge,
			AmiListenerEventLocalOptimizationBegin,
			AmiListenerEventLocalOptimizationEnd,
			AmiListenerEventMCID,
			AmiListenerEventMeetMeEnd,
			AmiListenerEventMeetMeJoin,
			AmiListenerEventMeetMeLeave,
			AmiListenerEventMeetMeMute,
			AmiListenerEventMeetMeTalking,
			AmiListenerEventMeetMeTalkRequest,
			AmiListenerEventMessageWaiting,
			AmiListenerEventMiniVoiceMail,
			AmiListenerEventMonitorStart,
			AmiListenerEventMonitorStop,
			AmiListenerEventMusicOnHoldStart,
			AmiListenerEventMusicOnHoldStop,
			AmiListenerEventNewAccountCode,
			AmiListenerEventNewCallerId,
			AmiListenerEventNewChannel,
			AmiListenerEventNewConnectedLine,
			AmiListenerEventNewState,
			AmiListenerEventOriginateResponse,
			AmiListenerEventParkedCall,
			AmiListenerEventParkedCallGiveUp,
			AmiListenerEventParkedCallSwap,
			AmiListenerEventParkedCallTimeOut,
			AmiListenerEventPickup,
			AmiListenerEventPresenceStateChange,
			AmiListenerEventPresenceStateListComplete,
			AmiListenerEventPresenceStatus,
			AmiListenerEventReceiveFAX,
			AmiListenerEventRename,
			AmiListenerEventSendFAX,
			AmiListenerEventSessionTimeout,
			AmiListenerEventSIPQualifyPeerDone,
			AmiListenerEventSoftHangupRequest,
			AmiListenerEventStatus,
			AmiListenerEventStatusComplete,
			AmiListenerEventConfbridgeJoin,
			AmiListenerEventConfbridgeLeave,
		},
	}
	AmiClassSystems map[string][]string = map[string][]string{
		AmiClassSystem: {
			AmiListenerEventUnload,
			AmiListenerEventAlarm,
			AmiListenerEventAlarmClear,
			AmiListenerEventAuthMethodNotAllowed,
			AmiListenerEventContactStatus,
			AmiListenerEventDNDState,
			AmiListenerEventFullyBooted,
			AmiListenerEventLoad,
			AmiListenerEventLogChannel,
			AmiListenerEventPeerStatus,
			AmiListenerEventRegistry,
			AmiListenerEventReload,
			AmiListenerEventShutdown,
			AmiListenerEventSpanAlarm,
			AmiListenerEventSpanAlarmClear,
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
			AmiListenerEventHangupHandlerPop,
			AmiListenerEventHangupHandlerPush,
			AmiListenerEventHangupHandlerRun,
			AmiListenerEventNewExtension,
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
			AmiListenerEventQueueCallerAbandon,
			AmiListenerEventQueueCallerJoin,
			AmiListenerEventQueueCallerLeave,
			AmiListenerEventQueueMemberAdded,
			AmiListenerEventQueueMemberPause,
			AmiListenerEventQueueMemberPenalty,
			AmiListenerEventQueueMemberRemoved,
			AmiListenerEventQueueMemberRinginuse,
			AmiListenerEventQueueMemberStatus,
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
	AmiClassCallDetailRecords map[string][]string = map[string][]string{
		AmiClassCallDetailRecord: {
			AmiListenerEventCdr,
		},
	}
	AmiClassChannelEventLoggings map[string][]string = map[string][]string{
		AmiClassChannelEventLogging: {
			AmiListenerEventCel,
		},
	}
	AmiClasses map[string][]string = map[string][]string{
		AmiClass: {
			AmiListenerEventChannelTalkingStart,
			AmiListenerEventChannelTalkingStop,
		},
	}
	AmiClassReports map[string][]string = map[string][]string{
		AmiClassReporting: {
			AmiListenerEventConfbridgeList,
			AmiListenerEventFAXSession,
			AmiListenerEventFAXSessionsEntry,
			AmiListenerEventFAXStats,
			AmiListenerEventMWIGet,
			AmiListenerEventMWIGetComplete,
			AmiListenerEventRTCPReceived,
			AmiListenerEventRTCPSent,
		},
	}
	AmiClassDualToneMultiFrequencies map[string][]string = map[string][]string{
		AmiClassDualToneMultiFrequency: {
			AmiListenerEventDTMFBegin,
			AmiListenerEventDTMFEnd,
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

const (
	// Answer channel, Answers channel if not already in answer state. Returns -1 on channel failure, or 0 if successful.
	AmiAgiCommandAnswer = "ANSWER"
	// Interrupts Async AGI, Interrupts expected flow of Async AGI commands and returns control to previous source (typically, the PBX dialplan)
	AmiAgiCommandAsyncAgiBreak = "ASYNCAGI BREAK"
	// Returns status of the connected channel,
	// Returns the status of the specified channel name. If no channel name is given then returns the status of the current channel.
	// Return values:
	// 0 - Channel is down and available.
	// 1 - Channel is down, but reserved.
	// 2 - Channel is off hook.
	// 3 - Digits (or equivalent) have been dialed.
	// 4 - Line is ringing.
	// 5 - Remote end is ringing.
	// 6 - Line is up.
	// 7 - Line is busy.
	// Syntax: CHANNEL STATUS SAMPLE-CHANNEL-NAME
	AmiAgiCommandChannelStatus = "CHANNEL STATUS"
	// Sends audio file on channel and allows the listener to control the stream
	// Send the given file, allowing playback to be controlled by the given digits, if any. Use double quotes for the digits if you wish none to be permitted. If
	// offsets is provided then the audio will seek to offsets before play starts. Returns 0 if playback completes without a digit being pressed, or the ASCII
	// numerical value of the digit if one was pressed, or -1 on error or if the channel was disconnected. Returns the position where playback was terminated as endpoint.
	// It sets the following channel variables upon completion:
	// CPLAYBACKSTATUS - Contains the status of the attempt as a text string
	// 			- SUCCESS
	// 			- USERSTOPPED
	// 			- REMOTESTOPPED
	// 			- ERROR
	// CPLAYBACKOFFSET - Contains the offset in ms into the file where playback was at when it stopped. -1 is end of file
	// CPLAYBACKSTOPKEY - If the playback is stopped by the user this variable contains the key that was pressed.
	// Syntax: CONTROL STREAM FILE FILENAME ESCAPE_DIGITS SKIPMS FFCHAR REWCHR PAUSECHR OFFSETMS
	AmiAgiCommandControlStreamFile = "CONTROL STREAM FILE"
	// Removes database key/value
	// Deletes an entry in the Asterisk database for a given family and key.
	// Returns 1 if successful, 0 otherwise.
	// Syntax: DATABASE DEL FAMILY KEY
	AmiAgiCommandDatabaseDelete = "DATABASE DEL"
	// Removes database keytree/value
	// Deletes a family or specific keytree within a family in the Asterisk database
	// Returns 1 if successful, 0 otherwise.
	// Syntax: DATABASE DELTREE FAMILY KEYTREE
	AmiAgiCommandDatabaseDeleteTree = "DATABASE DELTREE"
	// Gets database value
	// Retrieves an entry in the Asterisk database for a given family and key.
	// Returns 0 if key is not set. Returns 1 if key is set and returns the variable in parenthesis
	// Example return code: 200 result=1 (test variable)
	// Syntax: DATABASE GET FAMILY KEY
	AmiAgiCommandDatabaseGet = "DATABASE GET"
	// Adds/updates database value
	// Adds or updates an entry in the Asterisk database for a given family, key, and value.
	// Returns 1 if successful, 0 otherwise.
	// Syntax: DATABASE PUT FAMILY KEY VALUE
	AmiAgiCommandDatabasePut = "DATABASE PUT"
	// Executes a given Application
	// Executes application with given options.
	// Returns whatever the application returns, or -2 on failure to find application
	// Syntax: EXEC APPLICATION OPTIONS
	AmiAgiCommandExecute = "EXEC"
	// Prompts for DTMF on a channel
	// Stream the given file, and receive DTMF data
	// Returns the digits received from the channel at the other end.
	// Syntax: GET DATA FILE TIMEOUT MAXDIGITS
	AmiAgiCommandGetData = "GET DATA"
	// Evaluates a channel expression
	// Evaluates the given expression against the channel specified by channel name, or the current channel if channel name is not provided.
	// Unlike GET VARIABLE, the expression is processed in a manner similar to dialplan evaluation, allowing complex and built-in variables to be accessed, e.g.
	// The time is ${EPOCH}
	// Returns 0 if no channel matching channel name exists, 1 otherwise.
	// Example return code: 200 result=1 (The time is 1578493800)
	// Syntax: GET FULL VARIABLE EXPRESSION CHANNELNAME
	AmiAgiCommandGetFullVariable = "GET FULL VARIABLE"
	// Stream file, prompt for DTMF, with timeout.
	// Behaves similar to STREAM FILE but used with a timeout option.
	// Syntax: GET OPTION FILENAME ESCAPE_DIGITS TIMEOUT
	AmiAgiCommandGetOption = "GET OPTION"
	// Gets a channel variable.
	// Returns 0 if variable name is not set. Returns 1 if variable name is set and returns the variable in parentheses.
	// Example return code: 200 result=1 (test variable)
	// Syntax: GET VARIABLE VARIABLENAME
	AmiAgiCommandGetVariable = "GET VARIABLE"
	// Cause the channel to execute the specified dialplan subroutine.
	// Cause the channel to execute the specified dialplan subroutine, returning to the dialplan with execution of a Return().
	// Syntax: GOSUB CONTEXT EXTENSION PRIORITY OPTIONAL-ARGUMENT
	AmiAgiCommandGoSub = "GOSUB"
	// Hangup a channel
	// Hangs up the specified channel. If no channel name is given, hangs up the current channel
	// Syntax: HANGUP CHANNELNAME
	AmiAgiCommandHangup = "HANGUP"
	// Does nothing.
	// Syntax: NOOP
	AmiAgiCommandNoop = "NOOP"
	// Receives one character from channels supporting it.
	// Receives a character of text on a channel. Most channels do not support the reception of text. Returns the decimal value of the character if one is received,
	// or 0 if the channel does not support text reception. Returns -1 only on error/hangup.
	// Syntax: RECEIVE CHAR TIMEOUT
	AmiAgiCommandReceiveChar = "RECEIVE CHAR"
	// Receives text from channels supporting it.
	// Receives a string of text on a channel. Most channels do not support the reception of text. Returns -1 for failure or 1 for success, and the string in parenthesis.
	// Syntax: RECEIVE TEXT TIMEOUT
	AmiAgiCommandReceiveText = "RECEIVE TEXT"
	// Records to a given file.
	// Record to a file until a given dtmf digit in the sequence is received. Returns -1 on hangup or error. The format will specify what kind of file will be recorded
	// The timeout is the maximum record time in milliseconds, or -1 for no timeout. offset samples is optional, and, if provided, will seek to the offset without
	// exceeding the end of the file. beep can take any value, and causes Asterisk to play a beep to the channel that is about to be recorded. silence is the
	// number of seconds of silence allowed before the function returns despite the lack of dtmf digits or reaching timeout. silence value must be preceded by s=
	// and is also optional.
	// Syntax: RECORD FILE FILENAME FORMAT ESCAPE_DIGITS TIMEOUT OFFSET_SAMPLES BEEP S=SILENCE
	AmiAgiCommandRecordFile = "RECORD FILE"
	// Says a given character string.
	// Say a given character string, returning early if any of the given DTMF digits are received on the channel. Returns 0 if playback completes without a digit
	// being pressed, or the ASCII numerical value of the digit if one was pressed or -1 on error/hangup
	// Syntax: SAY ALPHA NUMBER ESCAPE_DIGITS
	AmiAgiCommandSayAlpha = "SAY ALPHA"
	// Says a given date.
	// Say a given date, returning early if any of the given DTMF digits are received on the channel. Returns 0 if playback completes without a digit being
	// pressed, or the ASCII numerical value of the digit if one was pressed or -1 on error/hangup
	// Syntax: SAY DATE DATE ESCAPE_DIGITS
	AmiAgiCommandSayDate = "SAY DATE"
	// Says a given time as specified by the format given
	// Say a given time, returning early if any of the given DTMF digits are received on the channel. Returns 0 if playback completes without a digit being pressed
	// or the ASCII numerical value of the digit if one was pressed or -1 on error/hangup.
	// Syntax: SAY DATETIME TIME ESCAPE_DIGITS FORMAT TIMEZONE
	AmiAgiCommandSayDateTime = "SAY DATETIME"
	// Says a given digit string
	// Say a given digit string, returning early if any of the given DTMF digits are received on the channel. Returns 0 if playback completes without a digit being
	// pressed, or the ASCII numerical value of the digit if one was pressed or -1 on error/hangup.
	// Syntax: SAY DIGITS NUMBER ESCAPE_DIGITS
	AmiAgiCommandSayDigits = "SAY DIGITS"
	// Says a given number
	// Say a given number, returning early if any of the given DTMF digits are received on the channel. Returns 0 if playback completes without a digit being
	// pressed, or the ASCII numerical value of the digit if one was pressed or -1 on error/hangup.
	// Syntax: SAY NUMBER NUMBER ESCAPE_DIGITS GENDER
	AmiAgiCommandSayNumber = "SAY NUMBER"
	// Says a given character string with phonetics.
	// Say a given character string with phonetics, returning early if any of the given DTMF digits are received on the channel. Returns 0 if playback completes
	// without a digit pressed, the ASCII numerical value of the digit if one was pressed, or -1 on error/hangup.
	// Syntax: SAY PHONETIC STRING ESCAPE_DIGITS
	AmiAgiCommandSayPhonetic = "SAY PHONETIC"
	// Says a given time.
	// Say a given time, returning early if any of the given DTMF digits are received on the channel. Returns 0 if playback completes without a digit being pressed,
	// or the ASCII numerical value of the digit if one was pressed or -1 on error/hangup.
	// Syntax: SAY TIME TIME ESCAPE_DIGITS
	AmiAgiCommandSayTime = "SAY TIME"
	// Sends images to channels supporting it.
	// Sends the given image on a channel. Most channels do not support the transmission of images. Returns 0 if image is sent, or if the channel does not
	// support image transmission. Returns -1 only on error/hangup. Image names should not include extensions.
	// Syntax: SEND IMAGE IMAGE
	AmiAgiCommandSendImage = "SEND IMAGE"
	// Sends text to channels supporting it.
	// Sends the given text on a channel. Most channels do not support the transmission of text. Returns 0 if text is sent, or if the channel does not support text
	// transmission. Returns -1 only on error/hangup.
	// Syntax: SEND TEXT TEXT TO SEND
	AmiAgiCommandSendText = "SEND TEXT"
	// Autohangup channel in some time.
	// Cause the channel to automatically hangup at time seconds in the future. Of course it can be hung up before then as well. Setting to 0 will cause the
	// autohangup feature to be disabled on this channel.
	// Syntax: SET AUTOHANGUP TIME
	AmiAgiCommandSetAutoHangup = "SET AUTOHANGUP"
	// Sets callerid for the current channel.
	// Changes the callerid of the current channel.
	// Syntax: SET CALLERID NUMBER
	AmiAgiCommandSetCallerId = "SET CALLERID"
	// Sets channel context.
	// Sets the context for continuation upon exiting the application
	// Syntax: SET CONTEXT DESIRED CONTEXT
	AmiAgiCommandSetContext = "SET CONTEXT"
	// Changes channel extension.
	// Changes the extension for continuation upon exiting the application.
	// Syntax: SET EXTENSION NEW EXTENSION
	AmiAgiCommandSetExtension = "SET EXTENSION"
	// Enable/Disable Music on hold generator
	// Enables/Disables the music on hold generator. If class is not specified, then the default music on hold class will be used. This generator will be stopped
	// automatically when playing a file.
	// Always returns 0.
	// Syntax: SET MUSIC CLASS
	AmiAgiCommandSetMusic = "SET MUSIC"
	// Set channel dialplan priority
	// Changes the priority for continuation upon exiting the application. The priority must be a valid priority or label.
	// Syntax: SET PRIORITY PRIORITY
	AmiAgiCommandSetPriority = "SET PRIORITY"
	// Sets a channel variable.
	// Sets a variable to the current channel.
	// Syntax: SET VARIABLE VARIABLENAME VALUE
	AmiAgiCommandSetVariable = "SET VARIABLE"
	// Activates a grammar.
	// Activates the specified grammar on the speech object.
	// Syntax: SPEECH ACTIVATE GRAMMAR GRAMMAR NAME
	AmiAgiCommandSpeechActivateGrammar = "SPEECH ACTIVATE GRAMMAR"
	// Creates a speech object
	// Create a speech object to be used by the other Speech AGI commands
	// Syntax: SPEECH CREATE ENGINE
	AmiAgiCommandSpeechCreate = "SPEECH CREATE"
	// Deactivates a grammar.
	// Deactivates the specified grammar on the speech object.
	// Syntax: SPEECH DEACTIVATE GRAMMAR GRAMMAR NAME
	AmiAgiCommandSpeechDeactivateGrammar = "SPEECH DEACTIVATE GRAMMAR"
	// Destroys a speech object.
	// Destroy the speech object created by SPEECH CREATE.
	// Syntax: SPEECH DESTROY
	AmiAgiCommandSpeechDestroy = "SPEECH DESTROY"
	// Loads a grammar.
	// Loads the specified grammar as the specified name
	// Syntax: SPEECH LOAD GRAMMAR GRAMMAR NAME PATH TO GRAMMAR
	AmiAgiCommandSpeechLoadGrammar = "SPEECH LOAD GRAMMAR"
	// Recognizes speech
	// Plays back given prompt while listening for speech and dtmf.
	// Syntax: SPEECH RECOGNIZE PROMPT TIMEOUT OFFSET
	AmiAgiCommandSpeechRecognize = "SPEECH RECOGNIZE"
	// Sets a speech engine setting.
	// Set an engine-specific setting.
	// Syntax: SPEECH SET NAME VALUE
	AmiAgiCommandSpeechSet = "SPEECH SET"
	// Unloads a grammar.
	// Unloads the specified grammar.
	// Syntax: SPEECH UNLOAD GRAMMAR GRAMMAR NAME
	AmiAgiCommandSpeechUnloadGrammar = "SPEECH UNLOAD GRAMMAR"
	// Sends audio file on channel.
	// Send the given file, allowing playback to be interrupted by the given digits, if any. Returns 0 if playback completes without a digit being pressed, or the
	// ASCII numerical value of the digit if one was pressed, or -1 on error or if the channel was disconnected. If music-on-hold is playing before calling stream file
	// it will be automatically stopped and will not be restarted after completion.
	// It sets the following channel variables upon completion:
	// - PLAYBACKSTATUS - The status of the playback attempt as a text string.
	// 		SUCCESS
	// 		FAILED
	// Syntax: STREAM FILE
	AmiAgiCommandStreamFile = "STREAM FILE"
	// Toggles TDD mode (for the deaf).
	// Enable/Disable TDD transmission/reception on a channel. Returns 1 if successful, or 0 if channel is not TDD-capable.
	// Syntax: TDD MODE BOOLEAN
	AmiAgiCommandTddMode = "TDD MODE"
	// Logs a message to the asterisk verbose log.
	// Sends message to the console via verbose message system. level is the verbose level (1-4). Always returns 1
	// Syntax: VERBOSE MESSAGE LEVEL
	AmiAgiCommandVerbose = "VERBOSE"
	// Waits for a digit to be pressed.
	// Waits up to timeout milliseconds for channel to receive a DTMF digit. Returns -1 on channel failure, 0 if no digit is received in the timeout, or the numerical
	// value of the ascii of the digit if one is received. Use -1 for the timeout value if you desire the call to block indefinitely.
	// Syntax: WAIT FOR DIGIT TIMEOUT
	AmiAgiCommandWaitForDigit = "WAIT FOR DIGIT"
)

var (
	AmiAgiCommands []string = []string{
		AmiAgiCommandAnswer,
		AmiAgiCommandAsyncAgiBreak,
		AmiAgiCommandChannelStatus,
		AmiAgiCommandControlStreamFile,
		AmiAgiCommandDatabaseDelete,
		AmiAgiCommandDatabaseDeleteTree,
		AmiAgiCommandDatabaseGet,
		AmiAgiCommandDatabasePut,
		AmiAgiCommandExecute,
		AmiAgiCommandGetData,
		AmiAgiCommandGetFullVariable,
		AmiAgiCommandGetOption,
		AmiAgiCommandGetVariable,
		AmiAgiCommandGoSub,
		AmiAgiCommandHangup,
		AmiAgiCommandNoop,
		AmiAgiCommandReceiveChar,
		AmiAgiCommandReceiveText,
		AmiAgiCommandRecordFile,
		AmiAgiCommandSayAlpha,
		AmiAgiCommandSayDate,
		AmiAgiCommandSayDateTime,
		AmiAgiCommandSayDigits,
		AmiAgiCommandSayNumber,
		AmiAgiCommandSayPhonetic,
		AmiAgiCommandSayTime,
		AmiAgiCommandSendImage,
		AmiAgiCommandSendText,
		AmiAgiCommandSetAutoHangup,
		AmiAgiCommandSetCallerId,
		AmiAgiCommandSetContext,
		AmiAgiCommandSetExtension,
		AmiAgiCommandSetMusic,
		AmiAgiCommandSetPriority,
		AmiAgiCommandSetVariable,
		AmiAgiCommandSpeechActivateGrammar,
		AmiAgiCommandSpeechCreate,
		AmiAgiCommandSpeechDeactivateGrammar,
		AmiAgiCommandSpeechDestroy,
		AmiAgiCommandSpeechLoadGrammar,
		AmiAgiCommandSpeechRecognize,
		AmiAgiCommandSpeechSet,
		AmiAgiCommandSpeechUnloadGrammar,
		AmiAgiCommandStreamFile,
		AmiAgiCommandTddMode,
		AmiAgiCommandVerbose,
		AmiAgiCommandWaitForDigit,
	}
)
