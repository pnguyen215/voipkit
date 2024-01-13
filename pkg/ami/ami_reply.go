package ami

import "github.com/pnguyen215/voipkit/pkg/ami/config"

func NewAmiReply() *AmiReply {
	s := &AmiReply{}
	return s
}

func NewAmiReplies() *AmiReplies {
	s := &AmiReplies{}
	return s
}

// Get retrieves the value associated with the specified key in the AmiReply map.
// If the key is not found or the value is empty, an empty string is returned.
func (s AmiReply) Get(key string) string {
	if s == nil {
		return ""
	}
	if len(s) == 0 {
		return ""
	}
	v := s[key]
	if len(v) == 0 {
		return ""
	}
	return v
}

// GetOrFallback retrieves the value associated with the specified key in the AmiReply map.
// If the value is empty, it falls back to the value associated with the fallback_key.
func (s AmiReply) GetOrFallback(key, fallback_key string) string {
	_v := s.Get(key)
	if len(_v) == 0 {
		return s.Get(fallback_key)
	}
	return _v
}

// GetOrFallbacks retrieves the value associated with the specified key in the AmiReply map.
// If the value is empty, it falls back to the values associated with the fallback_keys, checking in order.
func (s AmiReply) GetOrFallbacks(key string, fallback_keys ...string) string {
	if len(fallback_keys) == 0 {
		return s.GetOrFallback(key, "")
	}
	_v := ""
	for _, v := range fallback_keys {
		_v = s.GetOrFallback(key, v)
		if len(_v) > 0 {
			break
		}
	}
	return _v
}

// Values returns a slice containing unique values from the AmiReply map.
// Values are filtered based on the fields specified in config.AmiJsonIgnoringFieldType.
func (s AmiReply) Values() []string {
	if len(s) == 0 {
		return []string{}
	}
	var result []string
	for k := range s {
		if config.AmiJsonIgnoringFieldType[k] {
			continue
		}
		v := s.Get(k)
		if !Contains(result, v) {
			result = append(result, v)
		}
	}
	return result
}

// Size returns the number of unique values in the AmiReply map.
func (s AmiReply) Size() int {
	return len(s.Values())
}

// Get retrieves the value associated with the specified key in the AmiReplies map.
// If the key is not found or the value is empty, an empty string is returned.
// If the key has multiple values, it returns a JSON string representation of the values.
func (s AmiReplies) Get(key string) string {
	if s == nil {
		return ""
	}
	if len(s) == 0 {
		return ""
	}
	v := s[key]
	if len(v) == 0 {
		return ""
	}
	if len(v) == 1 {
		return v[0]
	}
	return JsonString(v)
}

// GetOrFallback retrieves the value associated with the specified key in the AmiReplies map.
// If the value is empty, it falls back to the value associated with the fallback_key.
func (s AmiReplies) GetOrFallback(key, fallback_key string) string {
	_v := s.Get(key)
	if len(_v) == 0 {
		return s.Get(fallback_key)
	}
	return _v
}

// GetOrFallbacks retrieves the value associated with the specified key in the AmiReplies map.
// If the value is empty, it falls back to the values associated with the fallback_keys, checking in order.
func (s AmiReplies) GetOrFallbacks(key string, fallback_keys ...string) string {
	if len(fallback_keys) == 0 {
		return s.GetOrFallback(key, "")
	}
	_v := ""
	for _, v := range fallback_keys {
		_v = s.GetOrFallback(key, v)
		if len(_v) > 0 {
			break
		}
	}
	return _v
}

// Values returns a slice containing unique values from the AmiReplies map.
// Values are filtered based on the fields specified in config.AmiJsonIgnoringFieldType.
func (s AmiReplies) Values() []string {
	if len(s) == 0 {
		return []string{}
	}
	var result []string
	for k := range s {
		if config.AmiJsonIgnoringFieldType[k] {
			continue
		}
		v := s.Get(k)
		if !Contains(result, v) {
			result = append(result, v)
		}
	}
	return result
}

// Size returns the number of unique values in the AmiReplies map.
func (s AmiReplies) Size() int {
	return len(s.Values())
}
