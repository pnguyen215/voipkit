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

func (s AmiReply) GetVal(key string) string {
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

func (s AmiReply) Values() []string {
	if len(s) == 0 {
		return []string{}
	}
	var result []string
	for k := range s {
		if config.AmiJsonIgnoringFieldType[k] {
			continue
		}
		v := s.GetVal(k)
		if !Contains(result, v) {
			result = append(result, v)
		}
	}
	return result
}

func (s AmiReply) Size() int {
	return len(s.Values())
}

func (s AmiReply) GetValOrPref(key, pref string) string {
	_v := s.GetVal(key)

	if len(_v) == 0 {
		return s.GetVal(pref)
	}
	return _v
}

func (s AmiReply) GetValOrPrefers(key string, pref ...string) string {
	if len(pref) == 0 {
		return s.GetValOrPref(key, "")
	}
	_v := ""
	for _, v := range pref {
		_v = s.GetValOrPref(key, v)
		if len(_v) > 0 {
			break
		}
	}
	return _v
}

func (s AmiReplies) GetVal(key string) string {
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

func (s AmiReplies) Values() []string {
	if len(s) == 0 {
		return []string{}
	}
	var result []string
	for k := range s {
		if config.AmiJsonIgnoringFieldType[k] {
			continue
		}
		v := s.GetVal(k)
		if !Contains(result, v) {
			result = append(result, v)
		}
	}
	return result
}

func (s AmiReplies) Size() int {
	return len(s.Values())
}

func (s AmiReplies) GetValOrPref(key, pref string) string {
	_v := s.GetVal(key)

	if len(_v) == 0 {
		return s.GetVal(pref)
	}
	return _v
}

func (s AmiReplies) GetValOrPrefers(key string, pref ...string) string {
	if len(pref) == 0 {
		return s.GetValOrPref(key, "")
	}
	_v := ""
	for _, v := range pref {
		_v = s.GetValOrPref(key, v)
		if len(_v) > 0 {
			break
		}
	}
	return _v
}
