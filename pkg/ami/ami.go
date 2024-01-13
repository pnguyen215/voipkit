package ami

import (
	"bufio"
	"context"
	"net"
	"net/textproto"
	"strings"
	"time"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

// Action sends an AMI action message to the Asterisk server.
// If the action message does not have an ActionID, it adds one automatically.
// The method returns true if the action message is successfully sent, otherwise false.
// In case of an error during the write operation, it emits an error event and returns false.
//
// Parameters:
//   - message: Pointer to an AMIMessage representing the action message to be sent.
//
// Example:
//
//	actionMessage := NewAMIMessage("Command")
//	actionMessage.AddParameter("Action", "Status")
//	actionMessage.AddParameter("Command", "core show channels")
//	amiClient.Action(actionMessage)
//
// Note: This method is typically used to send action requests to the Asterisk Manager Interface (AMI)
// for tasks such as retrieving status information or executing commands.
func (c *AMI) Action(message *AMIMessage) bool {
	if message.GetActionId() == "" {
		message.AddActionId()
	}
	err := c.write(message.Bytes())
	if err != nil {
		c.EmitError(err)
		return false
	}
	return true
}

// AllEvents subscribes to any AMI message received from the Asterisk server.
// It returns a send-only channel of pointers to AMIMessage or nil if not subscribed.
//
// Example:
//
//	allEventsChannel := amiClient.AllEvents()
//	go func() {
//	    for message := range allEventsChannel {
//	        // Handle AMI messages received from the Asterisk server
//	        fmt.Println("Received AMI Message:", message)
//	    }
//	}()
//
// Note: This method allows you to subscribe to all Asterisk Manager Interface (AMI) messages,
// providing a channel through which you can receive and handle messages as they are received from the server.
// It is useful for capturing and reacting to various events happening within the Asterisk communication system.
func (c *AMI) AllEvents() <-chan *AMIMessage {
	return c.Subs.Subscribe(config.AmiPubSubKeyRef)
}

// OnEvent subscribes by event name (case insensitive) and
// returns send-only channel or nil
func (c *AMI) OnEvent(name string) <-chan *AMIMessage {
	return c.Subs.Subscribe(name)
}

// OnEvents subscribes by events name (case insensitive) and
// return send-only channel or nil
func (c *AMI) OnEvents(keys ...string) <-chan *AMIMessage {
	return c.Subs.Subscribes(keys...)
}

// EmitError sends an error to the error channel (c.Err) in a non-blocking manner.
// If the error or the error channel is nil, it returns immediately.
//
// Example:
//
//	amiClient.EmitError(errors.New("An error occurred"))
//
// Note: This method is designed to be used internally to safely send errors to the error channel
// without blocking the main execution flow.
func (c *AMI) EmitError(err error) {
	go func(err error) {
		c.Mutex.Lock()
		defer c.Mutex.Unlock()
		if err == nil || c.Err == nil {
			return
		}
		c.Err <- err
	}(err)
}

// Error returns a read-only channel for errors, signaling that the AMI client encountered an error,
// and the client might need to be restarted.
//
// Example:
//
//	errorChannel := amiClient.Error()
//	select {
//	case err := <-errorChannel:
//	  log.Printf("Error received: %v", err)
//	  // Handle the error, possibly restarting the AMI client.
//	}
//
// Note: This method provides a non-blocking way to receive errors from the AMI client.
// It returns the error channel, allowing clients to listen for errors and take appropriate actions.
func (c *AMI) Error() <-chan error {
	c.Mutex.RLock()
	defer c.Mutex.RUnlock()
	return c.Err
}

// Close closes the AMI client, terminating its connection and cleaning up resources.
// It also destroys all subscriptions to events and action responses.
//
// Example:
//
//	amiClient.Close()
//
// Note: Closing the AMI client is crucial to ensure proper cleanup of resources.
// It terminates the connection, cancels the context, destroys event subscriptions, and closes the error channel.
// Once closed, the AMI client should not be used further, and a new instance may be created if needed.
func (c *AMI) Close() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Cancel()
	c.Subs.Destroy()
	c.Conn.Close()
	close(c.Err)
	c.Err = nil
}

// publish sends the provided AMI message to all subscribers based on event type and general subscriptions.
// It utilizes the AMI Pub-Sub mechanism for broadcasting messages to interested parties.
// This method is typically called after receiving an AMI message to notify subscribers.
//
// Example:
//
//	amiClient.publish(amiMessage)
//
// Note: The AMI Pub-Sub mechanism allows different parts of the application to subscribe to specific events or receive all events.
// When a new AMI message is received, this method broadcasts it to relevant subscribers.
func (c *AMI) publish(message *AMIMessage) {
	if message != nil {
		c.Subs.Publish(message)
	}
}

// apply sends an action to Asterisk and waits for the acknowledgment prompt.
// It sets a timeout for the operation to prevent blocking for an extended period.
// This method is used to execute actions in the Asterisk Manager Interface (AMI).
//
// Parameters:
//   - _ctx: The parent context in which the action is executed.
//   - timeout: The maximum duration allowed for the operation, including sending the action and waiting for the acknowledgment.
//
// Returns:
//   - An error if the operation encounters any issues or if it times out.
//
// Example:
//
//	err := amiClient.apply(context.Background(), 5*time.Second)
//
// Note: The method sends the action to Asterisk and waits for a response. It validates the acknowledgment prompt
// to ensure that the operation is successful. If the prompt is not received or does not match the expected format,
// an error is returned. The provided context and timeout parameters help control the duration of the operation.
func (c *AMI) apply(_ctx context.Context, timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(_ctx, timeout)
	defer cancel()
	reader := func() (chan string, chan error) {
		prompt := make(chan string)
		fail := make(chan error)
		go func() {
			defer close(prompt)
			defer close(fail)
			line, err := c.Reader.ReadLine()
			if err != nil {
				fail <- err
				return
			}
			prompt <- line
		}()
		return prompt, fail
	}
	prompt, fail := reader()
	select {
	case <-ctx.Done():
		return ErrorAsteriskConnTimeout
	case err := <-fail:
		return ErrorAsteriskNetwork.ErrorWrap(err.Error())
	case promptLine := <-prompt:
		if !strings.HasPrefix(promptLine, config.AmiCallManagerKey) {
			return ErrorAsteriskInvalidPrompt.ErrorWrap(promptLine)
		}
	}
	return nil
}

// read reads a message from the Asterisk Manager Interface (AMI) connection.
// It parses the MIME headers and creates an AMIMessage from the received data.
//
// Returns:
//   - An AMIMessage containing the parsed MIME headers.
//   - An error if there is an issue reading or parsing the message.
//
// Example:
//
//	message, err := amiClient.read()
//
// Note: The read method is used to read messages from the AMI connection. It parses MIME headers
// and creates an AMIMessage struct to represent the received data. If an error occurs during the read
// operation, it is returned as an error.
func (c *AMI) read() (*AMIMessage, error) {
	headers, err := c.Reader.ReadMIMEHeader()
	if err != nil {
		if err.Error() == ErrorEOF || err.Error() == ErrorIO {
			return nil, ErrorAsteriskNetwork
		}
		return nil, err
	}
	message := ofMessage(headers)
	c.Raw = message
	return message, nil
}

// write sends the provided bytes to the Asterisk Manager Interface (AMI) connection.
//
// Parameters:
//   - bytes: The byte slice to be sent to the AMI connection.
//
// Returns:
//   - An error if there is an issue writing the bytes to the connection.
//
// Example:
//
//	err := amiClient.write([]byte("Action: Login\nUsername: admin\nSecret: mySecret\n\n"))
//
// Note: The write method is used to send bytes to the AMI connection. It acquires a lock on the
// connection to ensure thread safety and writes the bytes using the Writer. Any error during the
// write operation is returned as an error.
func (c *AMI) write(bytes []byte) error {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	_, err := c.Writer.Write(bytes)
	if err != nil {
		return err
	}
	err = c.Writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

// release initiates the message handling loop for incoming data from the Asterisk Manager Interface (AMI) connection.
// It continuously reads messages from the connection, processes them, and publishes them to the appropriate subscribers.
//
// Parameters:
//   - ctx: The context.Context used for cancellation and error handling.
//
// Example:
//
//	amiClient.release(ctx)
//
// Note: The release method is responsible for managing the continuous reading of messages from the AMI connection.
// It initializes a new PubSubQueue for event subscriptions and a channel for errors. The function runs in a goroutine,
// continuously reading messages from the connection, publishing them to subscribers, and handling errors. It terminates
// when the provided context is canceled or an error occurs during the reading process.
func (c *AMI) release(ctx context.Context) {
	c.Subs = NewPubSubQueue()
	c.Err = make(chan error)
	go func() {
		defer c.Subs.Disabled()
		for {
			select {
			case <-ctx.Done():
				c.EmitError(ctx.Err())
				return
			default:
				message, err := c.read()
				if ctx.Err() != nil {
					c.EmitError(ctx.Err())
					return
				}
				if err != nil {
					if err == ErrorAsteriskNetwork {
						c.EmitError(err)
						return
					}
				}
				c.Raw = message
				c.publish(message)
			}
		}
	}()
}

// authenticate performs the authentication process with the Asterisk Manager Interface (AMI) server using the provided
// AmiClient configuration. It sends an authentication request, reads the response, and verifies the success of the authentication.
//
// Parameters:
//   - _ctx: The context.Context used for cancellation and timeout.
//   - request: The AmiClient configuration containing authentication details.
//
// Returns:
//   - error: An error indicating any issues during the authentication process.
//
// Example:
//
//	err := amiClient.authenticate(ctx, AmiClient{
//	  username: "admin",
//	  password: "secret*password",
//	  timeout:  time.Second * 5,
//	})
//
// Note: The authenticate method sends an authentication request to the AMI server, reads the response, and checks if the
// authentication was successful. It uses the provided context for cancellation and timeout handling. If the context is canceled,
// it returns an error indicating a connection timeout. If there is an issue reading the response or the authentication fails,
// it returns an appropriate error. If the authentication is successful, it updates the Raw field of the AMI client with the
// authentication response.
func (c *AMI) authenticate(_ctx context.Context, request AmiClient) error {
	ctx, cancel := context.WithTimeout(_ctx, request.timeout)
	defer cancel()
	message := Authenticate(request.username, request.password)
	reader := func() (chan *AMIMessage, chan error) {
		chMessage := make(chan *AMIMessage)
		fail := make(chan error)
		go func() {
			defer close(chMessage)
			defer close(fail)
			_message, err := c.read()
			if err != nil {
				fail <- err
				return
			}
			chMessage <- _message
		}()
		return chMessage, fail
	}
	message_reader, fail := reader()
	if err := c.write(message.Bytes()); err != nil {
		return err
	}
	select {
	case <-ctx.Done():
		return ErrorAsteriskConnTimeout.ErrorWrap(ErrorAuthenticatedUnsuccessfully)
	case err := <-fail:
		return ErrorAsteriskNetwork.ErrorWrap(err.Error())
	case msg := <-message_reader:
		if !msg.IsSuccess() {
			return ErrorAsteriskAuthenticated
		} else {
			c.Raw = msg
		}
	}
	return nil
}

// create initializes a new AMI client with the provided network connection.
// It returns the initialized AMI client along with a cancellable context for managing its lifecycle.
//
// Parameters:
//   - conn: The network connection used by the AMI client.
//
// Returns:
//   - An initialized AMI client (*AMI).
//   - A cancellable context for managing the AMI client's lifecycle (context.Context).
//
// Example:
//
//	// Creating an AMI client with a network connection
//	conn, err := net.Dial("tcp", "localhost:5038")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	client, ctx := create(conn)
//	// Use the client and context for AMI operations.
//	// Make sure to close the connection and cancel the context when done.
//	defer client.Close()
//	defer client.Cancel()
func create(conn net.Conn) (*AMI, context.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	c := &AMI{
		Reader: textproto.NewReader(bufio.NewReader(conn)),
		Writer: bufio.NewWriter(conn),
		Conn:   conn,
		Cancel: cancel,
	}
	if conn != nil {
		addr := conn.RemoteAddr().String()
		_socket, err := WithSocket(ctx, addr)
		if err == nil {
			c.Socket = _socket
			D().Info("Ami network cloning (addr: %v) socket connection succeeded", addr)
		}
	}
	return c, ctx
}

// serve initializes a new AMI client, establishes a connection, performs the authentication process, and releases resources.
//
// Parameters:
//   - conn: The net.Conn representing the connection to the Asterisk Manager Interface (AMI) server.
//   - request: The AmiClient configuration containing connection and authentication details.
//
// Returns:
//   - *AMI: A pointer to the initialized AMI client if the process succeeds.
//   - error: An error indicating any issues during the initialization or authentication process.
//
// Example:
//
//	amiClient, err := serve(conn, AmiClient{
//	  username: "admin",
//	  password: "secret*password",
//	  timeout:  time.Second * 5,
//	})
//
// Note: The serve function creates a new AMI client instance, applies necessary settings and configurations, establishes
// a connection to the AMI server, performs the authentication process, and releases resources. If any step in the process
// fails, it returns an appropriate error. If the initialization and authentication are successful, it returns a pointer to
// the AMI client ready for further interaction with the Asterisk server.
func serve(conn net.Conn, request AmiClient) (*AMI, error) {
	ins, ctx := create(conn)
	err := ins.apply(ctx, request.timeout)
	if err != nil {
		return nil, err
	}
	err = ins.authenticate(ctx, request)
	if err != nil {
		return nil, err
	}
	ins.release(ctx)
	return ins, nil
}
