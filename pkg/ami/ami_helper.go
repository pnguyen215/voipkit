package ami

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/nyaruka/phonenumbers"
	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

// RemoveProtocol removes the protocol prefix and port from the given IP address.
// If the IP address is empty or the port is negative, the original IP address is returned unchanged.
// The function supports removing both "http://" and "https://" protocol prefixes.
// The resulting IP address is formatted as a string without the protocol prefix and with the specified port,
// or without the port if it was negative or not provided.
//
// Parameters:
//   - ip: The input IP address with an optional protocol prefix and port.
//   - port: The port number to be included in the formatted IP address.
//     If the port is negative, it is excluded from the formatted IP address.
//
// Returns:
//   - The formatted IP address as a string without the protocol prefix and with the specified port,
//     or without the port if it was negative or not provided.
//
// Example:
//
//	input: "http://127.0.0.1:8088", port: 8088
//	output: "127.0.0.1:8088"
//
//	input: "https://example.com", port: -1
//	output: "example.com"
//
//	input: "invalid", port: 5060
//	output: "invalid:5060"
func RemoveProtocol(ip string, port int) string {
	if IsStringEmpty(ip) || port < 0 {
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

// JoinHostPortString joins the given IP address and port into a formatted string.
// The function utilizes the RemoveProtocol function to handle the removal of protocol prefixes.
//
// Parameters:
//   - ip: The input IP address with an optional protocol prefix and port.
//   - port: The port number to be included in the formatted IP address.
//     If the port is negative, it is excluded from the formatted IP address.
//
// Returns:
//   - The formatted IP address as a string without the protocol prefix and with the specified port,
//     or without the port if it was negative or not provided.
//
// Example:
//
//	input: "http://127.0.0.1:8088", port: 8088
//	output: "127.0.0.1:8088"
//
//	input: "https://example.com", port: -1
//	output: "example.com"
//
//	input: "invalid", port: 5060
//	output: "invalid:5060"
func JoinHostPortString(ip string, port int) string {
	return RemoveProtocol(ip, port)
}

// JoinHostPortStrings joins the given slice of IP addresses with the specified port into formatted strings.
// Each IP address in the slice may have an optional protocol prefix and port, which is handled by the JoinHostPortString function.
//
// Parameters:
//   - ip: The input slice of IP addresses, where each IP address may have an optional protocol prefix and port.
//   - port: The port number to be included in the formatted IP addresses.
//     If the port is negative, it is excluded from the formatted IP addresses.
//
// Returns:
//   - A slice of formatted IP addresses as strings without the protocol prefix and with the specified port,
//     or without the port if it was negative or not provided.
//
// Example:
//
//	input: []string{"http://127.0.0.1:8088", "https://example.com", "invalid"}, port: 5060
//	output: []string{"127.0.0.1:8088", "example.com", "invalid:5060"}
func JoinHostPortStrings(ip []string, port int) (result []string) {
	if len(ip) == 0 {
		return ip
	}
	for _, v := range ip {
		result = append(result, JoinHostPortString(v, port))
	}
	return result
}

// WriteString writes a formatted key-value pair to the provided bytes.Buffer.
// The key (tag) and value are concatenated with appropriate separators, and a newline signal is added at the end.
// If the key (tag) is an empty string, it is excluded from the output.
//
// Parameters:
//   - buf: A pointer to the bytes.Buffer where the formatted string will be written.
//   - tag: The key (tag) representing the identifier of the value.
//   - value: The value associated with the key.
//
// Example:
//
//	buf := &bytes.Buffer{}
//	WriteString(buf, "Action", "Login")
//	result: "Action: Login\r\n"
//
//	buf := &bytes.Buffer{}
//	WriteString(buf, "", "ValueOnly")
//	result: "ValueOnly\r\n"
func WriteString(buf *bytes.Buffer, tag, value string) {
	if len(tag) > 0 {
		buf.WriteString(tag)
		buf.WriteString(": ")
	}
	buf.WriteString(value)
	buf.WriteString(config.AmiSignalLetter)
}

// IsOmitempty checks if the given struct field tag contains the "omitempty" flag.
//
// Parameters:
//   - tag: The struct field tag to be analyzed.
//
// Returns:
//   - A tuple containing the modified tag without the "omitempty" flag (if present),
//     a boolean indicating whether "omitempty" was present, and an error (if any).
//
// Example:
//
//	tag := `json:"name,omitempty"`
//	result, omitempty, err := IsOmitempty(tag)
//	// result: "json:\"name\"", omitempty: true, err: nil
//
//	tag := `json:"age"`
//	result, omitempty, err := IsOmitempty(tag)
//	// result: "json:\"age\"", omitempty: false, err: nil
//
//	tag := `json:"salary,omitempty,unsupported"`
//	result, omitempty, err := IsOmitempty(tag)
//	// result: "", omitempty: false, err: "unsupported flag \"unsupported\" in tag \"json:\"salary,omitempty,unsupported\""
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

// IsZero checks if the given reflect.Value is considered "zero" based on its kind.
//
// Parameters:
//   - v: The reflect.Value to be checked for being "zero."
//
// Returns:
//   - A boolean indicating whether the provided reflect.Value is considered "zero."
//
// Example:
//
//	var str string
//	result := IsZero(reflect.ValueOf(str))
//	// result: true
//
//	var num int
//	result := IsZero(reflect.ValueOf(num))
//	// result: true
//
//	var slice []int
//	result := IsZero(reflect.ValueOf(slice))
//	// result: true
//
//	var ptr *int
//	result := IsZero(reflect.ValueOf(ptr))
//	// result: true
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

// Encode writes the encoded Asterisk Manager Interface (AMI) command parameters to the provided bytes.Buffer.
// The encoding is based on the provided struct field tag and the reflect.Value of the corresponding struct field.
// Supported types for encoding include string, integer types, boolean, floating-point types, pointers, interfaces, structs, maps, and slices.
//
// Parameters:
//   - buf: A pointer to the bytes.Buffer where the encoded parameters will be written.
//   - tag: The struct field tag specifying the key (identifier) for the AMI command parameter.
//   - v: The reflect.Value representing the value to be encoded.
//
// Returns:
//   - An error if there's an issue during encoding; otherwise, returns nil.
//
// Example:
//
//	type ExampleStruct struct {
//	    Name   string `ami:"Name"`
//	    Age    int    `ami:"Age,omitempty"`
//	    Active bool   `ami:"Active"`
//	}
//
//	buf := &bytes.Buffer{}
//	data := ExampleStruct{Name: "John", Age: 30, Active: true}
//	err := Encode(buf, "ExampleAction", reflect.ValueOf(data))
//	// Encoded data in buf:
//	// "ExampleAction: Name: John\r\nAge: 30\r\nActive: true\r\n"
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

// EncodeStruct writes the encoded Asterisk Manager Interface (AMI) command parameters for a struct to the provided bytes.Buffer.
// The encoding is based on the struct field tags and the reflect.Value of each struct field.
// Supported types for encoding include string, integer types, boolean, floating-point types, pointers, interfaces, structs, maps, and slices.
//
// Parameters:
//   - buf: A pointer to the bytes.Buffer where the encoded parameters will be written.
//   - v: The reflect.Value representing the struct whose fields are to be encoded.
//
// Returns:
//   - An error if there's an issue during encoding; otherwise, returns nil.
//
// Example:
//
//	type ExampleStruct struct {
//	    Name   string `ami:"Name"`
//	    Age    int    `ami:"Age,omitempty"`
//	    Active bool   `ami:"Active"`
//	}
//
//	buf := &bytes.Buffer{}
//	data := ExampleStruct{Name: "John", Age: 30, Active: true}
//	err := EncodeStruct(buf, reflect.ValueOf(data))
//	// Encoded data in buf:
//	// "Name: John\r\nAge: 30\r\nActive: true\r\n"
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

// EncodeMap writes the encoded Asterisk Manager Interface (AMI) command parameters for a map to the provided bytes.Buffer.
// The encoding is based on the keys (as tags) and values of the provided reflect.Value representing the map.
// Supported types for encoding include string keys and values, integer types, boolean, floating-point types, pointers, interfaces, structs, maps, and slices.
//
// Parameters:
//   - buf: A pointer to the bytes.Buffer where the encoded parameters will be written.
//   - v: The reflect.Value representing the map whose keys and values are to be encoded.
//
// Returns:
//   - An error if there's an issue during encoding; otherwise, returns nil.
//
// Example:
//
//	data := map[string]interface{}{
//	    "Name":   "John",
//	    "Age":    30,
//	    "Active": true,
//	}
//
//	buf := &bytes.Buffer{}
//	err := EncodeMap(buf, reflect.ValueOf(data))
//	// Encoded data in buf:
//	// "Name: John\r\nAge: 30\r\nActive: true\r\n"
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

// Marshal serializes the provided data structure into a slice of bytes based on the Asterisk Manager Interface (AMI) command parameters.
// The encoding is performed using reflection, and the resulting byte slice includes the encoded parameters followed by a newline signal.
//
// Parameters:
//   - v: The data structure (e.g., struct, map) to be serialized.
//
// Returns:
//   - A slice of bytes containing the serialized representation of the provided data structure.
//   - An error if there's an issue during serialization; otherwise, returns nil.
//
// Example:
//
//	type ExampleStruct struct {
//	    Name   string `ami:"Name"`
//	    Age    int    `ami:"Age,omitempty"`
//	    Active bool   `ami:"Active"`
//	}
//
//	data := ExampleStruct{Name: "John", Age: 30, Active: true}
//	result, err := Marshal(data)
//	// Serialized data in result:
//	// "Name: John\r\nAge: 30\r\nActive: true\r\n"
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := Encode(&buf, "", reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	buf.WriteString(config.AmiSignalLetter)
	return buf.Bytes(), nil
}

// GenUUID generates a universally unique identifier (UUID) using a cryptographically secure random number source.
// The UUID is represented as a string with the format "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx" where each 'x' is a hexadecimal digit.
//
// Returns:
//   - A string representing the generated UUID.
//   - An error if there's an issue during UUID generation; otherwise, returns nil.
//
// Example:
//
//	uuid, err := GenUUID()
//	// Example output: "3d96b8b9-4a84-4f69-9f9b-8f7ab9e6a96a"
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

// GenUUIDShorten generates a universally unique identifier (UUID) using a cryptographically secure random number source.
// The UUID is represented as a string with the format "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx" where each 'x' is a hexadecimal digit.
//
// Returns:
//   - A string representing the generated UUID.
//   - An error if there's an issue during UUID generation; otherwise, returns nil.
//
// Example:
//
//	uuid, err := GenUUID()
//	// Example output: "3d96b8b9-4a84-4f69-9f9b-8f7ab9e6a96a"
func GenUUIDShorten() string {
	uuid, err := GenUUID()
	if err != nil {
		return ""
	}
	return uuid
}

// IsSuccess checks whether an Asterisk Manager Interface (AMI) command response indicates success.
// It evaluates the provided raw AMI result to determine if the response is non-empty, has a valid response key,
// and the response status is considered successful.
//
// Parameters:
//   - raw: The AMIResultRaw representing the raw result of an AMI command response.
//
// Returns:
//   - A boolean value indicating whether the response is successful.
//
// Example:
//
//	raw := AMIResultRaw{"Response": "Success", "ActionID": "123", "Event": "SomeEvent"}
//	success := IsSuccess(raw)
//	// success is true if the response indicates success, otherwise false.
func IsSuccess(raw AmiReply) bool {
	if len(raw) == 0 {
		return false
	}
	response := raw.Get(strings.ToLower(config.AmiResponseKey))
	return IsResponse(raw) &&
		strings.EqualFold(response, config.AmiStatusSuccessKey)
}

// IsFailure checks whether an Asterisk Manager Interface (AMI) command response indicates failure.
// It evaluates the provided raw AMI result to determine if the response is non-empty and does not have a successful status.
//
// Parameters:
//   - raw: The AMIResultRaw representing the raw result of an AMI command response.
//
// Returns:
//   - A boolean value indicating whether the response is a failure.
//
// Example:
//
//	raw := AMIResultRaw{"Response": "Error", "ActionID": "123", "Message": "Command failed"}
//	failure := IsFailure(raw)
//	// failure is true if the response indicates failure, otherwise false.
func IsFailure(raw AmiReply) bool {
	return !IsSuccess(raw)
}

// IsEvent checks whether an Asterisk Manager Interface (AMI) command response represents an event.
// It evaluates the provided raw AMI result to determine if the response is non-empty and contains a non-whitespace value for the 'Event' key.
//
// Parameters:
//   - raw: The AMIResultRaw representing the raw result of an AMI command response.
//
// Returns:
//   - A boolean value indicating whether the response is an event.
//
// Example:
//
//	raw := AMIResultRaw{"Event": "SomeEvent", "ActionID": "123"}
//	isEvent := IsEvent(raw)
//	// isEvent is true if the response represents an event, otherwise false.
func IsEvent(raw AmiReply) bool {
	if len(raw) == 0 {
		return false
	}
	event := raw.Get(strings.ToLower(config.AmiEventKey))
	return event != ""
}

// IsResponse
// Check result from asterisk server to console is response?
// Get key `response` and value of `response` is not equal whitespace
func IsResponse(raw AmiReply) bool {
	if len(raw) == 0 {
		return false
	}
	response := raw.Get(strings.ToLower(config.AmiResponseKey))
	return response != ""
}

// ParseReply
// Break line by line for parsing to map[string]string
func ParseReply(socket AMISocket, raw string) (AmiReply, error) {
	response := make(AmiReply)
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

// ParseReplies
// Break line by line for parsing to map[string]string
func ParseReplies(socket AMISocket, raw string) (AmiReplies, error) {
	response := make(AmiReplies)
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
func DoGetResult(ctx context.Context, s AMISocket, c *AMICommand, acceptedEvents []string, ignoreEvents []string) ([]AmiReply, error) {
	return c.DoGetResult(ctx, s, c, acceptedEvents, ignoreEvents)
}

// TransformKey
// Find the key transferred from dictionary
// Example:
// The field key is Response, so then transferred to response
// Or from ResponseEvent to response_event
func TransformKey(response AmiReply, d *AMIDictionary) AmiReply {
	if len(response) <= 0 {
		return response
	}

	if d == nil {
		return response
	}

	_m := make(AmiReply, len(response))
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
func TransformKeyLevel(response AmiReplies, d *AMIDictionary) AmiReplies {
	if len(response) <= 0 {
		return response
	}

	if d == nil {
		return response
	}

	_m := make(AmiReplies, len(response))
	for k, v := range response {
		_m[d.TranslateField(k)] = v
	}
	response = nil
	return _m
}

// VerifyPhoneNoCustomize checks whether a given phone number string is in a valid format based on a customized regular expression pattern.
//
// Parameters:
//   - phone: The phone number string to be verified.
//
// Returns:
//   - A boolean value indicating whether the provided phone number is in a valid format.
//
// Example:
//
//	isValid := VerifyPhoneNoCustomize("+1 (123) 456-7890")
//	// isValid is true if the phone number is in a valid format, otherwise false.
func VerifyPhoneNoCustomize(phone string) bool {
	if IsStringEmpty(phone) {
		return false
	}
	matcher := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	return matcher.MatchString(phone)
}

// VerifyPhoneNo checks whether a given phone number string is both a valid and possible phone number
// based on the specified region using the Google's libphonenumber library.
// Additionally, it verifies the phone number using a customized regular expression pattern.
//
// Parameters:
//   - phone:  The phone number string to be verified.
//   - region: The ISO 3166-1 alpha-2 country code representing the region associated with the phone number.
//
// Returns:
//   - A boolean value indicating whether the provided phone number is both valid and possible
//     for the specified region and passes the additional customization check.
//
// Example:
//
//	isValid := VerifyPhoneNo("+1 (123) 456-7890", "US")
//	// isValid is true if the phone number is both valid and possible for the US region,
//	// and it passes the additional customization check; otherwise, false.
func VerifyPhoneNo(phone string, region string) bool {
	if IsStringEmpty(phone) {
		return false
	}

	p, err := phonenumbers.Parse(phone, region)
	if err != nil {
		log.Printf(err.Error())
		return false
	}
	v := phonenumbers.IsValidNumber(p)
	l := phonenumbers.IsPossibleNumber(p)
	return v && l && VerifyPhoneNoCustomize(phone)
}

// RemoveStringPrefix removes specified prefixes from the beginning of a given string.
//
// Parameters:
//   - str:    The string from which prefixes should be removed.
//   - prefix: One or more prefixes to be removed from the beginning of the string.
//
// Returns:
//   - A string with the specified prefixes removed from the beginning.
//
// Example:
//
//	result := RemoveStringPrefix("example-string", "example-", "prefix-")
//	// result is "string" after removing both "example-" and "prefix-" prefixes.
func RemoveStringPrefix(str string, prefix ...string) string {
	if IsStringEmpty(str) {
		return str
	}
	if len(prefix) == 0 {
		return str
	}
	for _, v := range prefix {
		str = strings.TrimPrefix(str, v)
	}
	return str
}

// IsStringEmpty checks whether a given string is empty or consists only of whitespace characters.
//
// Parameters:
//   - str: The string to be checked for emptiness.
//
// Returns:
//   - A boolean value indicating whether the provided string is empty or consists only of whitespace characters.
//
// Example:
//
//	isEmpty := IsStringEmpty("   ")
//	// isEmpty is true if the string is empty or consists only of whitespace characters; otherwise, false.
func IsStringEmpty(str string) bool {
	return len(str) == 0 || str == "" || strings.TrimSpace(str) == ""
}

// ForkDictionaryFromLink retrieves a dictionary (map[string]string) from a specified link
// where the content is formatted as JSON. The function performs a GET request to the provided link.
//
// Parameters:
//   - link:  The URL link to the JSON-formatted file containing the dictionary data.
//   - debug: A boolean flag indicating whether to enable debugging for the HTTP requests.
//
// Returns:
//   - A pointer to a map[string]string representing the dictionary retrieved from the provided link.
//   - An error if the HTTP request or JSON decoding fails.
//
// Example:
//
//	dict, err := ForkDictionaryFromLink("https://example.com/dictionary.json", true)
//	// dict is a pointer to a map[string]string containing the dictionary data,
//	// and err is an error indicating any issues during the retrieval process.
func ForkDictionaryFromLink(link string, debug bool) (*map[string]string, error) {
	c := NewRestify(link)
	c.SetDebug(debug)
	c.SetMaxRetries(3)
	c.AddHeader("Content-Type", "application/json")
	c.SetRetryCondition(func(response *http.Response, err error) bool {
		return response.StatusCode >= 400 || response.StatusCode >= 500
	})
	result := &map[string]string{}
	err := c.Get("", make(map[string]string), &result)
	return result, err
}

// VarsMap creates a map[string]string from a slice of string values, typically representing
// key-value pairs in the format "key=value". It uses the VarsSplit function to extract
// individual key-value pairs from each string in the input slice.
//
// Parameters:
//   - values: A slice of strings representing key-value pairs, where each string is in the format "key=value".
//
// Returns:
//   - A map[string]string containing the extracted key-value pairs from the input slice.
//
// Example:
//
//	values := []string{"name=John", "age=25", "city=New York"}
//	result := VarsMap(values)
//	// result is a map[string]string{"name": "John", "age": "25", "city": "New York"}.
func VarsMap(values []string) map[string]string {
	r := make(map[string]string)
	for _, value := range values {
		k, v := VarsSplit(value)
		r[k] = v
	}
	return r
}

// VarsSplit splits a string in the format "key=value" into its key and value components.
//
// Parameters:
//   - value: A string in the format "key=value" to be split.
//
// Returns:
//   - Two strings representing the key and value components extracted from the input string.
//
// Example:
//
//	value := "name=John"
//	key, val := VarsSplit(value)
//	// key is "name" and val is "John".
func VarsSplit(value string) (string, string) {
	s := strings.SplitN(value, "=", 2)
	k := s[0]
	if len(s) == 1 {
		return k, ""
	}
	return k, s[1]
}

// UsableRConnection checks the usability of a remote connection to a specified IP address and port.
// It attempts to establish a TCP connection within a specified timeout duration.
//
// Parameters:
//   - ip:   The IP address of the remote server.
//   - port: The port number on the remote server.
//
// Returns:
//   - A boolean value indicating whether the connection to the remote server is usable.
//   - An error if there is an issue during the connection attempt.
//
// Example:
//
//	isUsable, err := UsableRConnection("127.0.0.1", 8080)
//	// isUsable is true if the connection is usable, and err is an error indicating any connection issues.
func UsableRConnection(ip string, port int) (bool, error) {
	timeout := time.Second
	conn, err := net.DialTimeout(config.AmiNetworkTcpKey, net.JoinHostPort(ip, strconv.Itoa(port)), timeout)
	if err != nil {
		D().Error("Connecting error: %v", err)
		return false, err
	}
	if conn != nil {
		defer conn.Close()
		D().Info("Opened on: %s", net.JoinHostPort(ip, strconv.Itoa(port)))
		return true, nil
	}
	return false, nil
}

// UsableRConnectionWith checks the usability of a remote connection to a specified IP address
// across multiple ports. It iterates through the provided list of ports and attempts to establish
// a TCP connection to each port within a specified timeout duration.
//
// Parameters:
//   - ip:    The IP address of the remote server.
//   - ports: A slice of port numbers on the remote server to check for usability.
//
// Returns:
//   - A boolean value indicating whether at least one of the connections to the remote server is usable.
//   - An error if there is an issue during any of the connection attempts.
//
// Example:
//
//	isUsable, err := UsableRConnectionWith("127.0.0.1", []int{8080, 8888, 9999})
//	// isUsable is true if at least one of the connections is usable, and err is an error indicating any connection issues.
func UsableRConnectionWith(ip string, ports []int) (bool, error) {
	for _, port := range ports {
		if ok, err := UsableRConnection(ip, port); err != nil {
			return ok, err
		}
	}
	return true, nil
}

// DecodeIp decodes an IP address into two parts: host and port.
//
// Parameters:
//   - ip: A string representing the IP address to decode.
//
// Returns:
//   - Two strings: host and port extracted from the input IP address.
//   - An error if there is an issue during the decoding process.
//
// Example:
//
//	host, port, err := DecodeIp("http://example.com:8080")
//	// host is "example.com", port is "8080", and err is an error indicating any decoding issues.
func DecodeIp(ip string) (string, string, error) {
	u, err := url.Parse(ip)
	if err != nil {
		return "", "", err
	}
	host, port, err := net.SplitHostPort(u.Host)
	return host, port, err
}

// GetKeyByVal searches for a value in a map and returns the corresponding key.
//
// Parameters:
//   - values: A map[string]string representing key-value pairs to search.
//   - value:  The value to search for within the map.
//
// Returns:
//   - The key associated with the specified value, if found.
//   - If the value is not found, the function returns the input value itself.
//
// Example:
//
//	values := map[string]string{"one": "1", "two": "2", "three": "3"}
//	key := GetKeyByVal(values, "2")
//	// key is "two" since it corresponds to the value "2" in the map.
func GetKeyByVal(values map[string]string, value string) string {
	if len(values) <= 0 {
		return value
	}
	for k, v := range values {
		if strings.EqualFold(v, value) {
			return k
		}
	}
	return value
}

// GetValByKey searches for a key in a map and returns the corresponding value.
//
// Parameters:
//   - values: A map[string]string representing key-value pairs to search.
//   - key:    The key to search for within the map.
//
// Returns:
//   - The value associated with the specified key, if found.
//   - If the key is not found, the function returns an empty string.
//
// Example:
//
//	values := map[string]string{"one": "1", "two": "2", "three": "3"}
//	value := GetValByKey(values, "two")
//	// value is "2" since it corresponds to the key "two" in the map.
func GetValByKey(values map[string]string, key string) string {
	if len(values) <= 0 {
		return key
	}
	for k, v := range values {
		if strings.EqualFold(k, key) {
			return v
		}
	}
	return ""
}

// GetKeys retrieves the keys from a map or converts the indices of a slice to strings.
//
// Parameters:
//   - in: An interface{} representing either a map or a slice.
//
// Returns:
//   - A slice of strings containing the keys of the map or the stringified indices of the slice.
//   - If the input is not a map or a slice, an empty slice is returned.
//
// Example:
//
//	// For a map
//	m := map[string]int{"one": 1, "two": 2, "three": 3}
//	keys := GetKeys(m)
//	// keys is ["one", "two", "three"]
//
//	// For a slice
//	s := []int{1, 2, 3}
//	indices := GetKeys(s)
//	// indices is ["0", "1", "2"]
//
//	// For an unsupported type
//	unsupported := "unsupported"
//	result := GetKeys(unsupported)
//	// result is an empty slice since the input type is not supported.
func GetKeys(in interface{}) (keys []string) {
	switch z := in.(type) {
	case map[string]int:
	case map[string]int32:
	case map[string]int64:
	case map[string]float32:
	case map[string]float64:
	case map[string]string:
	case map[string]bool:
		for k := range z {
			keys = append(keys, k)
		}
	case []int:
		for _, k := range z {
			keys = append(keys, strconv.Itoa(k))
		}
	default:
		return []string{}
	}
	return keys
}

// MergeMaps merges two maps into a new map, combining their key-value pairs.
//
// Parameters:
//   - m1: A map[K]V representing the first map.
//   - m2: A map[K]V representing the second map.
//
// Type Parameters:
//   - K: The key type of the maps (must be comparable).
//   - V: The value type of the maps.
//
// Returns:
//   - A new map[K]V containing the merged key-value pairs from m1 and m2.
//   - If both m1 and m2 are empty maps, an empty map is returned.
//
// Example:
//
//	// Merging maps with string keys and int values
//	map1 := map[string]int{"one": 1, "two": 2}
//	map2 := map[string]int{"three": 3, "four": 4}
//	mergedMap := MergeMaps(map1, map2)
//	// mergedMap is {"one": 1, "two": 2, "three": 3, "four": 4}
//
//	// Merging maps with int keys and string values
//	map3 := map[int]string{1: "one", 2: "two"}
//	map4 := map[int]string{3: "three", 4: "four"}
//	mergedMapStrings := MergeMaps(map3, map4)
//	// mergedMapStrings is {1: "one", 2: "two", 3: "three", 4: "four"}
func MergeMaps[K comparable, V any](m1 map[K]V, m2 map[K]V) map[K]V {
	merged := make(map[K]V)
	if len(m1) > 0 {
		for key, value := range m1 {
			merged[key] = value
		}
	}
	if len(m2) > 0 {
		for key, value := range m2 {
			merged[key] = value
		}
	}
	return merged
}

// Contains checks if a given element is present in a slice.
//
// Parameters:
//   - s: A slice of elements of type T.
//   - e: The element of type T to check for in the slice.
//
// Type Parameters:
//   - T: The type of elements in the slice (must be comparable).
//
// Returns:
//   - true if the element e is found in the slice s; otherwise, false.
//
// Example:
//
//	// Checking if an integer is present in a slice of integers
//	numbers := []int{1, 2, 3, 4, 5}
//	containsThree := Contains(numbers, 3) // true
//	containsTen := Contains(numbers, 10)  // false
//
//	// Checking if a string is present in a slice of strings
//	fruits := []string{"apple", "banana", "orange"}
//	containsBanana := Contains(fruits, "banana") // true
//	containsGrape := Contains(fruits, "grape")   // false
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// TrimStringSpaces trims leading and trailing spaces and collapses multiple spaces within a string into a single space.
//
// Parameters:
//   - s: The input string to be trimmed.
//
// Returns:
//   - A new string with leading and trailing spaces removed and multiple spaces collapsed into a single space.
//
// Example:
//
//	inputString := "   This    is   a   sample    string.   "
//	trimmedString := TrimStringSpaces(inputString)
//	// trimmedString is "This is a sample string."
func TrimStringSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// ApplyTimezone applies a specified timezone to a given time and returns the time in the specified timezone.
//
// Parameters:
//   - at: The time.Time value to be adjusted to the specified timezone.
//   - timezone: The IANA Time Zone identifier representing the desired timezone (e.g., "America/New_York").
//
// Returns:
//   - A new time.Time value adjusted to the specified timezone.
//   - An error, if any, encountered during the process of loading the timezone.
//
// Example:
//
//	// Applying the "America/New_York" timezone to a specific time
//	originalTime := time.Date(2023, time.March, 15, 12, 0, 0, 0, time.UTC)
//	adjustedTime, err := ApplyTimezone(originalTime, "America/New_York")
//	if err != nil {
//	    log.Printf("Error applying timezone: %v", err)
//	}
//	// adjustedTime now represents the original time adjusted to the "America/New_York" timezone.
func ApplyTimezone(at time.Time, timezone string) (time.Time, error) {
	loc, err := time.LoadLocation(timezone)
	now := at.In(loc)
	return now, err
}

// AdjustTimezone attempts to apply a specified timezone to a given time and returns the adjusted time.
// If an error occurs during the timezone adjustment, the original time is returned.
//
// Parameters:
//   - at: The time.Time value to be adjusted to the specified timezone.
//   - timezone: The IANA Time Zone identifier representing the desired timezone (e.g., "America/New_York").
//
// Returns:
//   - A new time.Time value adjusted to the specified timezone, or the original time if an error occurs.
//
// Example:
//
//	// Adjusting the time to the "America/New_York" timezone or keeping the original time in case of an error
//	originalTime := time.Date(2023, time.March, 15, 12, 0, 0, 0, time.UTC)
//	adjustedTime := AdjustTimezone(originalTime, "America/New_York")
//	// adjustedTime now represents the original time adjusted to the "America/New_York" timezone,
//	// or it remains the same as the original time in case of an error.
func AdjustTimezone(at time.Time, timezone string) time.Time {
	t, err := ApplyTimezone(at, timezone)
	if err != nil {
		return at
	}
	return t
}

// JsonString converts the given data to a JSON-formatted string.
// If the input data is already a string, it is returned as is.
// If the conversion to JSON fails, an empty string is returned.
//
// Parameters:
//   - data: The input data to be converted to a JSON-formatted string.
//
// Returns:
//   - A string representing the JSON-formatted version of the input data, or an empty string if conversion fails.
//
// Example:
//
//	// Converting a struct to a JSON-formatted string
//	myData := MyStruct{Field1: "value1", Field2: 42}
//	jsonString := JsonString(myData)
//	// jsonString now contains the JSON-formatted string representation of myData,
//	// or an empty string if the conversion fails.
func JsonString(data interface{}) string {
	s, ok := data.(string)
	if ok {
		return s
	}
	result, err := json.Marshal(data)
	if err != nil {
		D().Error(err.Error())
		return ""
	}
	return string(result)
}

// Base64Encode encodes the given data to a Base64-encoded string.
// The input data is first converted to a JSON-formatted string using the JsonString function.
//
// Parameters:
//   - v: The input data to be Base64-encoded.
//
// Returns:
//   - A string representing the Base64-encoded version of the input data.
//
// Example:
//
//	// Encoding a struct to a Base64-encoded string
//	myData := MyStruct{Field1: "value1", Field2: 42}
//	base64String := Base64Encode(myData)
//	// base64String now contains the Base64-encoded string representation of myData.
func Base64Encode(v interface{}) string {
	d := JsonString(v)
	return base64.StdEncoding.EncodeToString([]byte(d))
}

// Base64Decode decodes the given Base64-encoded string to its original representation.
// The decoded string is returned, or an empty string if decoding fails.
//
// Parameters:
//   - encoded: The Base64-encoded string to be decoded.
//
// Returns:
//   - A string representing the decoded version of the Base64-encoded input string,
//     or an empty string if decoding fails.
//
// Example:
//
//	// Decoding a Base64-encoded string to its original representation
//	base64String := "eyJmb28iOiJiYXIifQ=="
//	decodedString := Base64Decode(base64String)
//	// decodedString now contains the original string representation of the Base64-encoded data.
func Base64Decode(encoded string) string {
	if IsStringEmpty(encoded) {
		return encoded
	}
	d, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return ""
	}
	return string(d)
}
