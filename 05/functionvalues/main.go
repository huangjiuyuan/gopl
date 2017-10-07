package main

import (
	"fmt"
	"strings"
)

func main() {
	f := square
	fmt.Println(f(3))
	f = negative
	fmt.Println(f(3))
	fmt.Printf("%T\n", f)
	g := product
	fmt.Println(g(3, 4))
	fmt.Printf("%v, %v, %v\n", square, negative, product)

	// func can only be compared to nil.
	var h func(int) int
	fmt.Println(h == nil)

	// strings.Map applies a function to each character of a string, joining
	// the results to make another string.
	fmt.Println(strings.Map(addOne, "HAL-9000"))
	fmt.Println(strings.Map(addOne, "VMS"))
	fmt.Println(strings.Map(addOne, "Admix"))
}

func square(n int) int {
	return n * n
}

func negative(n int) int {
	return -n
}

func product(m, n int) int {
	return m * n
}

func addOne(r rune) rune {
	return r + 1
}
