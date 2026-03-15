package log

type Level string

const (
	_INFO  Level = "INFO"
	_DEBUG Level = "DEBUG"
	_WARN  Level = "WARN"
	_ERROR Level = "ERROR"
	_FATAL Level = "FATAL"
)

type log struct {
	Time  string         `json:"t"`
	Level Level          `json:"level,omitempty"`
	Msg   string         `json:"msg,omitempty"`
	Er    error          `json:"er,omitempty"`
	Meta  map[string]any `json:"meta,omitempty"`
}
