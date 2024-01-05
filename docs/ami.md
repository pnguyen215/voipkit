# Asterisk Manager Interface (AMI)

## Usage

1. Open connection to asterisk server

```Go
func NewAmi(host string, port int, username, password string) (*AMI, error)
func NewAmiDial(conn net.Conn, username, password string) (*AMI, error)
func NewAmiWithTimeout(host string, port int, username, password string, timeout time.Duration) (*AMI, error)
func NewAmiWith(conn net.Conn, username, password string, timeout time.Duration) (*AMI, error)
```

<i>To create new connection using net conn</i>

```Go
func OpenDial(ip string, port int) (net.Conn, error)
func OpenDialWith(network, ip string, port int) (net.Conn, error)
```

##### Example

```Go
package main

import (
	"log"

	"github.com/pnguyen215/voipkit/pkg/ami"
)

func NewClient() (*ami.AMI, error) {
	client, err := ami.NewAmi("127.0.0.1", 5038, "u_username", "u_password")
	return client, err
}

func main() {

	client, err := NewClient()

	if err != nil {
		log.Fatal(err)
		return
	}
}
```

2. Custom listen all events from asterisk
   // note: using go func(){...}

```Go
// AllEvents subscribes to any AMI message received from Asterisk server
// returns send-only channel or nil
func (c *AMI) AllEvents() <-chan *AMIMessage
// OnEvent subscribes by event name (case insensitive) and
// returns send-only channel or nil
func (c *AMI) OnEvent(name string) <-chan *AMIMessage
// OnEvents subscribes by events name (case insensitive) and
// return send-only channel or nil
func (c *AMI) OnEvents(keys ...string) <-chan *AMIMessage
```

##### Example

```Go
package main

import (
	"log"
	"strings"

	"github.com/pnguyen215/voipkit/pkg/ami"
	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

func NewClient() (*ami.AMI, error) {
	client, err := ami.NewAmi("127.0.0.1", 5038, "u_username", "u_password")
	return client, err
}

func main() {

	client, err := NewClient()

	if err != nil {
		log.Fatal(err)
		return
	}

	go func() {
		all := client.AllEvents() // or client.OnEvents(keys...)
		defer client.Close()

		for {
			select {
			case message := <-all:
				log.Printf("ami event: '%s' received = %s", message.Field(strings.ToLower(config.AmiEventKey)), message.Json())
			case err := <-client.Error():
				client.Close()
				log.Printf("ami listener has error occurred = %s", err.Error())
			}
		}
	}()
}
```

3. Listen event by scratch

##### Example

```Go
package main

import (
	"log"

	"github.com/pnguyen215/voipkit/pkg/ami"
)

func NewClient() (*ami.AMI, error) {
	client, err := ami.NewAmi("127.0.0.1", 5038, "u_username", "u_password")
	return client, err
}

func main() {

	client, err := NewClient()

	if err != nil {
		log.Fatal(err)
		return
	}

	// Or using listener event shorten
	event := ami.NewEventListener()
	go func() {
		event.OpenFullEvents(client)
	}()
}
```

4. Listen events and get field from event by adding key-value you want to take value from that field

##### Example

```Go
package main

import (
	"log"

	"github.com/pnguyen215/voipkit/pkg/ami"
)

func NewClient() (*ami.AMI, error) {
	client, err := ami.NewAmi("127.0.0.1", 5038, "u_username", "u_password")
	return client, err
}

func main() {

	client, err := NewClient()

	if err != nil {
		log.Fatal(err)
		return
	}

	// Or using listener event shorten
	event := ami.NewEventListener()
	dictionary := ami.NewDictionary() // called translator key-value via field event

	dictionary.AddKeyTranslator("field_event_ast", "to_field_event")

	go func() {
		event.OpenFullEventsTranslator(client, dictionary)
	}()
}
```

5. Listen events and get field from event by adding key-value you want to take value from that field.
   You can customize event by adding new callback function

##### Example

```Go
package main

import (
	"log"

	"github.com/pnguyen215/voipkit/pkg/ami"
)

func NewClient() (*ami.AMI, error) {
	client, err := ami.NewAmi("127.0.0.1", 5038, "u_username", "u_password")
	return client, err
}

func main() {

	client, err := NewClient()

	if err != nil {
		log.Fatal(err)
		return
	}

	// Or using listener event shorten
	event := ami.NewEventListener()
	dictionary := ami.NewDictionary()

	dictionary.AddKeyTranslator("field_event_ast", "to_field_event")

	go func() {
		event.OpenFullEventsCallbackTranslator(client, dictionary, func(e *ami.AMIMessage, json string, err error) {
			// e is an original event
			// json is event has been marshall to json string
			// err if you get any error from asterisk server feedback
		})
	}()
}
```
