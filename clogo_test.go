package clogo

import "fmt"

// Show all registered namespaces:
func Example_all() {
	// Executed with:
	// DEBUG=* go run test.go

	log := NewLog("package")

	log.Println("This is a log")
	fmt.Println("<end>")

	// Output: package: This is a log!
	// <end>
}

// Hide all logging:
func Example_none() {
	// Executed with:
	// go run test.go

	log := NewLog("package")

	log.Println("This is a log")
	fmt.Println("<end>")
	// Output:
	// <end>
}

// Show logs from a specific namespace
func Example_specificNamespace() {
	// Executed with:
	// DEBUG=foo go run test.go

	foolog := NewLog("foo")
	barlog := NewLog("bar")

	foolog.Println("This is a log")
	barlog.Println("This is a log")
	fmt.Println("<end>")
	// Output:
	// foo: This is a log
	// <end>
}

// Match namespaces with wildcards
func Example_namespaceWildcard() {
	// Executed with:
	// DEBUG=foo:* go run test.go

	barlog := NewLog("foo:bar")
	bazlog := NewLog("foo:baz")
	quxlog := NewLog("qux")

	barlog.Println("This is a log")
	bazlog.Println("This is a log")
	quxlog.Println("This is a log")
	fmt.Println("<end>")
	// Output:
	// foo:bar This is a log
	// foo:baz This is a log
	// <end>
}
