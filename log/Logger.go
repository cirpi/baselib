package log

import (
	"io"
	"time"

	json "github.com/bytedance/sonic"
)

type Logger struct {
	out io.Writer
}

func NewLogger(out io.Writer) *Logger {
	return &Logger{out: out}
}

func (l *Logger) Info(msg string) {
	l.log("INFO", msg, nil, nil)
}

func (l *Logger) Error(msg string, err error) {
	l.log("ERROR", msg, nil, err)
}

func (l *Logger) InfoMeta(msg string, meta map[string]any) {
	l.log("INFO", msg, meta, nil)
}

func (l *Logger) log(level Level, msg string, meta map[string]any, er error) {
	entry := log{
		Time:  time.Now().Local().Format(time.UnixDate),
		Level: level,
		Msg:   msg,
		Meta:  meta,
		Er:    er,
	}

	b, _ := json.Marshal(entry)
	l.out.Write(b)
	l.out.Write([]byte("\n"))
}
