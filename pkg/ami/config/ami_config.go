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
	AmiSignalLetter = "\r\n"
)
