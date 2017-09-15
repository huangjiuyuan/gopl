package main

import (
	"fmt"
	"log"
	"os"
)

func f() {}

func g() int {
	return 100
}

func h(s int) int {
	return s
}

var v = "v"

var cwd string

func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	log.Printf("Working directory = %s", cwd)
}

func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}

func main() {
	f := "f"
	// "f"; local var f shadows package-level func f
	fmt.Println(f)
	// "v"; package-level var
	fmt.Println(v)

	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c\n", x)
		}
	}

	x = "yeah"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c\n", x)
	}

	// The second if statement is nested within the first, so variables
	// declared within the first statement's initializer are visible within
	// the second.
	if a := g(); a == 0 {
		fmt.Println(a)
	} else if b := h(a); a == b {
		fmt.Println(a, b)
	} else {
		fmt.Println(a, b)
	}

	if f, err := os.Open("error"); err != nil {
		fmt.Println(err)
	} else {
		// f and err are visible here too.
		f.Name()
		f.Close()
	}

	fmt.Println(cwd)
}
