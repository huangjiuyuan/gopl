package main

import (
	"fmt"
)

func main() {
	x, y := 1, 2
	a := []int{0, 1, 2, 3, 4}

	x, y = y, x
	a[0], a[1] = a[1], a[0]
	fmt.Println(x, y, a)
	fmt.Println(gcd(24, 36))
	fmt.Println(fib(8))
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
