# Google CloudFunctions support package

## Logging

example for trigger http function:

```go
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
```

deploy:

```
$ gcloud functions deploy function-1 --runtime=go113 --region=asia-northeast1 --timeout=60s --trigger-http --entry-point=ExampleLogging --update-env-vars=FUNCTION_REGION=asia-northeast1
```


### requirement for logging

- [Logs Writer (`roles/logging.logWriter`)](https://cloud.google.com/logging/docs/access-control) role for service account. (if you change the runtime service account from *App Engine default service account*.
  )


