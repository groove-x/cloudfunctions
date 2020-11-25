package function

import (
	"fmt"
	"net/http"

	"github.com/groove-x/cloudfunctions/log"
)

func ExampleLogging(w http.ResponseWriter, r *http.Request) {
	log.WithRequest(r)
	defer log.Flush()

	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")

	logger := log.WithField("key1", 1)
	logger.Info("structured log")

	log.WithFields(map[string]interface{}{"foo": 1, "bar": 2}).Debug("structured debug log")

	fmt.Fprintln(w, "hello")
}
