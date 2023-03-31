package config

const (
	// Execute Asterisk CLI Command.
	// Run a CLI command.
	// Syntax:
	// Action: Command
	// ActionID: <value>
	// Command: <value>
	AmiActionCommand = "Command"
	// Originate a call.
	// Generates an outgoing call to a Extension/Context/Priority or Application/Data
	// Syntax:
	// Action: Originate
	// ActionID: <value>
	// Channel: <value>
	// Exten: <value>
	// Context: <value>
	// Priority: <value>
	// Application: <value>
	// Data: <value>
	// Timeout: <value>
	// CallerID: <value>
	// Variable: <value>
	// Account: <value>
	// EarlyMedia: <value>
	// Async: <value>
	// Codecs: <value>
	// ChannelId: <value>
	// OtherChannelId: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned
	// Channel - Channel name to call.
	// Exten - Extension to use (requires Context and Priority)
	// Context - Context to use (requires Exten and Priority)
	// Priority - Priority to use (requires Exten and Context)
	// Application - Application to execute
	// Data - Data to use (requires Application)
	// Timeout - How long to wait for call to be answered (in ms.).
	// CallerID - Caller ID to be set on the outgoing channel.
	// Variable - Channel variable to set, multiple Variable: headers are allowed.
	// Account - Account code.
	// EarlyMedia - Set to true to force call bridge on early media..
	// Async - Set to true for fast origination.
	// Codecs - Comma-separated list of codecs to use for this call.
	// ChannelId - Channel UniqueId to be set on the channel.
	// OtherChannelId - Channel UniqueId to be set on the second local channel.
	AmiActionOriginate = "Originate"
	// Set absolute timeout.
	// Hangup a channel after a certain time. Acknowledges set time with Timeout Set message
	// Syntax:
	// Action: AbsoluteTimeout
	// ActionID: <value>
	// Channel: <value>
	// Timeout: <value>
	AmiActionAbsoluteTimeout = "AbsoluteTimeout"
	// Sets an agent as no longer logged in.
	// Syntax:
	// Action: AgentLogoff
	// ActionID: <value>
	// Agent: <value>
	// Soft: <value>
	AmiActionAgentLogOff = "AgentLogoff"
	// Lists agents and their status.
	// Will list info about all defined agents.
	// Syntax:
	// Action: Agents
	// ActionID: <value>
	AmiActionAgents = "Agents"
	// Add an AGI command to execute by Async AGI.
	// Add an AGI command to the execute queue of the channel in Async AGI.
	// Action: AGI
	// ActionID: <value>
	// Channel: <value>
	// Command: <value>
	// CommandID: <value>
	AmiActionAgi = "AGI"
	// Generate an Advice of Charge message on a channel.
	// Generates an AOC-D or AOC-E message on a channel.
	// Syntax:
	// Action: AOCMessage
	// ActionID: <value>
	// Channel: <value>
	// ChannelPrefix: <value>
	// MsgType: <value>
	// ChargeType: <value>
	// UnitAmount(0): <value>
	// UnitType(0): <value>
	// CurrencyName: <value>
	// CurrencyAmount: <value>
	// CurrencyMultiplier: <value>
	// TotalType: <value>
	// AOCBillingId: <value>
	// ChargingAssociationId: <value>
	// ChargingAssociationNumber: <value>
	// ChargingAssociationPlan: <value>
	AmiActionAocMessage = "AOCMessage"
	// Attended transfer.
	// Syntax:
	// Action: Atxfer
	// ActionID: <value>
	// Channel: <value>
	// Exten: <value>
	// Context: <value>
	AmiActionAtxfer = "Atxfer"
	// Blind transfer channel(s) to the given destination
	// Redirect all channels currently bridged to the specified channel to the specified destination.
	// Syntax:
	// Action: BlindTransfer
	// Channel: <value>
	// Context: <value>
	// Exten: <value>
	AmiActionBlindTransfer = "BlindTransfer"
	// Bridge two channels already in the PBX.
	// Syntax:
	// Action: Bridge
	// ActionID: <value>
	// Channel1: <value>
	// Channel2: <value>
	// Tone: <value>
	AmiActionBridge = "Bridge"
	// Destroy a bridge.
	// Deletes the bridge, causing channels to continue or hang up.
	// Syntax:
	// Action: BridgeDestroy
	// ActionID: <value>
	// BridgeUniqueid: <value>
	AmiActionBridgeDestroy = "BridgeDestroy"
	// Get information about a bridge.
	// Returns detailed information about a bridge and the channels in it
	// Syntax:
	// Action: BridgeInfo
	// ActionID: <value>
	// BridgeUniqueid: <value>
	AmiActionBridgeInfo = "BridgeInfo"
	// Kick a channel from a bridge.
	// The channel is removed from the bridge
	// Syntax:
	// Action: BridgeKick
	// ActionID: <value>
	// [BridgeUniqueid:] <value>
	// Channel: <value>
	AmiActionBridgeKick = "BridgeKick"
	// Get a list of bridges in the system
	// Returns a list of bridges, optionally filtering on a bridge type.
	// Syntax:
	// Action: BridgeList
	// ActionID: <value>
	// BridgeType: <value>
	AmiActionBridgeList = "BridgeList"
	// List available bridging technologies and their statuses
	// Returns detailed information about the available bridging technologies.
	// Syntax:
	// Action: BridgeTechnologyList
	// ActionID: <value>
	AmiActionBridgeTechnologyList = "BridgeTechnologyList"
	// Suspend a bridging technology
	// Marks a bridging technology as suspended, which prevents subsequently created bridges from using it.
	// Syntax:
	// Action: BridgeTechnologySuspend
	// ActionID: <value>
	// BridgeTechnology: <value>
	AmiActionBridgeTechnologySuspend = "BridgeTechnologySuspend"
	// Unsuspend a bridging technology.
	// Clears a previously suspended bridging technology, which allows subsequently created bridges to use it.
	// Syntax:
	// Action: BridgeTechnologyUnsuspend
	// ActionID: <value>
	// BridgeTechnology: <value>
	AmiActionBridgeTechnologyUnsuspend = "BridgeTechnologyUnsuspend"
	// Cancel an attended transfer.
	// Cancel an attended transfer. Note, this uses the configured cancel attended transfer feature option (atxferabort) to cancel the transfer. If not available this
	// action will fail.
	// Syntax:
	// Action: CancelAtxfer
	// ActionID: <value>
	// Channel: <value>
	AmiActionCancelAtxfer = "CancelAtxfer"
	// Generate Challenge for MD5 Auth.
	// Generate a challenge for MD5 authentication.
	// Syntax:
	// Action: Challenge
	// ActionID: <value>
	// AuthType: <value>
	AmiActionChallenge = "Challenge"
	// Change monitoring filename of a channel.
	// This action may be used to change the file started by a previous 'Monitor' action.
	// Syntax:
	// Action: ChangeMonitor
	// ActionID: <value>
	// Channel: <value>
	// File: <value>
	AmiActionChangeMonitor = "ChangeMonitor"
	// Kick a Confbridge user.
	// Syntax:
	// Action: ConfbridgeKick
	// ActionID: <value>
	// Conference: <value>
	// Channel: <value>
	AmiActionConfbridgeKick = "ConfbridgeKick"
	// List participants in a conference.
	// Lists all users in a particular ConfBridge conference. ConfbridgeList will follow as separate events, followed by a final event called ConfbridgeListComplete.
	// Syntax:
	// Action: ConfbridgeList
	// ActionID: <value>
	// Conference: <value>
	AmiActionConfbridgeList = "ConfbridgeList"
	// List active conferences
	// Lists data about all active conferences. ConfbridgeListRooms will follow as separate events, followed by a final event called ConfbridgeListRoomsComplete.
	// Syntax:
	// Action: ConfbridgeListRooms
	// ActionID: <value>
	AmiActionConfbridgeListRooms = "ConfbridgeListRooms"
	// Lock a Confbridge conference.
	// Syntax:
	// Action: ConfbridgeLock
	// ActionID: <value>
	// Conference: <value>
	AmiActionConfbridgeLock = "ConfbridgeLock"
	// Mute a Confbridge user
	// Syntax:
	// Action: ConfbridgeMute
	// ActionID: <value>
	// Conference: <value>
	// Channel: <value>
	AmiActionConfbridgeMute = "ConfbridgeMute"
	// Set a conference user as the single video source distributed to all other participants
	// Syntax:
	// Action: ConfbridgeSetSingleVideoSrc
	// ActionID: <value>
	// Conference: <value>
	// Channel: <value>
	AmiActionConfbridgeSetSingleVideoSrc = "ConfbridgeSetSingleVideoSrc"
	// Start recording a Confbridge conference.
	// Start recording a conference. If recording is already present an error will be returned. If RecordFile is not provided, the default record file specified in the
	// conference's bridge profile will be used, if that is not present either a file will automatically be generated in the monitor directory.
	// Syntax:
	// Action: ConfbridgeStartRecord
	// ActionID: <value>
	// Conference: <value>
	// [RecordFile:] <value>
	AmiActionConfbridgeStartRecord = "ConfbridgeStartRecord"
	// Stop recording a Confbridge conference
	// Syntax:
	// Action: ConfbridgeStopRecord
	// ActionID: <value>
	// Conference: <value>
	AmiActionConfbridgeStopRecord = "ConfbridgeStopRecord"
	// Unlock a Confbridge conference.
	// Syntax:
	// Action: ConfbridgeUnlock
	// ActionID: <value>
	// Conference: <value>
	AmiActionConfbridgeUnlock = "ConfbridgeUnlock"
	// Unmute a Confbridge user.
	// Syntax:
	// Action: ConfbridgeUnmute
	// ActionID: <value>
	// Conference: <value>
	// Channel: <value>
	AmiActionConfbridgeUnmute = "ConfbridgeUnmute"
	// Control the playback of a file being played to a channel.
	// Control the operation of a media file being played back to a channel. Note that this AMI action does not initiate playback of media to channel, but rather
	// controls the operation of a media operation that was already initiated on the channel.
	// Syntax:
	// Action: ControlPlayback
	// ActionID: <value>
	// Channel: <value>
	// Control: <value>
	AmiActionControlPlayback = "ControlPlayback"
	// Show PBX core settings (version etc).
	// Query for Core PBX settings.
	// Syntax:
	// Action: CoreSettings
	// ActionID: <value>
	AmiActionCoreSettings = "CoreSettings"
	// List currently active channels.
	// List currently defined channels and some information about them.
	// Syntax:
	// Action: CoreShowChannels
	// ActionID: <value>
	AmiActionCoreShowChannels = "CoreShowChannels"
	// Show PBX core status variables
	// Query for Core PBX status.
	// Syntax:
	// Action: CoreStatus
	// ActionID: <value>
	AmiActionCoreStatus = "CoreStatus"
	// Creates an empty file in the configuration directory.
	// This action will create an empty file in the configuration directory. This action is intended to be used before an UpdateConfig action.
	// Syntax:
	// Action: CreateConfig
	// ActionID: <value>
	// Filename: <value>
	AmiActionCreateConfig = "CreateConfig"
	// Dial over DAHDI channel while off hook
	// Generate DTMF control frames to the bridged peer.
	// Syntax:
	// Action: DAHDIDialOffhook
	// ActionID: <value>
	// DAHDIChannel: <value>
	// Number: <value>
	AmiActionDAHDIDialOffhook = "DAHDIDialOffhook"
	// Toggle DAHDI channel Do Not Disturb status OFF.
	// Equivalent to the CLI command "dahdi set dnd channel off".
	// Syntax:
	// Action: DAHDIDNDoff
	// ActionID: <value>
	// DAHDIChannel: <value>
	AmiActionDAHDIDNDoff = "DAHDIDNDoff"
	// Toggle DAHDI channel Do Not Disturb status ON.
	// Equivalent to the CLI command "dahdi set dnd channel on".
	// Syntax:
	// Action: DAHDIDNDon
	// ActionID: <value>
	// DAHDIChannel: <value>
	AmiActionDAHDIDNDon = "DAHDIDNDon"
	// Hangup DAHDI Channel
	// Simulate an on-hook event by the user connected to the channel.
	// Syntax:
	// Action: DAHDIHangup
	// ActionID: <value>
	// DAHDIChannel: <value>
	AmiActionDAHDIHangup = "DAHDIHangup"
	// Fully Restart DAHDI channels (terminates calls)
	// Equivalent to the CLI command "dahdi restart"
	// Syntax:
	// Action: DAHDIRestart
	// ActionID: <value>
	AmiActionDAHDIRestart = "DAHDIRestart"
	// Show status of DAHDI channels.
	// Similar to the CLI command "dahdi show channels".
	// Syntax:
	// Action: DAHDIShowChannels
	// ActionID: <value>
	// DAHDIChannel: <value>
	AmiActionDAHDIShowChannels = "DAHDIShowChannels"
	// Transfer DAHDI Channel
	// Simulate a flash hook event by the user connected to the channel.
	// Syntax:
	// Action: DAHDITransfer
	// ActionID: <value>
	// DAHDIChannel: <value>
	AmiActionDAHDITransfer = "DAHDITransfer"
	// Delete DB entry.
	// Syntax:
	// Action: DBDel
	// ActionID: <value>
	// Family: <value>
	// Key: <value>
	AmiActionDBDel = "DBDel"
	// Delete DB Tree.
	// Syntax:
	// Action: DBDelTree
	// ActionID: <value>
	// Family: <value>
	// Key: <value>
	AmiActionDBDelTree = "DBDelTree"
	// Get DB Entry
	// Syntax:
	// Action: DBGet
	// ActionID: <value>
	// Family: <value>
	// Key: <value>
	AmiActionDBGet = "DBGet"
	// Put DB entry.
	// Syntax:
	// Action: DBPut
	// ActionID: <value>
	// Family: <value>
	// Key: <value>
	// Val: <value>
	AmiActionDBPut = "DBPut"
	// List the current known device states
	// This will list out all known device states in a sequence of DeviceStateChange events. When finished, a DeviceStateListComplete event will be emitted
	// Syntax:
	// Action: DeviceStateList
	// ActionID: <value>
	AmiActionDeviceStateList = "DeviceStateList"
	// Add an extension to the dialplan
	// Syntax:
	// Action: DialplanExtensionAdd
	// ActionID: <value>
	// Context: <value>
	// Extension: <value>
	// Priority: <value>
	// Application: <value>
	// [ApplicationData:] <value>
	// [Replace:] <value>
	AmiActionDialplanExtensionAdd = "DialplanExtensionAdd"
	// Remove an extension from the dialplan
	// Syntax:
	// Action: DialplanExtensionRemove
	// ActionID: <value>
	// Context: <value>
	// Extension: <value>
	// [Priority:] <value>
	AmiActionDialplanExtensionRemove = "DialplanExtensionRemove"
	// Control Event Flow.
	// Enable/Disable sending of events to this manager client.
	// Syntax:
	// Action: Events
	// ActionID: <value>
	// EventMask: <value>
	AmiActionEvents = "Events"
	// Check Extension Status.
	// Report the extension state for given extension. If the extension has a hint, will use devicestate to check the status of the device connected to the extension.
	// Will return an Extension Status message. The response will include the hint for the extension and the status.
	// Syntax:
	// Action: ExtensionState
	// ActionID: <value>
	// Exten: <value>
	// Context: <value>
	AmiActionExtensionState = "ExtensionState"
	// List the current known extension states.
	// This will list out all known extension states in a sequence of ExtensionStatus events. When finished, a ExtensionStateListComplete event will be emitted
	// Syntax:
	// Action: ExtensionStateList
	// ActionID: <value>
	AmiActionExtensionStateList = "ExtensionStateList"
	// Responds with a detailed description of a single FAX session
	// Provides details about a specific FAX session. The response will include a common subset of the output from the CLI command 'fax show session
	// <session_number>' for each technology. If the FAX technology used by this session does not include a handler for FAXSession, then this action will fail.
	// Syntax:
	// Action: FAXSession
	// ActionID: <value>
	// SessionNumber: <value>
	AmiActionFAXSession = "FAXSession"
	// Lists active FAX sessions
	// Will generate a series of FAXSession events with information about each FAXSession. Closes with a FAXSessionsComplete event which includes a count
	// of the included FAX sessions. This action works in the same manner as the CLI command 'fax show sessions'
	// Syntax:
	// Action: FAXSessions
	// ActionID: <value>
	AmiActionFAXSessions = "FAXSessions"
	// Responds with fax statistics
	// Provides FAX statistics including the number of active sessions, reserved sessions, completed sessions, failed sessions, and the number of
	// receive/transmit attempts. This command provides all of the non-technology specific information provided by the CLI command 'fax show stats'
	// Syntax:
	// Action: FAXStats
	// ActionID: <value>
	AmiActionFAXStats = "FAXStats"
	// Dynamically add filters for the current manager session.
	// The filters added are only used for the current session. Once the connection is closed the filters are removed
	// This command requires the system permission because this command can be used to create filters that may bypass filters defined in manager.conf
	// Syntax:
	// Action: Filter
	// ActionID: <value>
	// Operation: <value>
	// Filter: <value>
	AmiActionFilter = "Filter"
	// Retrieve configuration.
	// This action will dump the contents of a configuration file by category and contents or optionally by specified category only. In the case where a category
	// name is non-unique, a filter may be specified to match only categories with matching variable values.
	// Syntax:
	// Action: GetConfig
	// ActionID: <value>
	// Filename: <value>
	// Category: <value>
	// Filter: <value>
	AmiActionGetConfig = "GetConfig"
	// Retrieve configuration (JSON format).
	// This action will dump the contents of a configuration file by category and contents in JSON format or optionally by specified category only. This only makes
	// sense to be used using raw man over the HTTP interface. In the case where a category name is non-unique, a filter may be specified to match only
	// categories with matching variable values.
	// Syntax:
	// Action: GetConfigJSON
	// ActionID: <value>
	// Filename: <value>
	// Category: <value>
	// Filter: <value>
	AmiActionGetConfigJson = "GetConfigJSON"
	// Gets a channel variable or function value
	// Get the value of a channel variable or function return.
	// Syntax:
	// Action: Getvar
	// ActionID: <value>
	// Channel: <value>
	// Variable: <value>
	AmiActionGetVar = "Getvar"
	// Hangup channel.
	// Syntax:
	// Action: Hangup
	// ActionID: <value>
	// Channel: <value>
	// Cause: <value>
	AmiActionHangup = "Hangup"
	// Show IAX Netstats.
	// Show IAX channels network statistics.
	// Syntax:
	// Action: IAXnetstats
	AmiActionIAXnetstats = "IAXnetstats"
	// List all the IAX peers.
	// Syntax:
	// Action: IAXpeerlist
	// ActionID: <value>
	AmiActionIAXpeerlist = "IAXpeerlist"
	// List IAX peers.
	// Syntax:
	// Action: IAXpeers
	// ActionID: <value>
	AmiActionIAXpeers = "IAXpeers"
	// Show IAX registrations.
	// Syntax:
	// Action: IAXregistry
	// ActionID: <value>
	AmiActionIAXregistry = "IAXregistry"
	// Sends a message to a Jabber Client.
	// Syntax:
	// Action: JabberSend
	// ActionID: <value>
	// Jabber: <value>
	// Message: <value>
	// JID: <value>
	AmiActionJabberSend = "JabberSend"
	// List categories in configuration file.
	// This action will dump the categories in a given file.
	// Syntax:
	// Action: ListCategories
	// ActionID: <value>
	// Filename: <value>
	AmiActionListCategories = "ListCategories"
	// List available manager commands
	// Returns the action name and synopsis for every action that is available to the user.
	// Syntax:
	// Action: ListCommands
	// ActionID: <value>
	AmiActionListCommands = "ListCommands"
	// Optimize away a local channel when possible.
	// A local channel created with "/n" will not automatically optimize away. Calling this command on the local channel will clear that flag and allow it to optimize
	// away if it's bridged or when it becomes bridged
	// Syntax:
	// Action: LocalOptimizeAway
	// ActionID: <value>
	// Channel: <value>
	AmiActionLocalOptimizeAway = "LocalOptimizeAway"
	// Reload and rotate the Asterisk logger.
	// Reload and rotate the logger. Analogous to the CLI command 'logger rotate'.
	// Syntax:
	// Action: LoggerRotate
	// ActionID: <value>
	AmiActionLoggerRotate = "LoggerRotate"
	// Login Manager.
	// Syntax:
	// Action: Login
	// ActionID: <value>
	// Username: <value>
	// Secret: <value>
	// Arguments:
	// ActionID - ActionID for this transaction. Will be returned.
	// Username - Username to login with as specified in manager.conf.
	// Secret - Secret to login with as specified in manager.conf.
	AmiActionLogin = "Login"
	// Logoff Manager.
	// Logoff the current manager session
	// Syntax:
	// Action: Logoff
	// ActionID: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	AmiActionLogoff = "Logoff"
	// Check Mailbox Message Count
	// Checks a voicemail account for new messages.
	// Returns number of urgent, new and old messages
	// Message: Mailbox Message Count
	// Mailbox: mailboxid
	// UrgentMessages: count
	// NewMessages: count
	// OldMessages: count
	// Syntax:
	// Action: MailboxCount
	// ActionID: <value>
	// Mailbox: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Mailbox - Full mailbox ID mailbox@vm-context.
	AmiActionMailboxCount = "MailboxCount"
	// Check mailbox.
	// Checks a voicemail account for status.
	// Returns whether there are messages waiting.
	// Message: Mailbox Status.
	// Mailbox: mailboxid.
	// Waiting: 0 if messages waiting, 1 if no messages waiting.
	// Syntax:
	// Action: MailboxStatus
	// ActionID: <value>
	// Mailbox: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Mailbox - Full mailbox ID mailbox@vm-context.
	AmiActionMailboxStatus = "MailboxStatus"
	// List participants in a conference.
	// Lists all users in a particular MeetMe conference. MeetmeList will follow as separate events, followed by a final event called MeetmeListComplete
	// Syntax:
	// Action: MeetmeList
	// ActionID: <value>
	// [Conference:] <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned
	// Conference - Conference number
	AmiActionMeetMeList = "MeetmeList"
	// List active conferences.
	// Lists data about all active conferences. MeetmeListRooms will follow as separate events, followed by a final event called MeetmeListRoomsComplete
	// Syntax:
	// Action: MeetmeListRooms
	// ActionID: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	AmiActionMeetMeListRooms = "MeetmeListRooms"
	// Mute a Meetme user.
	// Syntax:
	// Action: MeetmeMute
	// ActionID: <value>
	// Meetme: <value>
	// Usernum: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	AmiActionMeetMeMute = "MeetmeMute"
	// Unmute a Meetme user.
	// Syntax:
	// Action: MeetmeUnmute
	// ActionID: <value>
	// Meetme: <value>
	// Usernum: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	AmiActionMeetMeUnmute = "MeetmeUnmute"
	// Send an out of call message to an endpoint.
	// Syntax:
	// Action: MessageSend
	// ActionID: <value>
	// To: <value>
	// From: <value>
	// Body: <value>
	// Base64Body: <value>
	// Variable: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// To - The URI the message is to be sent to.
	// Technology: PJSIP. Specifying a prefix of pjsip: will send the message as a SIP MESSAGE request
	// Technology: SIP. Specifying a prefix of sip: will send the message as a SIP MESSAGE request.
	// Technology: XMPP. Specifying a prefix of xmpp: will send the message as an XMPP chat message.
	// From - A From URI for the message if needed for the message technology being used to send this message.
	// Technology: PJSIP. The from parameter can be a configured endpoint or in the form of "display-name" <URI>.
	// Technology: SIP. The from parameter can be a configured peer name or in the form of "display-name" <URI>.
	// Technology: XMPP. Specifying a prefix of xmpp: will specify the account defined in xmpp.conf to send the message from. Note that this field is required for XMPP messages
	// Body - The message body text. This must not contain any newlines as that conflicts with the AMI protocol.
	// Base64Body - Text bodies requiring the use of newlines have to be base64 encoded in this field. Base64Body will be decoded before
	// being sent out. Base64Body takes precedence over Body.
	// Variable - Message variable to set, multiple Variable: headers are allowed. The header value is a comma separated list of name=value paris.
	AmiActionMessageSend = "MessageSend"
	// Record a call and mix the audio during the recording. Use of StopMixMonitor is required to guarantee the audio file is available for processing during
	// 	 dialplan execution.
	// This action records the audio on the current channel to the specified file
	// 	 MIXMONITOR_FILENAME - Will contain the filename used to record the mixed stream.
	// Syntax:
	// Action: MixMonitor
	// ActionID: <value>
	// Channel: <value>
	// File: <value>
	// options: <value>
	// Command: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned
	// Channel - Used to specify the channel to record.
	// File - Is the name of the file created in the monitor spool directory. Defaults to the same name as the channel (with slashes replaced
	//   with dashes). This argument is optional if you specify to record unidirectional audio with either the r(filename) or t(filename) options in the
	//   options field. If neither MIXMONITOR_FILENAME or this parameter is set, the mixed stream won't be recorded.
	// options - Options that apply to the MixMonitor in the same way as they would apply if invoked from the MixMonitor application. For a list
	// 	 of available options, see the documentation for the mixmonitor application.
	// Command - Will be executed when the recording is over. Any strings matching ^{X} will be unescaped to X. All variables will be evaluated
	// 	 at the time MixMonitor is called.
	AmiActionMixMonitor = "MixMonitor"
	// Mute / unMute a Mixmonitor recording.
	// This action may be used to mute a MixMonitor recording.
	// Syntax:
	// Action: MixMonitorMute
	// ActionID: <value>
	// Channel: <value>
	// Direction: <value>
	// State: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Channel - Used to specify the channel to mute.
	// Direction - Which part of the recording to mute: read, write or both (from channel, to channel or both channels)
	// State - Turn mute on or off : 1 to turn on, 0 to turn off.
	AmiActionMixMonitorMute = "MixMonitorMute"
	// Check if module is loaded.
	// Checks if Asterisk module is loaded. Will return Success/Failure. For success returns, the module revision number is included.
	// Syntax:
	// Action: ModuleCheck
	// ActionID: <value>
	// Module: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned
	// Module - Asterisk module name (not including extension).
	AmiActionModuleCheck = "ModuleCheck"
	// Module management.
	// Loads, unloads or reloads an Asterisk module in a running system
	// Syntax:
	// Action: ModuleLoad
	// ActionID: <value>
	// Module: <value>
	// LoadType: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Module - Asterisk module name (including .so extension) or subsystem identifier:
	/*
		cdr
		dnsmgr
		extconfig
		enum
		acl
		manager
		http
		logger
		features
		dsp
		udptl
		indications
		cel
		plc
	*/
	// LoadType - The operation to be done on module. Subsystem identifiers may only be reloaded.
	/*
		load
		unload
		reload
	*/
	// 	If no module is specified for a reload load type, all modules are reloaded.
	AmiActionModuleLoad = "ModuleLoad"
	// Monitor a channel.
	// This action may be used to record the audio on a specified channel
	// Syntax:
	// Action: Monitor
	// ActionID: <value>
	// Channel: <value>
	// File: <value>
	// Format: <value>
	// Mix: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Channel - Used to specify the channel to record.
	// File - Is the name of the file created in the monitor spool directory. Defaults to the same name as the channel (with slashes replaced
	// 	with dashes).
	// Format - Is the audio recording format. Defaults to wav.
	// Mix - Boolean parameter as to whether to mix the input and output channels together after the recording is finished.
	AmiActionMonitor = "Monitor"
	// Mute an audio stream.
	// Mute an incoming or outgoing audio stream on a channel.
	// Syntax:
	// Action: MuteAudio
	// ActionID: <value>
	// Channel: <value>
	// Direction: <value>
	// State: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned
	// Channel - The channel you want to mute.
	// Direction
	/*
		in - Set muting on inbound audio stream. (to the PBX)
		out - Set muting on outbound audio stream. (from the PBX)
		all - Set muting on inbound and outbound audio streams.
	*/
	// State
	/*
		on - Turn muting on.
		off - Turn muting off.
	*/
	AmiActionMuteAudio = "MuteAudio"
	// Delete selected mailboxes.
	// Syntax:
	// Action: MWIDelete
	// ActionID: <value>
	// Mailbox: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Mailbox - Mailbox ID in the form of / regex/ for all mailboxes matching the regular expression. Otherwise it is for a specific mailbox.
	AmiActionMWIDelete = "MWIDelete"
	// Get selected mailboxes with message counts.
	// Get a list of mailboxes with their message counts.
	// Syntax:
	// Action: MWIGet
	// ActionID: <value>
	// Mailbox: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Mailbox - Mailbox ID in the form of / regex/ for all mailboxes matching the regular expression. Otherwise it is for a specific mailbox
	AmiActionMWIGet = "MWIGet"
	// Update the mailbox message counts.
	// Syntax:
	// Action: MWIUpdate
	// ActionID: <value>
	// Mailbox: <value>
	// OldMessages: <value>
	// NewMessages: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Mailbox - Specific mailbox ID.
	// OldMessages - The number of old messages in the mailbox. Defaults to zero if missing.
	// NewMessages - The number of new messages in the mailbox. Defaults to zero if missing.
	AmiActionMWIUpdate = "MWIUpdate"
	// Wait for an event to occur.
	// This action will elicit a Success response. Whenever a manager event is queued. Once WaitEvent has been called on an HTTP manager session, events
	// will be generated and queued.
	// Syntax:
	// Action: WaitEvent
	// ActionID: <value>
	// Timeout: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Timeout - Maximum time (in seconds) to wait for events, -1 means forever.
	AmiActionWaitEvent = "WaitEvent"
	// Show the status of given voicemail user's info.
	// Retrieves the status of the given voicemail user
	// Syntax:
	// Action: VoicemailUserStatus
	// ActionID: <value>
	// Context: <value>
	// Mailbox: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned
	// Context - The context you want to check.
	// Mailbox - The mailbox you want to check.
	AmiActionVoicemailUserStatus = "VoicemailUserStatus"
	// List All Voicemail User Information.
	// Syntax:
	// Action: VoicemailUsersList
	// ActionID: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	AmiActionVoicemailUsersList = "VoicemailUsersList"
	// Tell Asterisk to poll mailboxes for a change
	// Normally, MWI indicators are only sent when Asterisk itself changes a mailbox. With external programs that modify the content of a mailbox from outside
	// the application, an option exists called pollmailboxes that will cause voicemail to continually scan all mailboxes on a system for changes. This can
	// cause a large amount of load on a system. This command allows external applications to signal when a particular mailbox has changed, thus permitting
	// external applications to modify mailboxes and MWI to work without introducing considerable CPU load.
	// If Context is not specified, all mailboxes on the system will be polled for changes. If Context is specified, but Mailbox is omitted, then all mailboxes within
	// Context will be polled. Otherwise, only a single mailbox will be polled for changes.
	// Syntax:
	// Action: VoicemailRefresh
	// ActionID: <value>
	// Context: <value>
	// Mailbox: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	AmiActionVoicemailRefresh = "VoicemailRefresh"
	// Send an arbitrary event.
	// Send an event to manager sessions.
	// Syntax:
	// Action: UserEvent
	// ActionID: <value>
	// UserEvent: <value>
	// Header1: <value>
	// HeaderN: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// UserEvent - Event string to send.
	// Header1 - Content1.
	// HeaderN - ContentN.
	AmiActionUserEvent = "UserEvent"
	// Update basic configuration.
	// This action will modify, create, or delete configuration elements in Asterisk configuration files.
	// Syntax:
	// Action: UpdateConfig
	// ActionID: <value>
	// SrcFilename: <value>
	// DstFilename: <value>
	// Reload: <value>
	// PreserveEffectiveContext: <value>
	// Action-000000: <value>
	// Cat-000000: <value>
	// Var-000000: <value>
	// Value-000000: <value>
	// Match-000000: <value>
	// Line-000000: <value>
	// Options-000000: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned
	// SrcFilename - Configuration filename to read (e.g. foo.conf).
	// DstFilename - Configuration filename to write (e.g. foo.conf)
	// Reload - Whether or not a reload should take place (or name of specific module)
	// PreserveEffectiveContext - Whether the effective category contents should be preserved on template change. Default is true (pre13.2 behavior)
	// Action-000000 - Action to take. 0's represent 6 digit number beginning with 000000
	/*
		NewCat
		RenameCat
		DelCat
		EmptyCat
		Update
		Delete
		Append
		Insert
	*/
	// Cat-000000 - Category to operate on. 0's represent 6 digit number beginning with 000000.
	// Var-000000 - Variable to work on. 0's represent 6 digit number beginning with 000000.
	// Value-000000 - Value to work on. 0's represent 6 digit number beginning with 000000.
	// Match-000000 - Extra match required to match line. 0's represent 6 digit number beginning with 000000.
	// Line-000000 - Line in category to operate on (used with delete and insert actions). 0's represent 6 digit number beginning with 000000.
	// Options-000000 - A comma separated list of action-specific options.
	// 	NewCat - One or more of the following...
	/*
		allowdups - Allow duplicate category names.
		template - This category is a template.
		inherit="template,..." - Templates from which to inherit.
	*/
	// The following actions share the same options...
	/*
		RenameCat
		DelCat
		EmptyCat
		Update
		Delete
		Append
		Insert -
		catfilter="<expression>,..." -
				A comma separated list of name_regex=value_regex expressions which will cause only categories whose variables
				match all expressions to be considered. The special variable name TEMPLATES can be used to control whether
				templates are included. Passing include as the value will include templates along with normal categories. Passing res
				trict as the value will restrict the operation to ONLY templates. Not specifying a TEMPLATES expression results in the
				default behavior which is to not include templates.
				catfilter is most useful when a file contains multiple categories with the same name and you wish to operate on specific
				ones instead of all of them.
				0's represent 6 digit number beginning with 000000.
	*/
	AmiActionUpdateConfig = "UpdateConfig"
	// Unpause monitoring of a channel.
	// This action may be used to re-enable recording of a channel after calling PauseMonitor.
	// Syntax:
	// Action: UnpauseMonitor
	// ActionID: <value>
	// Channel: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Channel - Used to specify the channel to record.
	AmiActionUnpauseMonitor = "UnpauseMonitor"
	// Stop monitoring a channel
	// This action may be used to end a previously started 'Monitor' action.
	// Syntax:
	// Action: StopMonitor
	// ActionID: <value>
	// Channel: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Channel - The name of the channel monitored.
	AmiActionStopMonitor = "StopMonitor"
	// Stop recording a call through MixMonitor, and free the recording's file handle.
	// This action stops the audio recording that was started with the MixMonitor action on the current channel.
	// Syntax:
	// Action: StopMixMonitor
	// ActionID: <value>
	// Channel: <value>
	// [MixMonitorID:] <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Channel - The name of the channel monitored.
	// MixMonitorID - If a valid ID is provided, then this command will stop only that specific MixMonitor.
	AmiActionStopMixMonitor = "StopMixMonitor"
	// List channel status.
	// Will return the status information of each channel along with the value for the specified channel variables.
	// Syntax:
	// Action: Status
	// ActionID: <value>
	// [Channel:] <value>
	// Variables: <value>
	// AllVariables: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Channel - The name of the channel to query for status.
	// Variables - Comma , separated list of variable to include.
	// AllVariables - If set to "true", the Status event will include all channel variables for the requested channel(s). True or false
	AmiActionStatus = "Status"
	// Mark an object in a sorcery memory cache as stale.
	// Marks an object as stale within a sorcery memory cache.
	// Syntax:
	// Action: SorceryMemoryCacheStaleObject
	// ActionID: <value>
	// Cache: <value>
	// Object: <value>
	// [Reload:] <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Cache - The name of the cache to mark the object as stale in.
	// Object - The name of the object to mark as stale.
	// Reload - If true, then immediately reload the object from the backend cache instead of waiting for the next retrieval
	AmiActionSorceryMemoryCacheStaleObject = "SorceryMemoryCacheStaleObject"
	// Marks ALL objects in a sorcery memory cache as stale.
	// Marks ALL objects in a sorcery memory cache as stale.
	// Syntax:
	// Action: SorceryMemoryCacheStale
	// ActionID: <value>
	// Cache: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Cache - The name of the cache to mark all object as stale in.
	AmiActionSorceryMemoryCacheStale = "SorceryMemoryCacheStale"
	// Expire all objects from a memory cache and populate it with all objects from the backend.
	// Expires all objects from a memory cache and populate it with all objects from the backend
	// Syntax:
	// Action: SorceryMemoryCachePopulate
	// ActionID: <value>
	// Cache: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Cache - The name of the cache to populate.
	AmiActionSorceryMemoryCachePopulate = "SorceryMemoryCachePopulate"
	// Expire (remove) an object from a sorcery memory cache.
	// Expires (removes) an object from a sorcery memory cache.
	// Syntax:
	// Action: SorceryMemoryCacheExpireObject
	// ActionID: <value>
	// Cache: <value>
	// Object: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned
	// Cache - The name of the cache to expire the object from.
	// Object - The name of the object to expire.
	AmiActionSorceryMemoryCacheExpireObject = "SorceryMemoryCacheExpireObject"
	// Expire (remove) ALL objects from a sorcery memory cache.
	// Expires (removes) ALL objects from a sorcery memory cache.
	// Syntax:
	// Action: SorceryMemoryCacheExpire
	// ActionID: <value>
	// Cache: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Cache - The name of the cache to expire all objects from.
	AmiActionSorceryMemoryCacheExpire = "SorceryMemoryCacheExpire"
	// Show SKINNY line (text format).
	// Show one SKINNY line with details on current status.
	// Syntax:
	// Action: SKINNYshowline
	// ActionID: <value>
	// Line: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Line - The line name you want to check.
	AmiActionSKINNYShowLine = "SKINNYshowline"
	// Show SKINNY device (text format).
	// Show one SKINNY device with details on current status.
	// Syntax:
	// Action: SKINNYshowdevice
	// ActionID: <value>
	// Device: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Device - The device name you want to check.
	AmiActionSKINNYShowDevice = "SKINNYshowdevice"
	// List SKINNY lines (text format).
	// Lists Skinny lines in text format with details on current status. Linelist will follow as separate events, followed by a final event called LinelistComplete.
	// Syntax:
	// Action: SKINNYlines
	// ActionID: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	AmiActionSKINNYlines = "SKINNYlines"
	// List SKINNY devices (text format).
	// Lists Skinny devices in text format with details on current status. Devicelist will follow as separate events, followed by a final event called DevicelistComplete.
	// Syntax:
	// Action: SKINNYdevices
	// ActionID: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	AmiActionSKINNYdevices = "SKINNYdevices"
	// Show SIP registrations (text format).
	// Lists all registration requests and status. Registrations will follow as separate events followed by a final event called RegistrationsComplete.
	// Syntax:
	// Action: SIPshowregistry
	// ActionID: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	AmiActionSIPShowRegistry = "SIPshowregistry"
	// show SIP peer (text format).
	// Show one SIP peer with details on current status.
	// Syntax:
	// Action: SIPshowpeer
	// ActionID: <value>
	// Peer: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Peer - The peer name you want to check.
	AmiActionSIPShowPeer = "SIPshowpeer"
	// Qualify SIP peers.
	// Syntax:
	// Action: SIPqualifypeer
	// ActionID: <value>
	// Peer: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Peer - The peer name you want to qualify.
	AmiActionSIPQualifyPeer = "SIPqualifypeer"
	// Show the status of one or all of the sip peers.
	// Retrieves the status of one or all of the sip peers. If no peer name is specified, status for all of the sip peers will be retrieved
	// Syntax:
	// Action: SIPpeerstatus
	// ActionID: <value>
	// [Peer:] <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Peer - The peer name you want to check.
	AmiActionSIPPeerStatus = "SIPpeerstatus"
	// List SIP peers (text format).
	// Lists SIP peers in text format with details on current status. Peerlist will follow as separate events, followed by a final event called PeerlistComplete.
	// Syntax:
	// Action: SIPpeers
	// ActionID: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	AmiActionSIPPeers = "SIPpeers"
	// Send a SIP notify.
	// Sends a SIP Notify event.
	// All parameters for this event must be specified in the body of this request via multiple Variable: name=value sequences.
	// Syntax:
	// Action: SIPnotify
	// ActionID: <value>
	// Channel: <value>
	// Variable: <value>
	// [Call-ID:] <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Channel - Peer to receive the notify.
	// Variable - At least one variable pair must be specified. name=value
	// Call-ID - When specified, SIP notify will be sent as a part of an existing dialog.
	AmiActionSIPNotify = "SIPnotify"
	// Show dialplan contexts and extensions
	// Show dialplan contexts and extensions. Be aware that showing the full dialplan may take a lot of capacity.
	// Syntax:
	// Action: ShowDialPlan
	// ActionID: <value>
	// Extension: <value>
	// Context: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Extension - Show a specific extension.
	// Context - Show a specific context.
	AmiActionShowDialPlan = "ShowDialPlan"
	// Sets a channel variable or function value.
	// This command can be used to set the value of channel variables or dialplan functions.
	// Syntax:
	// Action: Setvar
	// ActionID: <value>
	// Channel: <value>
	// Variable: <value>
	// Value: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Channel - Channel to set variable for.
	// Variable - Variable name, function or expression.
	// Value - Variable or function value.
	AmiActionSetVar = "Setvar"
	// Send text message to channel.
	// Sends A Text Message to a channel while in a call.
	// Syntax:
	// Action: SendText
	// ActionID: <value>
	// Channel: <value>
	// Message: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Channel - Channel to send message to.
	// Message - Message to send.
	AmiActionSendText = "SendText"
	// Send a reload event.
	// Syntax:
	// Action: Reload
	// ActionID: <value>
	// Module: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Module - Name of the module to reload.
	AmiActionReload = "Reload"
	// Redirect (transfer) a call.
	// Syntax:
	/*
		Action: Redirect
		ActionID: <value>
		Channel: <value>
		ExtraChannel: <value>
		Exten: <value>
		ExtraExten: <value>
		Context: <value>
		ExtraContext: <value>
		Priority: <value>
		ExtraPriority: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Channel - Channel to redirect.
		ExtraChannel - Second call leg to transfer (optional).
		Exten - Extension to transfer to.
		ExtraExten - Extension to transfer extra channel to (optional).
		Context - Context to transfer to.
		ExtraContext - Context to transfer extra channel to (optional).
		Priority - Priority to transfer to.
		ExtraPriority - Priority to transfer extra channel to (optional).
	*/
	AmiActionRedirect = "Redirect"
	// Show queue summary.
	// Request the manager to send a QueueSummary event.
	// Syntax:
	// Action: QueueSummary
	// ActionID: <value>
	// Queue: <value>
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Queue - Queue for which the summary is requested.
	AmiActionQueueSummary = "QueueSummary"
	// Show queue status.
	// Check the status of one or more queues.
	// Syntax:
	/*
		Action: QueueStatus
		ActionID: <value>
		Queue: <value>
		Member: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Queue - Limit the response to the status of the specified queue.
		Member - Limit the response to the status of the specified member.
	*/
	AmiActionQueueStatus = "QueueStatus"
	// Queue Rules.
	// List queue rules defined in queuerules.conf
	// Syntax:
	/*
		Action: QueueRule
		ActionID: <value>
		Rule: <value>
	*/
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Rule - The name of the rule in queuerules.conf whose contents to list.
	AmiActionQueueRule = "QueueRule"
	// Reset queue statistics.
	// Reset the statistics for a queue.
	// Syntax:
	/*
		Action: QueueReset
		ActionID: <value>
		Queue: <value>
	*/
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	// Queue - The name of the queue on which to reset statistics.
	AmiActionQueueReset = "QueueReset"
	// Remove interface from queue.
	// Syntax:
	/*
		Action: QueueRemove
		ActionID: <value>
		Queue: <value>
		Interface: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Queue - The name of the queue to take action on.
		Interface - The interface (tech/name) to remove from queue
	*/
	AmiActionQueueRemove = "QueueRemove"
	// Reload a queue, queues, or any sub-section of a queue or queues.
	// Syntax:
	/*
		Action: QueueReload
		ActionID: <value>
		Queue: <value>
		Members: <value>
		Rules: <value>
		Parameters: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Queue - The name of the queue to take action on. If no queue name is specified, then all queues are affected.
		Members - Whether to reload the queue's members.
				yes
				no
		Rules - Whether to reload queuerules.conf
				yes
				no
		Parameters - Whether to reload the other queue options.
				yes
				no
	*/
	AmiActionQueueReload = "QueueReload"
	// Set the penalty for a queue member.
	// Change the penalty of a queue member
	// Syntax:
	/*
		Action: QueuePenalty
		ActionID: <value>
		Interface: <value>
		Penalty: <value>
		Queue: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Interface - The interface (tech/name) of the member whose penalty to change.
		Penalty - The new penalty (number) for the member. Must be nonnegative.
		Queue - If specified, only set the penalty for the member of this queue. Otherwise, set the penalty for the member in all queues to which
		the member belongs.
	*/
	AmiActionQueuePenalty = "QueuePenalty"
	// Makes a queue member temporarily unavailable.
	// Pause or unpause a member in a queue.
	// Syntax:
	/*
		Action: QueuePause
		ActionID: <value>
		Interface: <value>
		Paused: <value>
		Queue: <value>
		Reason: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Interface - The name of the interface (tech/name) to pause or unpause.
		Paused - Pause or unpause the interface. Set to 'true' to pause the member or 'false' to unpause.
		Queue - The name of the queue in which to pause or unpause this member. If not specified, the member will be paused or unpaused in
		all the queues it is a member of.
		Reason - Text description, returned in the event QueueMemberPaused.
	*/
	AmiActionQueuePause = "QueuePause"
	// Set the ringinuse value for a queue member.
	// Syntax:
	/*
		Action: QueueMemberRingInUse
		ActionID: <value>
		Interface: <value>
		RingInUse: <value>
		Queue: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Interface
		RingInUse
		Queue
	*/
	AmiActionQueueMemberRingInUse = "QueueMemberRingInUse"
	// Adds custom entry in queue_log
	// Syntax:
	/*
		Action: QueueLog
		ActionID: <value>
		Queue: <value>
		Event: <value>
		Uniqueid: <value>
		Interface: <value>
		Message: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Queue
		Event
		Uniqueid
		Interface
		Message
	*/
	AmiActionQueueLog = "QueueLog"
	// Change priority of a caller on queue.
	// Syntax:
	/*
		Action: QueueChangePriorityCaller
		ActionID: <value>
		Queue: <value>
		Caller: <value>
		Priority: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Queue - The name of the queue to take action on.
		Caller - The caller (channel) to change priority on queue.
		Priority - Priority value for change for caller on queue.
	*/
	AmiActionQueueChangePriorityCaller = "QueueChangePriorityCaller"
	// Add interface to queue.
	// Syntax:
	/*
		Action: QueueAdd
		ActionID: <value>
		Queue: <value>
		Interface: <value>
		Penalty: <value>
		Paused: <value>
		MemberName: <value>
		StateInterface: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Queue - Queue's name.
		Interface - The name of the interface (tech/name) to add to the queue.
		Penalty - A penalty (number) to apply to this member. Asterisk will distribute calls to members with higher penalties only after
		attempting to distribute calls to those with lower penalty.
		Paused - To pause or not the member initially (true/false or 1/0).
		MemberName - Text alias for the interface.
		StateInterface
	*/
	AmiActionQueueAdd = "QueueAdd"
	// Show status of PRI spans.
	// Similar to the CLI command "pri show spans".
	// Syntax:
	/*
		Action: PRIShowSpans
		ActionID: <value>
		Span: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Span - Specify the specific span to show. Show all spans if zero or not present.
	*/
	AmiActionPRIShowSpans = "PRIShowSpans"
	// Set PRI debug levels for a span
	// Equivalent to the CLI command "pri set debug <level> span <span>".
	// Syntax:
	/*
		Action: PRIDebugSet
		ActionID: <value>
		Span: <value>
		Level: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Span - Which span to affect.
		Level - What debug level to set. May be a numerical value or a text value from the list below
			off
			on
			hex
			intense
	*/
	AmiActionPRIDebugSet = "PRIDebugSet"
	// Disables file output for PRI debug messages
	// Syntax:
	/*
		Action: PRIDebugFileUnset
		ActionID: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
	*/
	AmiActionPRIDebugFileUnset = "PRIDebugFileUnset"
	// Set the file used for PRI debug message output
	// Equivalent to the CLI command "pri set debug file <output-file>"
	// Syntax:
	/*
		Action: PRIDebugFileSet
		ActionID: <value>
		File: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		File - Path of file to write debug output.
	*/
	AmiActionPRIDebugFileSet = "PRIDebugFileSet"
	// List the current known presence states.
	// This will list out all known presence states in a sequence of PresenceStateChange events. When finished, a PresenceStateListComplete event will be emitted.
	// Syntax:
	/*
		Action: PresenceStateList
		ActionID: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
	*/
	AmiActionPresenceStateList = "PresenceStateList"
	// Check Presence State
	// Report the presence state for the given presence provider.
	// Will return a Presence State message. The response will include the presence state and, if set, a presence subtype and custom message.
	// Syntax:
	/*
		Action: PresenceState
		ActionID: <value>
		Provider: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Provider - Presence Provider to check the state of
	*/
	AmiActionPresenceState = "PresenceState"
	// Play DTMF signal on a specific channel.
	// Syntax:
	/*
		Action: PlayDTMF
		ActionID: <value>
		Channel: <value>
		Digit: <value>
		[Duration:] <value>
		[Receive:] <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Channel - Channel name to send digit to.
		Digit - The DTMF digit to play.
		Duration - The duration, in milliseconds, of the digit to be played.
		Receive - Emulate receiving DTMF on this channel instead of sending it out.
	*/
	AmiActionPlayDtmf = "PlayDTMF"
	// Unregister an outbound registration.
	// Unregister the specified (or all) outbound registration(s) and stops future registration attempts. Call PJSIPRegister to start registration and schedule
	// re-registrations according to configuration.
	// Syntax:
	/*
		Action: PJSIPUnregister
		ActionID: <value>
		Registration: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Registration - The outbound registration to unregister or '*all' to unregister them all.
	*/
	AmiActionPJSIPUnregister = "PJSIPUnregister"
	// Lists subscriptions.
	// Provides a listing of all outbound subscriptions. An event OutboundSubscriptionDetail is issued for each subscription object. Once all detail events
	// are completed an OutboundSubscriptionDetailComplete event is issued.
	// Syntax:
	// Action: PJSIPShowSubscriptionsOutbound
	AmiActionPJSIPShowSubscriptionsOutbound = "PJSIPShowSubscriptionsOutbound"
	// Lists subscriptions
	// Provides a listing of all inbound subscriptions. An event InboundSubscriptionDetail is issued for each subscription object. Once all detail events are
	// completed an InboundSubscriptionDetailComplete event is issued.
	// Syntax:
	// Action: PJSIPShowSubscriptionsInbound
	AmiActionPJSIPShowSubscriptionsInbound = "PJSIPShowSubscriptionsInbound"
	// Displays settings for configured resource lists.
	// Provides a listing of all resource lists. An event ResourceListDetail is issued for each resource list object. Once all detail events are completed a ResourceListDetailComplete event is issued.
	// Syntax:
	// Action: PJSIPShowResourceLists
	AmiActionPJSIPShowResourceLists = "PJSIPShowResourceLists"
	// Lists PJSIP outbound registrations.
	// In response OutboundRegistrationDetail events showing configuration and status information are raised for each outbound registration object.
	// AuthDetail events are raised for each associated auth object as well. Once all events are completed an OutboundRegistrationDetailComplete is issued
	// Syntax:
	// Action: PJSIPShowRegistrationsOutbound
	AmiActionPJSIPShowRegistrationsOutbound = "PJSIPShowRegistrationsOutbound"
	// Lists PJSIP inbound registrations.
	// In response, InboundRegistrationDetail events showing configuration and status information are raised for all contacts, static or dynamic. Once all
	// events are completed an InboundRegistrationDetailComplete is issued.
	// Syntax:
	// Action: PJSIPShowRegistrationsInbound
	AmiActionPJSIPShowRegistrationsInbound = "PJSIPShowRegistrationsInbound"
	// Lists ContactStatuses for PJSIP inbound registrations.
	// In response, ContactStatusDetail events showing status information are raised for each inbound registration (dynamic contact) object. Once all events
	// are completed a ContactStatusDetailComplete event is issued.
	// Syntax:
	// Action: PJSIPShowRegistrationInboundContactStatuses
	AmiActionPJSIPShowRegistrationInboundContactStatuses = "PJSIPShowRegistrationInboundContactStatuses"
	// Lists PJSIP endpoints.
	// Provides a listing of all endpoints. For each endpoint an EndpointList event is raised that contains relevant attributes and status information. Once all
	// endpoints have been listed an EndpointListComplete event is issued.
	// Syntax:
	// Action: PJSIPShowEndpoints
	AmiActionPJSIPShowEndpoints = "PJSIPShowEndpoints"
	// Detail listing of an endpoint and its objects.
	// Provides a detailed listing of options for a given endpoint. Events are issued showing the configuration and status of the endpoint and associated objects.
	// These events include EndpointDetail, AorDetail, AuthDetail, TransportDetail, and IdentifyDetail. Some events may be listed multiple
	// times if multiple objects are associated (for instance AoRs). Once all detail events have been raised a final EndpointDetailComplete event is issued
	// Syntax:
	/*
		Action: PJSIPShowEndpoint
		ActionID: <value>
		Endpoint: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Endpoint - The endpoint to list.
	*/
	AmiActionPJSIPShowEndpoint = "PJSIPShowEndpoint"
	// Lists PJSIP Contacts.
	// Provides a listing of all Contacts. For each Contact a ContactList event is raised that contains relevant attributes and status information. Once all
	// contacts have been listed a ContactListComplete event is issued
	// Syntax:
	/*
		Action: PJSIPShowContacts
	*/
	AmiActionPJSIPShowContacts = "PJSIPShowContacts"
	// Lists PJSIP Auths.
	// Provides a listing of all Auths. For each Auth an AuthList event is raised that contains relevant attributes and status information. Once all auths have
	// been listed an AuthListComplete event is issued.
	// Syntax:
	// Action: PJSIPShowAuths
	AmiActionPJSIPShowAuths = "PJSIPShowAuths"
	// Lists PJSIP AORs.
	// Provides a listing of all AORs. For each AOR an AorList event is raised that contains relevant attributes and status information. Once all aors have been
	// listed an AorListComplete event is issued.
	// Syntax:
	// Action: PJSIPShowAors
	AmiActionPJSIPShowAors = "PJSIPShowAors"
	// Register an outbound registration.
	// Unregisters the specified (or all) outbound registration(s) then starts registration and schedules re-registrations according to configuration
	// Syntax:
	/*
		Action: PJSIPRegister
		ActionID: <value>
		Registration: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Registration - The outbound registration to register or '*all' to register them all.
	*/
	AmiActionPJSIPRegister = "PJSIPRegister"
	// Qualify a chan_pjsip endpoint.
	// Syntax:
	/*
		Action: PJSIPQualify
		ActionID: <value>
		Endpoint: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Endpoint - The endpoint you want to qualify.
	*/
	AmiActionPJSIPQualify = "PJSIPQualify"
	// Send a NOTIFY to either an endpoint, an arbitrary URI, or inside a SIP dialog.
	// Sends a NOTIFY to an endpoint, an arbitrary URI, or inside a SIP dialog.
	// All parameters for this event must be specified in the body of this requestvia multiple Variable: name=value sequences.
	// Syntax:
	/*
		Action: PJSIPNotify
		ActionID: <value>
		[Endpoint:] <value>
		[URI:] <value>
		[channel:] <value>
		Variable: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Endpoint - The endpoint to which to send the NOTIFY.
		URI - Arbitrary URI to which to send the NOTIFY.
		channel - Channel name to send the NOTIFY. Must be a PJSIP channel.
		Variable - Appends variables as headers/content to the NOTIFY. If the variable is named Content, then the value will compose the
		body of the message if another variable sets Content-Type. name=value
	*/
	AmiActionPJSIPNotify = "PJSIPNotify"
	// Keepalive command.
	// A 'Ping' action will elicit a 'Pong' response. Used to keep the manager connection open.
	// Syntax:
	/*
		Action: Ping
		ActionID: <value>
	*/
	AmiActionPing = "Ping"
	// Pause monitoring of a channel.
	// This action may be used to temporarily stop the recording of a channel.
	// Syntax:
	/*
		Action: PauseMonitor
		ActionID: <value>
		Channel: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Channel - Used to specify the channel to record.
	*/
	AmiActionPauseMonitor = "PauseMonitor"
	// Get a list of parking lots
	// List all parking lots as a series of AMI events
	// Syntax:
	/*
		Action: Parkinglots
		ActionID: <value>
	*/
	// Args:
	// ActionID - ActionID for this transaction. Will be returned.
	AmiActionParkingLots = "Parkinglots"
	// List parked calls.
	// Syntax:
	/*
		Action: ParkedCalls
		ActionID: <value>
		ParkingLot: <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		ParkingLot - If specified, only show parked calls from the parking lot with this name.
	*/
	AmiActionParkedCalls = "ParkedCalls"
	// Park a channel.
	// Park an arbitrary channel with optional arguments for specifying the parking lot used, how long the channel should remain parked, and what dial string to
	// use as the parker if the call times out.
	// Syntax:
	/*
		Action: Park
		ActionID: <value>
		Channel: <value>
		[TimeoutChannel:] <value>
		[AnnounceChannel:] <value>
		[Timeout:] <value>
		[Parkinglot:] <value>
	*/
	// Args:
	/*
		ActionID - ActionID for this transaction. Will be returned.
		Channel - Channel name to park.
		TimeoutChannel - Channel name to use when constructing the dial string that will be dialed if the parked channel times out. If Timeout
			Channel is in a two party bridge with Channel, then TimeoutChannel will receive an announcement and be treated as having
			parked Channel in the same manner as the Park Call DTMF feature.
		AnnounceChannel - If specified, then this channel will receive an announcement when Channel is parked if AnnounceChannel is in a
			state where it can receive announcements (AnnounceChannel must be bridged). AnnounceChannel has no bearing on the actual state
			of the parked call.
		Timeout - Overrides the timeout of the parking lot for this park action. Specified in milliseconds, but will be converted to seconds. Use a
			value of 0 to disable the timeout.
		Parkinglot - The parking lot to use when parking the channel
	*/
	AmiActionPark     = "Park"
	AmiActionDataGet  = "DataGet"
	AmiActionKSendSMS = "KSendSMS"
)
