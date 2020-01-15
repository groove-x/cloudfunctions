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

	log.Debug("debug")
	log.Info("info")
	log.Warn("warn")
	log.Error("error")

	fmt.Fprintln(w, "hello")
}
```
