package ami

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"net/textproto"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/pnguyen215/gobase-voip-core/pkg/ami/config"
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/fatal"
	"github.com/pnguyen215/gobase-voip-core/pkg/ami/utils"
)

// OpenContext
func OpenContext(conn net.Conn) (*AMI, context.Context) {
	ctx, cancel := context.WithCancel(context.Background())

	client := &AMI{
		Reader: textproto.NewReader(bufio.NewReader(conn)),
		Writer: bufio.NewWriter(conn),
		Conn:   conn,
		Cancel: cancel,
	}

	// checking conn available
	if conn != nil {
		addr := conn.RemoteAddr().String()
		_socket, err := NewAMISocketWith(ctx, addr)

		if err == nil {
			client.Socket = _socket
			log.Printf("OpenContext, cloning (addr: %v) socket connection succeeded", addr)
		}
	}

	return client, ctx
}

// OpenDial
func OpenDial(ip string, port int) (net.Conn, error) {
	return OpenDialWith(config.AmiNetworkTcpKey, ip, port)
}

// OpenDialWith
func OpenDialWith(network, ip string, port int) (net.Conn, error) {

	if !config.AmiNetworkKeys[network] {
		return nil, fatal.AMIErrorNew("AMI: Invalid network")
	}

	if ip == "" {
		return nil, fatal.AMIErrorNew("AMI: IP must be not empty")
	}

	if port <= 0 {
		return nil, fatal.AMIErrorNew("AMI: Port must be positive number")
	}

	host, _port, _ := utils.IPDecode(ip)

	if len(host) > 0 && len(_port) > 0 {
		form := net.JoinHostPort(host, _port)
		log.Printf("AMI: (IP decoded) dial connection = %v", form)
		return net.Dial(network, form)
	}

	form := RemoveProtocol(ip, port)
	log.Printf("AMI: dial connection = %v", form)
	return net.Dial(network, form)
}

// RemoveProtocol
// Return form as string: <ip>:<port>
// Example:
// Ip: http://127.0.0.1 or https://127.0.0.1
// Port: 18080
// Result: 127.0.0.1:18080
func RemoveProtocol(ip string, port int) string {
	if ip == "" {
		return ip
	}

	if port < 0 {
		return ip
	}

	if strings.HasPrefix(ip, config.AmiProtocolHttpKey) {
		ip = strings.Replace(ip, config.AmiProtocolHttpKey, "", -1)
	}

	if strings.HasPrefix(ip, config.AmiProtocolHttpsKey) {
		ip = strings.Replace(ip, config.AmiProtocolHttpsKey, "", -1)
	}

	_ip := strings.Split(ip, ":")
	ip = _ip[0]

	form := net.JoinHostPort(ip, strconv.Itoa(port))
	return form
}

// JoinHostPortString
func JoinHostPortString(ip string, port int) string {
	return RemoveProtocol(ip, port)
}

// JoinHostPortStrings
func JoinHostPortStrings(ip []string, port int) (result []string) {
	if len(ip) == 0 {
		return ip
	}
	for _, v := range ip {
		result = append(result, JoinHostPortString(v, port))
	}
	return result
}

// WriteString
func WriteString(buf *bytes.Buffer, tag, value string) {
	if len(tag) > 0 {
		buf.WriteString(tag)
		buf.WriteString(": ")
	}
	buf.WriteString(value)
	buf.WriteString(config.AmiSignalLetter)
}

func IsOmitempty(tag string) (string, bool, error) {
	fields := strings.Split(tag, ",")
	if len(fields) > 1 {
		for _, flag := range fields[1:] {
			if strings.EqualFold(strings.TrimSpace(flag), config.AmiOmitemptyKeyRef) {
				return fields[0], true, nil
			}
			return tag, false, fmt.Errorf("unsupported flag %q in tag %q", flag, tag)
		}
	}
	return tag, false, nil
}

func IsZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return len(v.String()) == 0
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		return v.Int() == 0
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		return v.Uint() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Struct:
		for i := v.NumField() - 1; i >= 0; i-- {
			if !IsZero(v.Field(i)) {
				return false
			}
		}
		return true
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	}
	return false
}

func Encode(buf *bytes.Buffer, tag string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.String:
		WriteString(buf, tag, v.String())
	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		WriteString(buf, tag, strconv.FormatInt(v.Int(), 10))
	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		WriteString(buf, tag, strconv.FormatUint(v.Uint(), 10))
	case reflect.Bool:
		WriteString(buf, tag, strconv.FormatBool(v.Bool()))
	case reflect.Float32:
		WriteString(buf, tag, strconv.FormatFloat(v.Float(), 'E', -1, 32))
	case reflect.Float64:
		WriteString(buf, tag, strconv.FormatFloat(v.Float(), 'E', -1, 64))
	case reflect.Ptr, reflect.Interface:
		if !v.IsNil() {
			return Encode(buf, tag, v.Elem())
		}
	case reflect.Struct:
		return EncodeStruct(buf, v)
	case reflect.Map:
		return EncodeMap(buf, v)
	case reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			elem := v.Index(i)
			if !IsZero(elem) {
				if err := Encode(buf, tag, elem); err != nil {
					return err
				}
			}
		}
	default:
		return fmt.Errorf("unsupported kind %v", v.Kind())
	}
	return nil
}

func EncodeStruct(buf *bytes.Buffer, v reflect.Value) error {
	var omitempty bool
	var err error
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		tag, ok := field.Tag.Lookup(config.AmiTagKeyRef)
		switch {
		case !ok:
			tag = string(field.Tag)
		case tag == "-":
			continue
		}
		tag, omitempty, err = IsOmitempty(tag)
		if err != nil {
			return err
		}
		value := v.Field(i)
		if omitempty && IsZero(value) {
			continue
		}

		if err := Encode(buf, tag, value); err != nil {
			return err
		}
	}
	return nil
}

func EncodeMap(buf *bytes.Buffer, v reflect.Value) error {
	for _, key := range v.MapKeys() {
		value := v.MapIndex(key)
		if key.Kind() == reflect.String {
			tag := key.String()
			if err := Encode(buf, tag, value); err != nil {
				return err
			}
		}
	}
	return nil
}

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := Encode(&buf, "", reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	buf.WriteString(config.AmiSignalLetter)
	return buf.Bytes(), nil
}

// GenUUID returns a new UUID based on /dev/urandom (unix).
func GenUUID() (string, error) {
	file, err := os.Open("/dev/urandom")
	if err != nil {
		return "", fmt.Errorf("open /dev/urandom error:[%v]", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}()
	b := make([]byte, 16)

	_, err = file.Read(b)
	if err != nil {
		return "", err
	}
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return uuid, nil
}

func GenUUIDShorten() string {
	uuid, err := GenUUID()
	if err != nil {
		return ""
	}
	return uuid
}

// IsSuccess
// Check event from asterisk feedback to console is succeeded
func IsSuccess(raw AMIResultRaw) bool {
	if len(raw) == 0 {
		return false
	}
	response := raw.GetVal(strings.ToLower(config.AmiResponseKey))
	return IsResponse(raw) &&
		strings.EqualFold(response, config.AmiStatusSuccessKey)
}

// IsFailure
// Check event from asterisk feedback to console is failure
func IsFailure(raw AMIResultRaw) bool {
	return !IsSuccess(raw)
}

// IsEvent
// Check result from asterisk server to console is event?
// Get key `Event` and value of `Event` is not equal whitespace
func IsEvent(raw AMIResultRaw) bool {
	if len(raw) == 0 {
		return false
	}
	event := raw.GetVal(strings.ToLower(config.AmiEventKey))
	return event != ""
}

// IsResponse
// Check result from asterisk server to console is response?
// Get key `response` and value of `response` is not equal whitespace
func IsResponse(raw AMIResultRaw) bool {
	if len(raw) == 0 {
		return false
	}
	response := raw.GetVal(strings.ToLower(config.AmiResponseKey))
	return response != ""
}

// ParseResult
// Break line by line for parsing to map[string]string
func ParseResult(socket AMISocket, raw string) (AMIResultRaw, error) {
	response := make(AMIResultRaw)
	lines := strings.Split(raw, config.AmiSignalLetter)

	for _, line := range lines {
		keys := strings.SplitAfterN(line, ":", 2)

		if len(keys) == 2 {
			key := strings.TrimSpace(strings.Trim(keys[0], ":"))
			value := strings.TrimSpace(keys[1])
			response[key] = value
		} else if strings.Contains(line, config.AmiSignalLetters) || line == "" {
			break
		}
	}

	return TransformKey(response, socket.Dictionary), nil
}

// ParseResultLevel
// Break line by line for parsing to map[string]string
func ParseResultLevel(socket AMISocket, raw string) (AMIResultRawLevel, error) {
	response := make(AMIResultRawLevel)
	lines := strings.Split(raw, config.AmiSignalLetter)

	for _, line := range lines {
		keys := strings.SplitAfterN(line, ":", 2)

		if len(keys) == 2 {
			key := strings.TrimSpace(strings.Trim(keys[0], ":"))
			value := strings.TrimSpace(keys[1])
			response[key] = append(response[key], value)
		} else if strings.Contains(line, config.AmiSignalLetters) || line == "" {
			break
		}
	}

	return TransformKeyLevel(response, socket.Dictionary), nil
}

// DoGetResult
// Get result while fetch response command has been sent to asterisk server
// Arguments:
// 1. AMISocket - to create new instance connection socket
// 2. AMICommand - to build command cli will be sent to server
// 3. acceptedEvents - select event will captured as response
// 4. ignoreEvents - the event will been stopped fetching command
func DoGetResult(ctx context.Context, s AMISocket, c *AMICommand, acceptedEvents []string, ignoreEvents []string) ([]AMIResultRaw, error) {
	return c.DoGetResult(ctx, s, c, acceptedEvents, ignoreEvents)
}

// TransformKey
// Find the key transferred from dictionary
// Example:
// The field key is Response, so then transferred to response
// Or from ResponseEvent to response_event
func TransformKey(response AMIResultRaw, d *AMIDictionary) AMIResultRaw {
	if len(response) <= 0 {
		return response
	}

	if d == nil {
		return response
	}

	_m := make(AMIResultRaw, len(response))
	for k, v := range response {
		_m[d.TranslateField(k)] = v
	}
	response = nil
	return _m
}

// TransformKeyLevel
// Find the key transferred from dictionary
// Example:
// The field key is Response, so then transferred to response
// Or from ResponseEvent to response_event
func TransformKeyLevel(response AMIResultRawLevel, d *AMIDictionary) AMIResultRawLevel {
	if len(response) <= 0 {
		return response
	}

	if d == nil {
		return response
	}

	_m := make(AMIResultRawLevel, len(response))
	for k, v := range response {
		_m[d.TranslateField(k)] = v
	}
	response = nil
	return _m
}
