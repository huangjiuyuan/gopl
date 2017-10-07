package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(hypot(3, 4))

	// The type of a function is sometimes called its signature. Two functions
	// have the same type or signature if they have the same sequence of
	// parameter types and the same sequence of result types.
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", first)
	fmt.Printf("%T\n", zero)
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) (z int) {
	z = x - y
	return
}

func first(x int, _ int) int {
	return x
}

func zero(int, int) int {
	return 0
}
