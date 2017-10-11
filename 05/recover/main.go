package main

import (
	"fmt"
)

func main() {
	Parse("Hello")
	soleTitle()
}

func Parse(input string) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()
	return nil
}

func soleTitle() (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
		// no panic
		case bailout{}:
			// "expected" panic
			err = fmt.Errorf("multiple title elements")
		default:
			// unexpected panic; carry on panicking
			panic(p)
		}
	}()
	return title, nil
}
