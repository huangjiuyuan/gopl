package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	// When a panic occurs, all deferred functions are run in reverse order,
	// starting with those of the topmost function on the stack and proceeding
	// up to main.
	defer printStack()
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
