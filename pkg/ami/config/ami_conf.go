package config

import (
	"time"
)

// Time format styles for date and time formatting.
const (
	// DateTimeFormat20060102T150405 represents the time format "2006-01-02T15:04:05",
	// which includes the full date in the format YYYY-MM-DD followed by the time in the
	// format HH:MM:SS.
	DateTimeFormat20060102T150405 = "2006-01-02T15:04:05"

	// DateTimeFormat20060102150405 represents the time format "2006-01-02 15:04:05",
	// which includes the full date in the format YYYY-MM-DD followed by the time in the
	// format HH:MM:SS with a space between the date and time.
	DateTimeFormat20060102150405 = "2006-01-02 15:04:05"

	// DateTimeFormat20060102 represents the date format "2006-01-02",
	// which includes the full date in the format YYYY-MM-DD without any time information.
	DateTimeFormat20060102 = "2006-01-02"

	// DateTimeFormat200601 represents the month format "2006-01",
	// which includes the year and month in the format YYYY-MM without any day or time information.
	DateTimeFormat200601 = "2006-01"
)

// Timezone constants representing default timezones for specific regions.
const (
	// DefaultTimezoneVietnam is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Vietnam, which is "Asia/Ho_Chi_Minh".
	DefaultTimezoneVietnam = "Asia/Ho_Chi_Minh"

	// DefaultTimezoneNewYork is a constant that holds the IANA Time Zone identifier
	// for the default timezone in New York, USA, which is "America/New_York".
	DefaultTimezoneNewYork = "America/New_York"

	// DefaultTimezoneLondon is a constant that holds the IANA Time Zone identifier
	// for the default timezone in London, United Kingdom, which is "Europe/London".
	DefaultTimezoneLondon = "Europe/London"

	// DefaultTimezoneTokyo is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Tokyo, Japan, which is "Asia/Tokyo".
	DefaultTimezoneTokyo = "Asia/Tokyo"

	// DefaultTimezoneSydney is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Sydney, Australia, which is "Australia/Sydney".
	DefaultTimezoneSydney = "Australia/Sydney"

	// DefaultTimezoneParis is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Paris, France, which is "Europe/Paris".
	DefaultTimezoneParis = "Europe/Paris"

	// DefaultTimezoneMoscow is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Moscow, Russia, which is "Europe/Moscow".
	DefaultTimezoneMoscow = "Europe/Moscow"

	// DefaultTimezoneLosAngeles is a constant that holds the IANA Time Zone identifier
	// for the default timezone in Los Angeles, USA, which is "America/Los_Angeles".
	DefaultTimezoneLosAngeles = "America/Los_Angeles"
)

// AMI protocol keys used for defining actions, events, responses, and other properties
// in Asterisk Manager Interface (AMI) communication.
const (
	// AmiActionKey represents the key used for specifying actions in AMI messages.
	AmiActionKey = "Action"

	// AmiEventKey represents the key used for specifying events in AMI messages.
	AmiEventKey = "Event"

	// AmiResponseKey represents the key used for specifying responses in AMI messages.
	AmiResponseKey = "Response"

	// AmiActionIdKey represents the key used for specifying Action IDs in AMI messages,
	// allowing correlation between actions and their corresponding responses.
	AmiActionIdKey = "ActionID"

	// AmiLoginKey represents the key used for specifying login information in AMI messages.
	AmiLoginKey = "Login"

	// AmiCallManagerKey represents the key used for identifying Asterisk Call Manager in AMI messages.
	AmiCallManagerKey = "Asterisk Call Manager"

	// AmiAuthTypeKey represents the key used for specifying authentication types in AMI messages.
	AmiAuthTypeKey = "AuthType"

	// AmiFilenameKey represents the key used for specifying filenames in AMI messages.
	AmiFilenameKey = "Filename"

	// AmiFullyBootedKey represents the key used for indicating that Asterisk is fully booted
	// and ready to handle commands and events.
	AmiFullyBootedKey = "Fully Booted"

	AmiStatusSuccessKey = "success"
	AmiStatusFailedKey  = "failed"
)

// AMI Network constants used for indicating the network type in Asterisk Manager Interface (AMI) configurations.
const (
	// AmiNetworkTcpKey represents the network key "tcp," indicating the TCP network.
	AmiNetworkTcpKey = "tcp"

	// AmiNetworkUdpKey represents the network key "udp," indicating the UDP network.
	AmiNetworkUdpKey = "udp"

	// AmiNetworkTcp4Key represents the network key "tcp4," indicating the IPv4 TCP network.
	AmiNetworkTcp4Key = "tcp4"

	// AmiNetworkTcp6Key represents the network key "tcp6," indicating the IPv6 TCP network.
	AmiNetworkTcp6Key = "tcp6"

	// AmiNetworkUdp4Key represents the network key "udp4," indicating the IPv4 UDP network.
	AmiNetworkUdp4Key = "udp4"

	// AmiNetworkUdp6Key represents the network key "udp6," indicating the IPv6 UDP network.
	AmiNetworkUdp6Key = "udp6"

	// AmiNetworkIpKey represents the network key "ip," indicating the generic IP network.
	AmiNetworkIpKey = "ip"

	// AmiNetworkIp4Key represents the network key "ip4," indicating the IPv4 network.
	AmiNetworkIp4Key = "ip4"

	// AmiNetworkIp6Key represents the network key "ip6," indicating the IPv6 network.
	AmiNetworkIp6Key = "ip6"

	// AmiNetworkUnixKey represents the network key "unix," indicating the Unix network.
	AmiNetworkUnixKey = "unix"

	// AmiNetworkUnixGramKey represents the network key "unixgram," indicating the Unix datagram network.
	AmiNetworkUnixGramKey = "unixgram"

	// AmiNetworkUnixPacketKey represents the network key "unixpacket," indicating the Unix packet network.
	AmiNetworkUnixPacketKey = "unixpacket"
)

// AMI Protocol Key constants used for indicating the protocol in Asterisk Manager Interface (AMI) URLs.
const (
	// AmiProtocolHttpKey represents the protocol key "http://," indicating the HTTP protocol.
	AmiProtocolHttpKey = "http://"

	// AmiProtocolHttpsKey represents the protocol key "https://," indicating the HTTPS protocol.
	AmiProtocolHttpsKey = "https://"
)

// AMI Class constants used for categorizing Asterisk Manager Interface (AMI) actions and events.
const (
	// AmiClassCommand represents the class "COMMAND," indicating actions related to command execution.
	AmiClassCommand = "COMMAND"

	// AmiClassSecurity represents the class "SECURITY," indicating actions related to security.
	AmiClassSecurity = "SECURITY"

	// AmiClassCall represents the class "CALL," indicating actions and events related to calls.
	AmiClassCall = "CALL"

	// AmiClassSystem represents the class "SYSTEM," indicating actions and events related to the Asterisk system.
	AmiClassSystem = "SYSTEM"

	// AmiClassUser represents the class "USER," indicating actions and events related to users.
	AmiClassUser = "USER"

	// AmiClassDialPlan represents the class "DIALPLAN," indicating actions and events related to the dial plan.
	AmiClassDialPlan = "DIALPLAN"

	// AmiClassAgent represents the class "AGENT," indicating actions and events related to agents.
	AmiClassAgent = "AGENT"

	// AmiClassAgi represents the class "AGI," indicating actions and events related to Asterisk Gateway Interface (AGI) scripts.
	AmiClassAgi = "AGI"

	// AmiClassAoc represents the class "AOC," indicating actions and events related to Advice of Charge (AOC).
	AmiClassAoc = "AOC"

	// AmiClassCallDetailRecord represents the class "CDR," indicating actions and events related to Call Detail Records (CDR).
	AmiClassCallDetailRecord = "CDR"

	// AmiClassChannelEventLogging represents the class "CEL," indicating actions and events related to Channel Event Logging (CEL).
	AmiClassChannelEventLogging = "CEL"

	// AmiClass represents the class "CLASS," indicating generic actions and events.
	AmiClass = "CLASS"

	// AmiClassReporting represents the class "REPORTING," indicating actions and events related to reporting.
	AmiClassReporting = "REPORTING"

	// AmiClassDualToneMultiFrequency represents the class "DTMF," indicating actions and events related to Dual-Tone Multi-Frequency (DTMF).
	AmiClassDualToneMultiFrequency = "DTMF"
)

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

const (
	NetworkTimeoutAfterSeconds              = time.Second * 3 // default is 3 seconds
	AmiMaxTimeoutInMsForCall                = 100000          // 100000 milliseconds
	AmiMinTimeoutInMsForCall                = 10000           // 10000 milliseconds
	AmiSignalLetter                         = "\r\n"
	AmiSignalLetters                        = "\r\n\r\n"
	AmiMaxConcurrencyMillis                 = 60000 // milliseconds
	AmiDigitExtensionRegexDefault    string = "^SIP/\\d{4}"
	AmiDigitExtensionRegexWithDigits string = "^SIP/\\d{%v}"
	AmiPubSubKeyRef                         = "ami-key"
	AmiOmitemptyKeyRef                      = "omitempty"
	AmiTagKeyRef                            = "ami"
)

const (
	AmiErrorFieldRequired           string = "%v is required"
	AmiErrorInvalidProtocol         string = "Invalid protocol"
	AmiErrorInvalidChanspy          string = "Invalid chanspy join"
	AmiErrorProtocolMessage         string = "Protocol must have values: %v"
	AmiErrorChanspyMessage          string = "Chanspy only supported values: %v"
	AmiErrorConsumeEvent            string = "Ami can not consume event, reason failed: %v"
	AmiErrorUsernameRequired        string = "(Ami Authentication). username was missing"
	AmiErrorPasswordRequired        string = "(Ami Authentication). password was missing"
	AmiErrorLoginFailedMessage      string = "(Ami Authentication). login failed for reason: %v"
	AmiErrorLogoutFailedMessage     string = "(Ami Authentication). logout failed for reason: %v"
	AmiErrorBreakSocketIgnoredEvent string = "Event '%v' was broken while fetching from server"
	AmiErrorMissingSocketEvent      string = "Event %v was missing while fetching from server, response = %v"
	AmiErrorNoExtensionConfigured   string = "There's no sip peers configured"
	AmiErrorNoExtensionsConfigured  string = "There's no extensions configured"
	AmiErrorLoginFailed             string = "(Ami Authentication). login failed"
	AmiErrorPingFailed              string = "(Ami Authentication). Ping failed for reason: %v"
)

// AMI Channel Protocols constants used for indicating the protocol of a channel
// in Asterisk Manager Interface (AMI) responses.
const (
	// AmiSIPChannelProtocol represents the channel protocol "SIP," indicating
	// that the channel uses the SIP protocol.
	AmiSIPChannelProtocol string = "SIP"

	// AmiIAXChannelProtocol represents the channel protocol "IAX," indicating
	// that the channel uses the IAX (Inter-Asterisk eXchange) protocol.
	AmiIAXChannelProtocol string = "IAX"

	// AmiZapChannelProtocol represents the channel protocol "Zap," indicating
	// that the channel uses the Zap (Zaptel) protocol.
	AmiZapChannelProtocol string = "Zap"

	// AmiH323ChannelProtocol represents the channel protocol "H323," indicating
	// that the channel uses the H.323 protocol.
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

// AMI Context constants used for indicating the context of a communication
// in Asterisk Manager Interface (AMI) responses.
const (
	// AmiContextOutbound represents the communication context "outbound-allroutes,"
	// indicating an outbound communication across all routes.
	AmiContextOutbound = "outbound-allroutes"

	// AmiContextDefault represents the communication context "default,"
	// indicating the default communication context.
	AmiContextDefault = "default"

	// AmiContextFromInternal represents the communication context "from-internal,"
	// indicating an internal communication context.
	AmiContextFromInternal = "from-internal"

	// AmiContextFromExternal represents the communication context "from-external,"
	// indicating an external communication context.
	AmiContextFromExternal = "from-external"
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

// AGIControl represents the control type to playback actions
type AGIControl string

const (
	// Stop the playback operation
	Stop AGIControl = "stop"
	// Forward move the current position in the media forward.
	Forward AGIControl = "forward"
	// Reverse move the current position in the media backward.
	Reverse AGIControl = "reverse"
	// Pause pause/unpause the playback operation.
	Pause AGIControl = "pause"
	// Restart the playback operation.
	Restart AGIControl = "restart"
)

// AMI Communication Direction constants used for indicating the direction of a communication
// in Asterisk Manager Interface (AMI) responses.
const (
	// AmiOutboundDirection represents the communication direction "outbound,"
	// indicating an outbound communication.
	AmiOutboundDirection = "outbound"

	// AmiInboundDirection represents the communication direction "inbound,"
	// indicating an inbound communication.
	AmiInboundDirection = "inbound"

	// AmiUnknownDirection represents the communication direction "Unknown,"
	// indicating an unknown or undefined communication direction.
	AmiUnknownDirection = "Unknown"
)

// AMI Direction Type constants used for indicating the direction of a communication
// in Asterisk Manager Interface (AMI) responses.
const (
	// AmiTypeOutboundNormalDirection represents the direction type "outbound_normal,"
	// indicating a normal outbound communication.
	AmiTypeOutboundNormalDirection = "outbound_normal"

	// AmiTypeInboundDialDirection represents the direction type "inbound_dial,"
	// indicating an inbound communication initiated from an application on the user's local machine.
	AmiTypeInboundDialDirection = "inbound_dial"

	// AmiTypeInboundQueueDirection represents the direction type "inbound_queue,"
	// indicating an inbound communication from a queue and not from a real user.
	AmiTypeInboundQueueDirection = "inbound_queue"

	// AmiTypeChanSpyDirection represents the direction type "chan_spy,"
	// indicating a communication direction for ChanSpy functionality.
	AmiTypeChanSpyDirection = "chan_spy"
)

var (
	AmiCallDirection map[string]bool = map[string]bool{
		AmiOutboundDirection: true,
		AmiInboundDirection:  true,
	}
)

// AMI Extension State constants used for indicating the state of an extension in
// Asterisk Manager Interface (AMI) responses.
const (
	// AmiExtensionRemoved represents the extension state "Removed," indicating
	// that the extension has been removed.
	AmiExtensionRemoved = -2

	// AmiExtensionDeactivated represents the extension state "Deactivated," indicating
	// that the extension hint has been removed.
	AmiExtensionDeactivated = -1

	// AmiExtensionNotInUse represents the extension state "Not In Use," indicating
	// that no devices associated with the extension are In-Use, Busy, or Idle.
	AmiExtensionNotInUse = 0

	// AmiExtensionInUse represents the extension state "In Use," indicating
	// that one or more devices associated with the extension are In-Use.
	AmiExtensionInUse = 1 << 0

	// AmiExtensionBusy represents the extension state "Busy," indicating
	// that all devices associated with the extension are Busy.
	AmiExtensionBusy = 1 << 1

	// AmiExtensionUnavailable represents the extension state "Unavailable," indicating
	// that all devices associated with the extension are Unavailable or Unregistered.
	AmiExtensionUnavailable = 1 << 2

	// AmiExtensionRinging represents the extension state "Ringing," indicating
	// that all devices associated with the extension are Ringing.
	AmiExtensionRinging = 1 << 3

	// AmiExtensionOnHold represents the extension state "On Hold," indicating
	// that all devices associated with the extension are On-Hold.
	AmiExtensionOnHold = 1 << 4
)

// AMI Device State constants used for indicating the state of a device in
// Asterisk Manager Interface (AMI) responses.
const (
	// AmiDeviceStateUnknown represents the device state "Unknown," indicating
	// that the device is valid but the channel didn't know its state.
	AmiDeviceStateUnknown = 0

	// AmiDeviceStateNotInUse represents the device state "Not In Use," indicating
	// that the device is not currently in use.
	AmiDeviceStateNotInUse = 1

	// AmiDeviceStateInUse represents the device state "In Use," indicating
	// that the device is currently in use.
	AmiDeviceStateInUse = 2

	// AmiDeviceStateBusy represents the device state "Busy," indicating
	// that the device is busy.
	AmiDeviceStateBusy = 3

	// AmiDeviceStateInvalid represents the device state "Invalid," indicating
	// that the device state is invalid.
	AmiDeviceStateInvalid = 4

	// AmiDeviceStateUnavailable represents the device state "Unavailable," indicating
	// that the device is unavailable.
	AmiDeviceStateUnavailable = 5

	// AmiDeviceStateRinging represents the device state "Ringing," indicating
	// that the device is currently ringing.
	AmiDeviceStateRinging = 6

	// AmiDeviceStateRingInUse represents the device state "Ring In Use," indicating
	// that the device is ringing and currently in use.
	AmiDeviceStateRingInUse = 7

	// AmiDeviceStateOnHold represents the device state "On Hold," indicating
	// that the device is currently on hold.
	AmiDeviceStateOnHold = 8

	// AmiDeviceStateTotal represents the total number of device states and is
	// used for testing purposes.
	AmiDeviceStateTotal = 9 // Total num of device states, used for testing
)

// AMI Device State string constants used for indicating the state of a device in
// Asterisk Manager Interface (AMI) responses.
const (
	// AmiDeviceStateUnknownString represents the device state "UNKNOWN," indicating
	// that the state of the device is unknown or undefined.
	AmiDeviceStateUnknownString = "UNKNOWN"

	// AmiDeviceStateNotInUseString represents the device state "NOT_INUSE," indicating
	// that the device is not currently in use.
	AmiDeviceStateNotInUseString = "NOT_INUSE"

	// AmiDeviceStateInUseString represents the device state "INUSE," indicating
	// that the device is currently in use.
	AmiDeviceStateInUseString = "INUSE"

	// AmiDeviceStateBusyString represents the device state "BUSY," indicating
	// that the device is busy.
	AmiDeviceStateBusyString = "BUSY"

	// AmiDeviceStateInvalidString represents the device state "INVALID," indicating
	// that the device state is invalid.
	AmiDeviceStateInvalidString = "INVALID"

	// AmiDeviceStateUnavailableString represents the device state "UNAVAILABLE," indicating
	// that the device is unavailable.
	AmiDeviceStateUnavailableString = "UNAVAILABLE"

	// AmiDeviceStateRingingString represents the device state "RINGING," indicating
	// that the device is currently ringing.
	AmiDeviceStateRingingString = "RINGING"

	// AmiDeviceStateRingInUseString represents the device state "RINGINUSE," indicating
	// that the device is ringing and currently in use.
	AmiDeviceStateRingInUseString = "RINGINUSE"

	// AmiDeviceStateOnHoldString represents the device state "ONHOLD," indicating
	// that the device is currently on hold.
	AmiDeviceStateOnHoldString = "ONHOLD"
)

// AMI Channel State constants used for indicating the state of a channel in
// Asterisk Manager Interface (AMI) responses.
const (
	// AmiChannelStateDown represents the channel state "Down," indicating that the
	// channel is down and available.
	AmiChannelStateDown = 0

	// AmiChannelStateReserved represents the channel state "Reserved," indicating
	// that the channel is down, but reserved.
	AmiChannelStateReserved = 1

	// AmiChannelStateOffHook represents the channel state "Off Hook," indicating
	// that the channel is off hook.
	AmiChannelStateOffHook = 2

	// AmiChannelStateDialing represents the channel state "Dialing," indicating
	// that digits (or equivalent) have been dialed.
	AmiChannelStateDialing = 3

	// AmiChannelStateRing represents the channel state "Ring," indicating
	// that the line is ringing.
	AmiChannelStateRing = 4

	// AmiChannelStateRinging represents the channel state "Ringing," indicating
	// that the remote end is ringing.
	AmiChannelStateRinging = 5

	// AmiChannelStateUp represents the channel state "Up," indicating
	// that the line is up.
	AmiChannelStateUp = 6

	// AmiChannelStateBusy represents the channel state "Busy," indicating
	// that the line is busy.
	AmiChannelStateBusy = 7

	// AmiChannelStateDialingOffHook represents the channel state "Dialing Off Hook,"
	// indicating that digits (or equivalent) have been dialed while off hook.
	AmiChannelStateDialingOffHook = 8

	// AmiChannelStatePreRing represents the channel state "Pre Ring," indicating
	// that the channel has detected an incoming call and is waiting for ring.
	AmiChannelStatePreRing = 9

	// AmiChannelStateMute represents the channel state "Mute," indicating
	// that voice data transmission is muted.
	AmiChannelStateMute = 1 << 16 // Do not transmit voice data
)

// Device state, extension state and cdr strings for printing
var (
	AmiExtensionStatusCodes map[int]string = map[int]string{
		AmiExtensionDeactivated: "Extension not found",
		AmiExtensionNotInUse:    "Idle",
		AmiExtensionInUse:       "In Use",
		AmiExtensionBusy:        "Busy",
		AmiExtensionUnavailable: "Unavailable",
		AmiExtensionRinging:     "Ringing",
		AmiExtensionOnHold:      "On Hold",
	}
	AmiDeviceStatesString map[string]string = map[string]string{
		AmiDeviceStateUnknownString:     "Unknown",     // Valid, but unknown state
		AmiDeviceStateNotInUseString:    "Not in use",  // Not used
		AmiDeviceStateInUseString:       "In use",      // In use
		AmiDeviceStateBusyString:        "Busy",        // Busy
		AmiDeviceStateInvalidString:     "Invalid",     // Invalid - not known to Asterisk
		AmiDeviceStateUnavailableString: "Unavailable", // Unavailable (not registered)
		AmiDeviceStateRingingString:     "Ringing",     // Ring, ring, ring
		AmiDeviceStateRingInUseString:   "Ring+Inuse",  // Ring and in use
		AmiDeviceStateOnHoldString:      "On Hold",     // On Hold
	}
	AmiDeviceStatesText map[int]string = map[int]string{
		AmiDeviceStateUnknown:     AmiDeviceStatesString[AmiDeviceStateUnknownString],
		AmiDeviceStateNotInUse:    AmiDeviceStatesString[AmiDeviceStateNotInUseString],
		AmiDeviceStateInUse:       AmiDeviceStatesString[AmiDeviceStateInUseString],
		AmiDeviceStateBusy:        AmiDeviceStatesString[AmiDeviceStateBusyString],
		AmiDeviceStateInvalid:     AmiDeviceStatesString[AmiDeviceStateInvalidString],
		AmiDeviceStateUnavailable: AmiDeviceStatesString[AmiDeviceStateUnavailableString],
		AmiDeviceStateRinging:     AmiDeviceStatesString[AmiDeviceStateRingingString],
		AmiDeviceStateRingInUse:   AmiDeviceStatesString[AmiDeviceStateRingInUseString],
		AmiDeviceStateOnHold:      AmiDeviceStatesString[AmiDeviceStateOnHoldString],
	}
	AmiChannelStatesText map[int]string = map[int]string{
		AmiChannelStateDown:           "down",
		AmiChannelStateReserved:       "reserved",
		AmiChannelStateOffHook:        "off-hook",
		AmiChannelStateDialing:        "dialing",
		AmiChannelStateRing:           "ring",
		AmiChannelStateRinging:        "ringing",
		AmiChannelStateUp:             "up",
		AmiChannelStateBusy:           "busy",
		AmiChannelStateDialingOffHook: "dialing-off-hook",
		AmiChannelStatePreRing:        "pre-ring",
		AmiChannelStateMute:           "mute",
	}
	AmiCdrDispositionText map[string]bool = map[string]bool{
		AmiCdrDispositionNoAnswer:   true,
		AmiCdrDispositionFailed:     true,
		AmiCdrDispositionBusy:       true,
		AmiCdrDispositionAnswered:   true,
		AmiCdrDispositionCongestion: true,
	}
)

// AMI Peer Status constants used for indicating the status of a peer (SIP/IAX) in
// Asterisk Manager Interface (AMI) responses.
const (
	// AmiPeerStatusUnknown represents the peer status "Unknown," indicating that the
	// status of the peer is unknown or undefined.
	AmiPeerStatusUnknown = "Unknown"

	// AmiPeerStatusRegistered represents the peer status "Registered," indicating that
	// the peer is successfully registered with the Asterisk server.
	AmiPeerStatusRegistered = "Registered"

	// AmiPeerStatusUnregistered represents the peer status "Unregistered," indicating
	// that the peer is not currently registered with the Asterisk server.
	AmiPeerStatusUnregistered = "Unregistered"

	// AmiPeerStatusRejected represents the peer status "Rejected," indicating that the
	// registration request from the peer was rejected by the Asterisk server.
	AmiPeerStatusRejected = "Rejected"

	// AmiPeerStatusLagged represents the peer status "Lagged," indicating that there
	// is a delay or lag in communication with the peer.
	AmiPeerStatusLagged = "Lagged"

	// AmiPeerStatusReachable represents the peer status "Reachable," indicating that
	// the peer is reachable and responsive.
	AmiPeerStatusReachable = "Reachable"
)

// AMI Call Detail Record (CDR) disposition constants used for indicating the result
// of a call in Asterisk Manager Interface (AMI) responses.
const (
	// AmiCdrDispositionNoAnswer represents the CDR disposition "NO ANSWER," indicating
	// that the channel was not answered. This is the default disposition.
	AmiCdrDispositionNoAnswer = "NO ANSWER"

	// AmiCdrDispositionFailed represents the CDR disposition "FAILED," indicating that
	// the channel attempted to dial but the call failed.
	AmiCdrDispositionFailed = "FAILED"

	// AmiCdrDispositionBusy represents the CDR disposition "BUSY," indicating that
	// the channel attempted to dial but the remote party was busy.
	AmiCdrDispositionBusy = "BUSY"

	// AmiCdrDispositionAnswered represents the CDR disposition "ANSWERED," indicating
	// that the channel was answered. The hang-up cause will no longer impact the disposition of the CDR.
	AmiCdrDispositionAnswered = "ANSWERED"

	// AmiCdrDispositionCongestion represents the CDR disposition "CONGESTION," indicating
	// that the channel attempted to dial but the remote party was congested.
	AmiCdrDispositionCongestion = "CONGESTION"
)

// AMI AMA (Automatic Message Accounting) flag constants used for indicating the purpose
// of Call Detail Record (CDR) entries in Asterisk Manager Interface (AMI) responses.
const (
	// AmiAmaFlagOmit represents the AMA flag "OMIT," indicating that this CDR should be ignored.
	AmiAmaFlagOmit = "OMIT"

	// AmiAmaFlagBilling represents the AMA flag "BILLING," indicating that this CDR contains
	// valid billing data.
	AmiAmaFlagBilling = "BILLING"

	// AmiAmaFlagDocumentation represents the AMA flag "DOCUMENTATION," indicating that
	// this CDR is for documentation purposes.
	AmiAmaFlagDocumentation = "DOCUMENTATION"
)

// AMI last application keys used for identifying the last application executed in
// Asterisk Manager Interface (AMI) responses.
const (
	// AmiLastApplicationDial represents the key used for identifying the last executed
	// application as "Dial" in AMI responses.
	AmiLastApplicationDial = "Dial"

	// AmiLastApplicationQueue represents the key used for identifying the last executed
	// application as "Queue" in AMI responses.
	AmiLastApplicationQueue = "Queue"

	// AmiLastApplicationChanSpy represents the key used for identifying the last executed
	// application as "ChanSpy" in AMI responses.
	AmiLastApplicationChanSpy = "ChanSpy"
)

// AmiChanspySpy, AmiChanspyBarge, and AmiChanspyWhisper are constants representing different modes
// of ChanSpy functionality in Asterisk Manager Interface (AMI).
const (
	// AmiChanspySpy represents the "spy" mode in ChanSpy, allowing monitoring without intervention.
	AmiChanspySpy = "spy"

	// AmiChanspyBarge represents the "barge" mode in ChanSpy, allowing the interceptor to join an ongoing call.
	AmiChanspyBarge = "barge"

	// AmiChanspyWhisper represents the "whisper" mode in ChanSpy, allowing the interceptor to listen and talk to
	// one party in a call without the other party hearing the interceptor's voice.
	AmiChanspyWhisper = "whisper"
)

var (
	AmiChanspy map[string]bool = map[string]bool{
		AmiChanspySpy:     true,
		AmiChanspyBarge:   true,
		AmiChanspyWhisper: true,
	}
)
