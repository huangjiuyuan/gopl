package main

import (
	"fmt"
	"strconv"
)

func main() {
	// To convert an integer to a string , one option is to use fmt.Sprintf;
	// another is to use the function strconv.Itoa.
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	// FormatInt and FormatUint can be used to format numbers in a different
	// base. The fmt.Printf verbs %b, %d, %u, and %x are often more convenient
	// than Format functions, especially if we want to include additional
	// information besides the number.
	fmt.Println(strconv.FormatInt(int64(x), 2))
	fmt.Println(fmt.Sprintf("x=%b", x))

	// To parse a string representing an integer, use the strconv functions
	// Atoi or ParseInt, or ParseUint for unsigned integers.
	a, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println(err)
	}
	b, err := strconv.ParseInt("123", 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a, b)
}
