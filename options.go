package timeout

import (
	"net/http"
)

type CallBackFunc func(*http.Request)
type Option func(*TimeoutWriter)

type TimeoutOptions struct {
	CallBack             CallBackFunc
	DefaultMsg           interface{}
	Timeout              uint64
	MaxTimeout           uint64
	MinTimeout           uint64
	ErrorHttpCode        int
	CustomHeader         map[string][]string
	AllowInfinityTimeout bool
}

func WithTimeout(d uint64) Option {
	return func(t *TimeoutWriter) {
		t.Timeout = d
	}
}

func WithCustomHeader(k string, v []string) Option {
	return func(t *TimeoutWriter) {
		if t.CustomHeader == nil {
			t.CustomHeader = make(map[string][]string)
		}
		t.CustomHeader[k] = v
	}
}

// Optional parameters
func WithErrorHttpCode(code int) Option {
	return func(t *TimeoutWriter) {
		t.ErrorHttpCode = code
	}
}

// Optional parameters
func WithDefaultMsg(resp interface{}) Option {
	return func(t *TimeoutWriter) {
		t.DefaultMsg = resp
	}
}

// Optional parameters
func WithCallBack(f CallBackFunc) Option {
	return func(t *TimeoutWriter) {
		t.CallBack = f
	}
}

func WithAllowInfinityTimeoutFlag(allowInfinityTimeout bool) Option {
	return func(t *TimeoutWriter) {
		t.AllowInfinityTimeout = allowInfinityTimeout
	}
}
