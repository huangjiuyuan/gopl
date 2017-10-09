package main

import (
	"fmt"
	"os"
)

func main() {
	// The sum function returns the sum of zero or more int arguments. Within
	// the body of the function, the type of vals is an []int slice. When sum
	// is called, any number of values may be provided for its vals parameter.
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4))

	// Implicitly, the caller allocates an array, copies the arguments into it,
	// and passes a slice of the entire array to the function. To invoke a
	// variadic function when the arguments are already in a slice, place an
	// ellipsis after the final argument.
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...))

	// Although the ...int parameter behaves like a slice within the function
	// body, the type of a variadic function is distinct from the type of a
	// function with an ordinary slice parameter.
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)

	linenum, a, b, c := 12, "var", "func", "map"
	errorf(linenum, "undefined: %s %s %s", a, b, c)
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func f(...int) {}

func g([]int) {}

func errorf(linenum int, format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Line %d: ", linenum)
	fmt.Fprintf(os.Stderr, format, args...)
	fmt.Fprintln(os.Stderr)
}
