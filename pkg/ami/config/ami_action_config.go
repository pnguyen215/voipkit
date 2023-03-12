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
	//
)
