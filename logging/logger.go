package logging

import (
	"fmt"
	"io"
)

type Severity byte

var (
	SeverityDebug Severity = 0
	SeverityInfo  Severity = 1
	SeverityWarn  Severity = 2
	SeverityError Severity = 3
)

const (
	severityDebugString = "DEBUG"
	severityInfoString  = "INFO"
	severityWarnString  = "WARN"
	severityErrorString = "ERROR"
)

func (s Severity) String() string {
	switch s {
	case SeverityDebug:
		return severityDebugString
	case SeverityInfo:
		return severityInfoString
	case SeverityWarn:
		return severityWarnString
	case SeverityError:
		return severityErrorString
	}
	panic(fmt.Sprintf("unknown severity: %d", s))
}

type Logger interface {
	Debug(msg string)
	Info(msg string)
	Warn(msg string)
	Error(msg string)
}

type WriterLogger struct {
	severity Severity
	w        io.Writer
}

func NewWriterLogger(w io.Writer, severity Severity) *WriterLogger {
	return &WriterLogger{w: w, severity: severity}
}

func (l *WriterLogger) Debug(msg string) {
	if l.severity > SeverityDebug {
		return
	}
	io.WriteString(l.w, SeverityDebug.String())
	io.WriteString(l.w, ": ")
	io.WriteString(l.w, msg)
	io.WriteString(l.w, "\n")
}

func (l *WriterLogger) Info(msg string) {
	if l.severity > SeverityInfo {
		return
	}
	io.WriteString(l.w, SeverityInfo.String())
	io.WriteString(l.w, ": ")
	io.WriteString(l.w, msg)
	io.WriteString(l.w, "\n")
}

func (l *WriterLogger) Warn(msg string) {
	if l.severity > SeverityWarn {
		return
	}
	io.WriteString(l.w, SeverityWarn.String())
	io.WriteString(l.w, ": ")
	io.WriteString(l.w, msg)
	io.WriteString(l.w, "\n")
}

func (l *WriterLogger) Error(msg string) {
	if l.severity > SeverityError {
		return
	}
	io.WriteString(l.w, SeverityError.String())
	io.WriteString(l.w, ": ")
	io.WriteString(l.w, msg)
	io.WriteString(l.w, "\n")
}
