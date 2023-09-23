package utils

import (
	"log"

	jsonI "github.com/json-iterator/go"
)

var _json = jsonI.ConfigCompatibleWithStandardLibrary

func ToJson(data interface{}) string {
	s, ok := data.(string)
	if ok {
		return s
	}
	// result, err := json.Marshal(data)
	result, err := MarshalToString(data)
	if err != nil {
		log.Printf(err.Error())
		return ""
	}
	return string(result)
}

func ToJsonPretty(data interface{}) string {
	s, ok := data.(string)

	if ok {
		return s
	}

	result, err := MarshalIndent(data, "", "    ")

	if err != nil {
		log.Printf(err.Error())
		return ""
	}

	return string(result)
}

func MarshalToString(v interface{}) (string, error) {
	return _json.MarshalToString(v)
}

func Marshal(v interface{}) ([]byte, error) {
	return _json.Marshal(v)
}

func Unmarshal(data []byte, v interface{}) error {
	return _json.Unmarshal(data, v)
}

func UnmarshalFromString(str string, v interface{}) error {
	return _json.UnmarshalFromString(str, v)
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return _json.MarshalIndent(v, prefix, indent)
}
