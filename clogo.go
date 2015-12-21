// Simple conditional logging for go.
//
// Leverages the environment variable "DEBUG" to conditionally output statements.
//
package clogo

import deflog "log"
import "io"
import "io/ioutil"
import "os"
import "strings"

var namespaces = strings.Split(os.Getenv("DEBUG"), ",")

// Logger interface collects the basic methods for logging.
type Logger interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

type noopLogger struct{}

func (s noopLogger) Print(v ...interface{})                 {}
func (s noopLogger) Printf(format string, v ...interface{}) {}
func (s noopLogger) Println(v ...interface{})               {}

var noop = deflog.New(ioutil.Discard, "", 0)

func createLogger(ns string, w io.Writer) Logger {
	return deflog.New(w, ns+": ", 0)
}

// NewLogWriter creates a logger that will write to the given io.Writer
func NewLogWriter(ns string, w io.Writer) Logger {
	for _, namespace := range namespaces {
		// Log everything
		if namespace == "*" {
			return createLogger(ns, w)
		}

		// Prefix wildcard match
		if strings.HasSuffix(namespace, ":*") && strings.HasPrefix(ns, namespace[:len(namespace)-2]) {
			return createLogger(ns, w)
		}

		// Exact match
		if ns == namespace {
			return createLogger(ns, w)
		}
	}

	return noop
}

// NewLog creates a conditional logger that writes to os.Stderr
func NewLog(ns string) Logger {
	return NewLogWriter(ns, os.Stderr)
}
