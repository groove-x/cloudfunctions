package log

import (
	"log"

	"cloud.google.com/go/logging"
)

type localLogger struct {
	logger *log.Logger
}

func (l *localLogger) Log(e logging.Entry) {
	l.logger.Printf("[%s] %v", e.Severity, e.Payload)
}

func (l *localLogger) Flush() error {
	return nil
}

func (l *localLogger) StandardLogger(s logging.Severity) *log.Logger {
	return l.logger
}
