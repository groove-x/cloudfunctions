package log

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/logging"
	mrpb "google.golang.org/genproto/googleapis/api/monitoredres"
)

var projectID string

func init() {
	projectID = os.Getenv("GCP_PROJECT")

	ctx := context.Background()
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("failed to create logging client: %s", err)
	}

	logName := "cloudfunctions.googleapis.com%2Fcloud-functions"
	functionName := os.Getenv("FUNCTION_NAME")
	region := os.Getenv("FUNCTION_REGION")

	std = client.Logger(logName, logging.CommonResource(&mrpb.MonitoredResource{
		Type: "cloud_function",
		Labels: map[string]string{
			"function_name": functionName,
			"project_id":    projectID,
			"region":        region,
		},
	}))
}

var (
	std   *logging.Logger
	entry logging.Entry
)

func WithRequest(r *http.Request) {
	entry = logging.Entry{
		Labels: map[string]string{
			"execution_id": r.Header.Get("Function-Execution-Id"),
		},
		HTTPRequest: &logging.HTTPRequest{Request: r},
	}
}

func StandardLogger(s logging.Severity) *log.Logger {
	return std.StandardLogger(s)
}

func Debug(args ...interface{}) {
	entry := entry
	entry.Severity = logging.Debug
	entry.Payload = fmt.Sprint(args...)
	std.Log(entry)
}

func Print(args ...interface{}) {
	entry := entry
	entry.Severity = logging.Info
	entry.Payload = fmt.Sprint(args...)
	std.Log(entry)
}

func Info(args ...interface{}) {
	entry := entry
	entry.Severity = logging.Info
	entry.Payload = fmt.Sprint(args...)
	std.Log(entry)
}

func Warn(args ...interface{}) {
	entry := entry
	entry.Severity = logging.Warning
	entry.Payload = fmt.Sprint(args...)
	std.Log(entry)
}

func Error(args ...interface{}) {
	entry := entry
	entry.Severity = logging.Error
	entry.Payload = fmt.Sprint(args...)
	std.Log(entry)
}

func Fatal(args ...interface{}) {
	entry := entry
	entry.Severity = logging.Critical
	entry.Payload = fmt.Sprint(args...)
	std.Log(entry)
	std.Flush()
	os.Exit(1)
}

func Debugf(format string, args ...interface{}) {
	entry := entry
	entry.Severity = logging.Debug
	entry.Payload = fmt.Sprintf(format, args...)
	std.Log(entry)
}

func Printf(format string, args ...interface{}) {
	entry := entry
	entry.Severity = logging.Info
	entry.Payload = fmt.Sprintf(format, args...)
	std.Log(entry)
}

func Infof(format string, args ...interface{}) {
	entry := entry
	entry.Severity = logging.Info
	entry.Payload = fmt.Sprintf(format, args...)
	std.Log(entry)
}

func Warnf(format string, args ...interface{}) {
	entry := entry
	entry.Severity = logging.Warning
	entry.Payload = fmt.Sprintf(format, args...)
	std.Log(entry)
}

func Errorf(format string, args ...interface{}) {
	entry := entry
	entry.Severity = logging.Error
	entry.Payload = fmt.Sprintf(format, args...)
	std.Log(entry)
}

func Fatalf(format string, args ...interface{}) {
	entry := entry
	entry.Severity = logging.Critical
	entry.Payload = fmt.Sprintf(format, args...)
	std.Log(entry)
	std.Flush()
	os.Exit(1)
}

func Debugln(args ...interface{}) {
	entry := entry
	entry.Severity = logging.Debug
	entry.Payload = fmt.Sprintln(args...)
	std.Log(entry)
}

func Println(args ...interface{}) {
	entry := entry
	entry.Severity = logging.Info
	entry.Payload = fmt.Sprintln(args...)
	std.Log(entry)
}

func Infoln(args ...interface{}) {
	entry := entry
	entry.Severity = logging.Info
	entry.Payload = fmt.Sprintln(args...)
	std.Log(entry)
}

func Warnln(args ...interface{}) {
	entry := entry
	entry.Severity = logging.Warning
	entry.Payload = fmt.Sprintln(args...)
	std.Log(entry)
}

func Errorln(args ...interface{}) {
	entry := entry
	entry.Severity = logging.Error
	entry.Payload = fmt.Sprintln(args...)
	std.Log(entry)
}

func Fatalln(args ...interface{}) {
	entry := entry
	entry.Severity = logging.Critical
	entry.Payload = fmt.Sprintln(args...)
	std.Log(entry)
	std.Flush()
	os.Exit(1)
}
