# clogo

Simple conditional logging for go.

See the [documentation](https://godoc.org/github.com/32bitkid/clogo).

## Installation

```bash
$ go get github.com/32bitkid/clogo
```

## Examples

```go
package main

import "github.com/32bitkid/clogo"
import "fmt"

var logger = clogo.NewLog("prefix")

func main() {
    logger.Println("This is a log")
    fmt.Println("<end>")
}
```

Silent by default:

```bash
$ go run main.go
<end>
```

Opt-in logging:

```bash
$ DEBUG=* go run main.go
prefix: This is a log
<end>
```

