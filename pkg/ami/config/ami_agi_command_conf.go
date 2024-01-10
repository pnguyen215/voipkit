package config

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
