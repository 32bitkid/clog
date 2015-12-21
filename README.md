# clogo

Simple conditional logging for go.

See the [documentation](https://godoc.org/github.com/32bitkid/clogo).

## Installation

```bash
$ go get github.com/32bitkid/clogo
```

## Examples
clogo is silent by default...

```go
log := NewLog("package")
log.Println("This is a log")
fmt.Println("<end>")
```

```bash
$ go run test.go
<end>
```

...but can easily be configured to show all registered loggers...


```go
log := NewLog("package")

log.Println("This is a log")
fmt.Println("<end>")
```

```bash
$ DEBUG=* go run test.go
package: This is a log!
<end>
```

...or only show logs from a specific namespace...

```go
 foolog := NewLog("foo")
 barlog := NewLog("bar")
 foolog.Println("This is a log")
 barlog.Println("This is a log")
 fmt.Println("<end>")
```

```bash
$ DEBUG=foo go run test.go
foo: This is a log
<end>
```

...and even match namespaces with wildcards.

```go
barlog := NewLog("foo:bar")
bazlog := NewLog("foo:baz")
quxlog := NewLog("qux")

barlog.Println("This is a log")
bazlog.Println("This is a log")
quxlog.Println("This is a log")

fmt.Println("<end>")
```

```bash
$ DEBUG=foo:* go run test.go
foo:bar This is a log
foo:baz This is a log
<end>
```
