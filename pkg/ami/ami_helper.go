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
func IsSuccess(raw AMIResultRaw) bool {
	if len(raw) == 0 {
		return false
	}
	response := raw.GetVal(strings.ToLower(config.AmiResponseKey))
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
func IsFailure(raw AMIResultRaw) bool {
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

func VarsMap(values []string) map[string]string {
	r := make(map[string]string)
	for _, value := range values {
		k, v := VarsSplit(value)
		r[k] = v
	}
	return r
}

func VarsSplit(value string) (string, string) {
	s := strings.SplitN(value, "=", 2)
	k := s[0]
	if len(s) == 1 {
		return k, ""
	}
	return k, s[1]
}

func UsableRConnection(ip string, port int) (bool, error) {
	timeout := time.Second
	conn, err := net.DialTimeout(config.AmiNetworkTcpKey, net.JoinHostPort(ip, strconv.Itoa(port)), timeout)
	if err != nil {
		log.Printf("Connecting error: %v", err)
		return false, err
	}
	if conn != nil {
		defer conn.Close()
		log.Printf("Opened on: %s", net.JoinHostPort(ip, strconv.Itoa(port)))
		return true, nil
	}
	return false, nil
}

func UsableRConnectionWith(ip string, ports []int) (bool, error) {
	for _, port := range ports {
		if ok, err := UsableRConnection(ip, port); err != nil {
			return ok, err
		}
	}
	return true, nil
}

// DecodeIp
// Decode IP into 2 parts: host, port
func DecodeIp(ip string) (string, string, error) {
	u, err := url.Parse(ip)
	if err != nil {
		return "", "", err
	}
	host, port, err := net.SplitHostPort(u.Host)
	return host, port, err
}

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

// Contains check slice contains value or not
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func TrimStringSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

// ApplyTimezone
func ApplyTimezone(at time.Time, timezone string) (time.Time, error) {
	loc, err := time.LoadLocation(timezone)
	now := at.In(loc)
	return now, err
}

// AdjustTimezone
func AdjustTimezone(at time.Time, timezone string) time.Time {
	t, err := ApplyTimezone(at, timezone)
	if err != nil {
		return at
	}
	return t
}

func Base64Encode(v interface{}) string {
	d := JsonString(v)
	return base64.StdEncoding.EncodeToString([]byte(d))
}

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

func JsonString(data interface{}) string {
	s, ok := data.(string)
	if ok {
		return s
	}
	result, err := json.Marshal(data)
	if err != nil {
		log.Printf(err.Error())
		return ""
	}
	return string(result)
}
