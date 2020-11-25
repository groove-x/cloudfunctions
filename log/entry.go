package log

import (
	"fmt"
	"os"

	"cloud.google.com/go/logging"
)

type Entry struct {
	entry   logging.Entry
	payload map[string]interface{}
}

func WithFields(fields map[string]interface{}) *Entry {
	return &Entry{entry: entry, payload: fields}
}

func WithField(key string, value interface{}) *Entry {
	return &Entry{entry: entry, payload: map[string]interface{}{key: value}}
}

func (e *Entry) WithField(key string, value interface{}) *Entry {
	e.payload[key] = value
	return e
}

func (e *Entry) Debug(args ...interface{}) {
	entry := e.entry
	entry.Severity = logging.Debug
	e.payload["message"] = fmt.Sprint(args...)
	entry.Payload = e.payload
	std.Log(entry)
}

func (e *Entry) Print(args ...interface{}) {
	entry := e.entry
	entry.Severity = logging.Info
	e.payload["message"] = fmt.Sprint(args...)
	entry.Payload = e.payload
	std.Log(entry)
}

func (e *Entry) Info(args ...interface{}) {
	entry := e.entry
	entry.Severity = logging.Info
	e.payload["message"] = fmt.Sprint(args...)
	entry.Payload = e.payload
	std.Log(entry)
}

func (e *Entry) Warn(args ...interface{}) {
	entry := e.entry
	entry.Severity = logging.Warning
	e.payload["message"] = fmt.Sprint(args...)
	entry.Payload = e.payload
	std.Log(entry)
}

func (e *Entry) Error(args ...interface{}) {
	entry := e.entry
	entry.Severity = logging.Error
	e.payload["message"] = fmt.Sprint(args...)
	entry.Payload = e.payload
	std.Log(entry)
}

func (e *Entry) Fatal(args ...interface{}) {
	entry := e.entry
	entry.Severity = logging.Critical
	e.payload["message"] = fmt.Sprint(args...)
	entry.Payload = e.payload
	std.Log(entry)
	std.Flush()
	os.Exit(1)
}

func (e *Entry) Debugf(format string, args ...interface{}) {
	entry := e.entry
	entry.Severity = logging.Debug
	e.payload["message"] = fmt.Sprintf(format, args...)
	entry.Payload = e.payload
	std.Log(entry)
}

func (e *Entry) Printf(format string, args ...interface{}) {
	entry := e.entry
	entry.Severity = logging.Info
	e.payload["message"] = fmt.Sprintf(format, args...)
	entry.Payload = e.payload
	std.Log(entry)
}

func (e *Entry) Infof(format string, args ...interface{}) {
	entry := e.entry
	entry.Severity = logging.Info
	e.payload["message"] = fmt.Sprintf(format, args...)
	entry.Payload = e.payload
	std.Log(entry)
}

func (e *Entry) Warnf(format string, args ...interface{}) {
	entry := e.entry
	entry.Severity = logging.Warning
	entry.Payload = fmt.Sprintf(format, args...)
	std.Log(entry)
}

func (e *Entry) Errorf(format string, args ...interface{}) {
	entry := e.entry
	entry.Severity = logging.Error
	e.payload["message"] = fmt.Sprintf(format, args...)
	entry.Payload = e.payload
	std.Log(entry)
}

func (e *Entry) Fatalf(format string, args ...interface{}) {
	entry := e.entry
	entry.Severity = logging.Critical
	e.payload["message"] = fmt.Sprintf(format, args...)
	entry.Payload = e.payload
	std.Log(entry)
	std.Flush()
	os.Exit(1)
}

func (e *Entry) Debugln(args ...interface{}) {
	entry := e.entry
	entry.Severity = logging.Debug
	e.payload["message"] = fmt.Sprintln(args...)
	entry.Payload = e.payload
	std.Log(entry)
}

func (e *Entry) Println(args ...interface{}) {
	entry := e.entry
	entry.Severity = logging.Info
	e.payload["message"] = fmt.Sprintln(args...)
	entry.Payload = e.payload
	std.Log(entry)
}

func (e *Entry) Infoln(args ...interface{}) {
	entry := e.entry
	entry.Severity = logging.Info
	e.payload["message"] = fmt.Sprintln(args...)
	entry.Payload = e.payload
	std.Log(entry)
}

func (e *Entry) Warnln(args ...interface{}) {
	entry := e.entry
	entry.Severity = logging.Warning
	entry.Payload = fmt.Sprintln(args...)
	std.Log(entry)
}

func (e *Entry) Errorln(args ...interface{}) {
	entry := e.entry
	entry.Severity = logging.Error
	e.payload["message"] = fmt.Sprintln(args...)
	entry.Payload = e.payload
	std.Log(entry)
}

func (e *Entry) Fatalln(args ...interface{}) {
	entry := e.entry
	entry.Severity = logging.Critical
	e.payload["message"] = fmt.Sprintln(args...)
	entry.Payload = e.payload
	std.Log(entry)
	std.Flush()
	os.Exit(1)
}
