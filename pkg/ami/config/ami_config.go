package config

import (
	"time"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/fatal"
)

const (
	DateTimeFormatYYYYMMDDTHHMMSS = "2006-01-02T15:04:05"
	DateTimeFormatYYYYMMDDHHMMSS  = "2006-01-02 15:04:05"
	DateTimeFormatYYYYMMDD        = "2006-01-02"
	DateTimeFormatYYYYMM          = "2006-01"
	DefaultTimezoneAsia           = "Asia/Ho_Chi_Minh"
)

const (
	AmiActionKey      = "Action"
	AmiEventKey       = "Event"
	AmiResponseKey    = "Response"
	AmiActionIdKey    = "ActionID"
	AmiLoginKey       = "Login"
	AmiCallManagerKey = "Asterisk Call Manager"
	AmiAuthTypeKey    = "AuthType"
	AmiFilenameKey    = "Filename"
	AmiFullyBootedKey = "Fully Booted"
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
	AmiPubSubKeyRef    = "ami-key"
	AmiOmitemptyKeyRef = "omitempty"
	AmiTagKeyRef       = "ami"
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
	AmiMaxTimeoutInMsForCall   = 100000          // 100000 milliseconds
	AmiMinTimeoutInMsForCall   = 10000           // 10000 milliseconds
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
	AmiContextOutbound     = "outbound-allroutes"
	AmiContextDefault      = "default"
	AmiContextFromInternal = "from-internal"
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

const (
	AmiSignalLetter  = "\r\n"
	AmiSignalLetters = "\r\n\r\n"
)

const (
	AmiMaxConcurrencyMillis = 60000 // millis
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

const (
	AmiOutboundDirection = "outbound"
	AmiInboundDirection  = "inbound"
	AmiUnknownDirection  = "Unknown"
)

const (
	AmiTypeOutboundNormalDirection = "outbound_normal"
	AmiTypeInboundDialDirection    = "inbound_dial"  // from application on user local machine
	AmiTypeInboundQueueDirection   = "inbound_queue" // from queue, not real user
)

var (
	AmiCallDirection map[string]bool = map[string]bool{
		AmiOutboundDirection: true,
		AmiInboundDirection:  true,
	}
)

const (
	AmiExtensionRemoved     = -2     // Extension removed
	AmiExtensionDeactivated = -1     // Extension hint removed
	AmiExtensionNotInUse    = 0      // No device In-Use or Busy or Idle
	AmiExtensionInUse       = 1 << 0 // One or more devices In-Use
	AmiExtensionBusy        = 1 << 1 // All devices Busy
	AmiExtensionUnavailable = 1 << 2 // All devices Unavailable or Unregistered
	AmiExtensionRinging     = 1 << 3 // All devices Ringing
	AmiExtensionOnHold      = 1 << 4 // All devices On-Hold
)

// Device States
// The order of these states may not change because they are included
// in Asterisk events which may be transmitted across the network to other servers.
const (
	AmiDeviceStateUnknown     = 0 // Device is valid but channel didn't know state
	AmiDeviceStateNotInUse    = 1 // Device is not used
	AmiDeviceStateInUse       = 2 // Device is in use
	AmiDeviceStateBusy        = 3 // Device is busy
	AmiDeviceStateInvalid     = 4 // Device is invalid
	AmiDeviceStateUnavailable = 5 // Device is unavailable
	AmiDeviceStateRinging     = 6 // Device is ringing
	AmiDeviceStateRingInUse   = 7 // Device is ringing *and* in use
	AmiDeviceStateOnHold      = 8 // Device is on hold
	AmiDeviceStateTotal       = 9 // Total num of device states, used for testing
)

const (
	AmiDeviceStateUnknownString     = "UNKNOWN"
	AmiDeviceStateNotInUseString    = "NOT_INUSE"
	AmiDeviceStateInUseString       = "INUSE"
	AmiDeviceStateBusyString        = "BUSY"
	AmiDeviceStateInvalidString     = "INVALID"
	AmiDeviceStateUnavailableString = "UNAVAILABLE"
	AmiDeviceStateRingingString     = "RINGING"
	AmiDeviceStateRingInUseString   = "RINGINUSE"
	AmiDeviceStateOnHoldString      = "ONHOLD"
)

const (
	AmiChannelStateDown           = 0       // Channel is down and available
	AmiChannelStateReserved       = 1       // Channel is down, but reserved
	AmiChannelStateOffHook        = 2       // Channel is off hook
	AmiChannelStateDialing        = 3       // Digits (or equivalent) have been dialed
	AmiChannelStateRing           = 4       // Line is ringing
	AmiChannelStateRinging        = 5       // Remote end is ringing
	AmiChannelStateUp             = 6       // Line is up
	AmiChannelStateBusy           = 7       // Line is busy
	AmiChannelStateDialingOffHook = 8       // Digits (or equivalent) have been dialed while offhook
	AmiChannelStatePreRing        = 9       // Channel has detected an incoming call and is waiting for ring
	AmiChannelStateMute           = 1 << 16 // Do not transmit voice data
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

const (
	AmiPeerStatusUnknown      = "Unknown"
	AmiPeerStatusRegistered   = "Registered"
	AmiPeerStatusUnregistered = "Unregistered"
	AmiPeerStatusRejected     = "Rejected"
	AmiPeerStatusLagged       = "Lagged"
	AmiPeerStatusReachable    = "Reachable"
)

const (
	AmiCdrDispositionNoAnswer   = "NO ANSWER"  // The channel was not answered. This is the default disposition.
	AmiCdrDispositionFailed     = "FAILED"     // The channel attempted to dial but the call failed.
	AmiCdrDispositionBusy       = "BUSY"       // The channel attempted to dial but the remote party was busy.
	AmiCdrDispositionAnswered   = "ANSWERED"   // The channel was answered. The hang up cause will no longer impact the disposition of the CDR.
	AmiCdrDispositionCongestion = "CONGESTION" // The channel attempted to dial but the remote party was congested.
)

const (
	AmiAmaFlagOmit          = "OMIT"          // This CDR should be ignored.
	AmiAmaFlagBilling       = "BILLING"       // This CDR contains valid billing data.
	AmiAmaFlagDocumentation = "DOCUMENTATION" // This CDR is for documentation purposes.
)

const (
	AmiLastApplicationDial  = "Dial"
	AmiLastApplicationQueue = "Queue"
)

const (
	AmiChanspySpy     = "spy"
	AmiChanspyBarge   = "barge"
	AmiChanspyWhisper = "whisper"
)

var (
	AmiChanspy map[string]bool = map[string]bool{
		AmiChanspySpy:     true,
		AmiChanspyBarge:   true,
		AmiChanspyWhisper: true,
	}
)
