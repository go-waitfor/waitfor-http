# waitfor-http
HTTP resource readiness assertion library

# Quick start

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-waitfor/waitfor"
	"github.com/go-waitfor/waitfor-http"
	"os"
)

func main() {
	runner := waitfor.New(http.Use())

	err := runner.Test(
		context.Background(),
		[]string{"http://locahost:5432", "https://www.google.com"},
		waitfor.WithAttempts(5),
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```