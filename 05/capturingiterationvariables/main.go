package main

import (
	"fmt"
)

func main() {
	ns := []int{0, 1, 2, 3, 4}

	var s1 []func()
	for _, v := range ns {
		val := v
		s1 = append(s1, func() {
			doSomething(val)
			fmt.Println(&val)
		})
	}
	for _, f := range s1 {
		f()
	}

	// The for loop introduces a new lexical block in which the variable v is
	// declared. All function values created by this loop capture and share the
	// same variable, an addressable storage location, not its value at that
	// particular moment. The value of v is updated in successive iterations,
	// so by the time the cleanup functions are called, the v variable has been
	// updated several times by the now-completed for loop. Thus v holds the
	// value from the final iteration, and consequently all calls to
	// doSomething will attempt to print the same number.
	var s2 []func()
	for _, v := range ns {
		s2 = append(s2, func() {
			doSomething(v)
			fmt.Println(&v)
		})
	}
	for _, f := range s2 {
		f()
	}

	// This will panic for index out of range. Because after the last execution
	// of for loop, the iteration variable increments to 5.
	var s3 []func()
	for i := 0; i < len(ns); i++ {
		s3 = append(s3, func() {
			doSomething(ns[i])
		})
	}
	for _, f := range s3 {
		f()
	}
}

func doSomething(i int) {
	fmt.Printf("%v ", i)
}
