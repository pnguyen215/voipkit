package utils

import "strings"

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