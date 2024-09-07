# VoipKit

_VoipKit_ is an Asterisk VOIP library implemented in the Go programming language (Golang). It provides a powerful and flexible toolkit for building scalable and high-performance VOIP applications. This user guide will walk you through the installation, requirements, and key features of VoipKit.

## Requirement

Go version 1.20 or higher

## Installation

You can install VoipKit using one of the following methods:

- Use a specific version (tag)

```bash
go get github.com/pnguyen215/voipkit@v1.0.1
```

- Latest version

```bash
go get -u github.com/pnguyen215/voipkit@latest
```

## Features

VoipKit provides a comprehensive set of features for building and managing VOIP applications, including:

- **Asterisk Manager Interface (AMI)**: Enables communication between the Go application and the Asterisk server.
- **Event Socket**: Allows real-time event handling and monitoring.
- **Call Control**: Provides advanced call handling capabilities.
- **Configuration Management**: Supports configuration management and easy integration with existing Asterisk configurations.
- **Error Handling**: Includes detailed error handling and logging for easy debugging.

## Usage

To start using VoipKit in your Go project, import the library:

```go
import (
    "github.com/pnguyen215/voipkit"
)
```

### Connecting to Asterisk

To establish a connection with the Asterisk server, you need to create an AMI instance and provide the necessary configuration:

```go
func CreateConn() (*ami.AMI, error) {
    c := ami.GetAmiClientSample().
        SetEnabled(true).
        SetHost("127.0.0.1").
        SetPort(5038).
        SetUsername("username").
        SetPassword("password").
        SetTimeout(10 * time.Second)
    ami.D().Debug("Asterisk server credentials: %v", c.String())
    return ami.NewClient(ami.NewTcp(), *c)
}
```

### Retrieving SIP Peers

To retrieve the list of SIP peers from Asterisk, you can use the _GetSIPPeers_ method:

```go
func GetSIPPeers() {
    c, err := CreateConn()
    if err != nil {
        ami.D().Error(err.Error())
        return
    }
    c.Core().AddSession()
    peers, err := c.Core().GetSIPPeers(c.Context())
    if err != nil {
        ami.D().Error(err.Error())
        return
    }
    for _, v := range peers {
        ami.D().Info("Peer: %v", v.Get("object_name"))
    }
}
```

### Retrieving SIP Peer Status

To retrieve the status of a specific SIP peer, you can use the _GetSIPPeerStatus_ method:

```go
func GetSIPPeerStatus() {
   c, err := createConn()
    if err != nil {
        ami.D().Error(err.Error())
        return
    }
    c.Core().AddSession()
    peer, err := c.Core().GetSIPPeerStatus(c.Context(), "1004")
    if err != nil {
        ami.D().Error(err.Error())
        return
    }
    ami.D().Info("Peer: %v| Status: %v", peer.Get("peer"), peer.Get("peer_status"))
}
```

### Executing Asterisk CLI Commands

To execute Asterisk CLI commands, you can use the Command method:

```go
func ExecuteCLICommand() {
    c, err := CreateConn()
    if err != nil {
        ami.D().Error(err.Error())
        return
    }
    c.Socket().SetDebugMode(true)
    c.Core().AddSession()
    response, err := c.Core().Command(c.Context(), "sip show users")
    if err != nil {
        ami.D().Error(err.Error())
        return
    }
    fmt.Println(ami.JsonString(response))
```

## Github Action

VoipKit leverages Github Actions for continuous integration and deployment. The following actions are used:

- [Action gh-release](https://github.com/softprops/action-gh-release): This action automates the creation of Github releases.
- [Github Security guides](https://docs.github.com/en/actions/security-guides): Follow these security guides to ensure the safety and integrity of your Github Actions workflows.

## Reference

- [Asterisk Documentation](https://docs.asterisk.org): Official documentation for Asterisk, providing comprehensive guides and reference materials.
- [Official Asterisk Repositories](https://github.com/asterisk/documentation): Github repositories containing Asterisk documentation and related resources.

## Conclusion

VoipKit is a powerful and flexible library for building VOIP applications using Go and Asterisk. With its rich set of features and modular architecture, it enables developers to create scalable and efficient VOIP solutions. By following this user guide and leveraging the provided resources, you can quickly get started with VoipKit and build robust VOIP applications.
