package ami

import "context"

type AMIService interface {
	// Connected returns the client status.
	Connected() bool
	// Close closes the client connection.
	Close(ctx context.Context) error
	// Send data from client to server.
	Send(message string) error
	// Received receives a string from server.
	Received(ctx context.Context) (string, error)
}
