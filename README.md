# waitfor-proc
OS process readiness assertion library

```go
package main

import (
	"context"
	"fmt"
	"github.com/go-waitfor/waitfor"
	"github.com/go-waitfor/waitfor-proc"
	"os"
)

func main() {
	runner := waitfor.New(proc.Use())

	err := runner.Test(
		context.Background(),
		[]string{"proc://node"},
		waitfor.WithAttempts(5),
	)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
```