package error

import (
	json "github.com/bytedance/sonic"
)

type Error struct {
	Op   string         `json:"op,omitempty"`
	Er   error          `json:"cause"`
	Meta map[string]any `json:"meta,omitempty"`
}

func (e *Error) Error() string {
	res, er := json.MarshalString(e)
	if er != nil {
		return er.Error()
	}
	return res
}

func (e *Error) Unwrap() error {
	return e.Er
}

func NewError(operation string, er error) error {
	return &Error{Er: er, Op: operation}
}

func AddMeta(er error, key string, value any) error {
	libEr, canCast := er.(*Error)
	if canCast {
		if libEr.Meta == nil {
			libEr.Meta = map[string]any{}
		}
		libEr.Meta[key] = value
	}
	return er
}
