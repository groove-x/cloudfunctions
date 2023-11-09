package log

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"cloud.google.com/go/compute/metadata"
	"cloud.google.com/go/logging"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	projectID string
	region    string
)

var (
	std   stdLogger
	entry logging.Entry
)

type stdLogger interface {
	Log(e logging.Entry)
	Flush() error
	StandardLogger(s logging.Severity) *log.Logger
}

func init() {
	if !metadata.OnGCE() {
		std = &localLogger{logger: log.New(os.Stderr, "", log.LstdFlags)}
		return
	}

	projectID, _ = metadata.ProjectID()
	instanceRegion, _ := metadata.Get("instance/region")
	region = path.Base(instanceRegion)

	ctx := context.Background()
	client, err := logging.NewClient(ctx, projectID)
	if err != nil {
		log.Fatalf("failed to create logging client: %s", err)
	}

	if err := client.Ping(ctx); err != nil {
		if status.Code(err) == codes.PermissionDenied {
			log.Fatalf(`Caller does not have required permission to use project %[1]s. Grant the caller the roles/logging.logWriter role, or a custom role with the logging.logEntries.create permission, by visiting https://console.developers.google.com/iam-admin/iam/project?project=%[1]s and then retry (propagation of new permission may take a few minutes)., forbidden`, projectID)
		} else {
			log.Fatalf("failed to ping logging server: %s", err)
		}
	}

	name := os.Getenv("FUNCTION_NAME")
	if name == "" {
		name = os.Getenv("K_SERVICE")
	}
	std = client.Logger(name)
}

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

func Flush() error {
	return std.Flush()
}
