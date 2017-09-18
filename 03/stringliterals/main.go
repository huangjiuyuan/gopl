package main

import (
	"fmt"
)

func main() {
	fmt.Printf("\btab\n")
	fmt.Printf("\vvertical tab\n")
	fmt.Printf("\bbackspace\n")
	fmt.Printf("\nnewline\n")

	const GoUsage = `Go is a tool for managing Go source code.`
	fmt.Println(GoUsage)
}
