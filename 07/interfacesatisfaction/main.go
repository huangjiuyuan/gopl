package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

type IntSet struct{}

func main() {
	var w io.Writer
	w = os.Stdout         // OK: *os.File has Write method
	w = new(bytes.Buffer) // OK: *bytes.Buffer has Write method
	// w = time.Second
	// compile error: time.Duration lacks Write method

	var rwc io.ReadWriteCloser
	rwc = os.Stdout // OK: *os.File has Read, Write, Close methods
	// rwc = new(bytes.Buffer)
	// compile error: *bytes.Buffer lacks Close method

	w = rwc // OK: io.ReadWriteCloser has Write method
	// rwc = w
	// compile error: io.Writer lacks Close method
	fmt.Println(w, rwc)

	// var _ = IntSet{}.String()
	// compile error: String requires *IntSet receiver
	var s IntSet
	var _ = s.String() // OK: s is a variable and &s has a String method

	// Only *IntSet has a String method, only *IntSet satisfies the
	// fmt.Stringer interface.
	var _ fmt.Stringer = &s // OK
	// var _ fmt.Stringer = s
	// compile error: IntSet lacks String method

	os.Stdout.Write([]byte("hello\n")) // OK: *os.File has Write method
	os.Stdout.Close()                  // OK: *os.File has Close method
	w = os.Stdout
	w.Write([]byte("hello\n")) // OK: io.Writer has Write method
	// w.Close()
	// compile error: io.Writer lacks Close method

	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	any = new(bytes.Buffer)
	fmt.Println(any)

	// The declaration below asserts at compile time that a value of type
	// *bytes.Buffer satisfies io.Writer.
	// *bytes.Buffer must satisfy io.Writer
	var _ io.Writer = new(bytes.Buffer)
	var _ io.Writer = (*bytes.Buffer)(nil)
}

func (*IntSet) String() string { return "" }
