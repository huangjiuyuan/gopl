package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Println(x, y, z)

	const Pi64 float64 = math.Pi
	var a float32 = float32(Pi64)
	var b float64 = Pi64
	var c complex128 = complex128(Pi64)
	fmt.Println(a, b, c)

	var f float64 = 212
	fmt.Println((f - 32) * 5 / 9)
	fmt.Println(5 / 9 * (f - 32))
	fmt.Println(5.0 / 9.0 * (f - 32))

	var g float64 = 3 + 0i
	fmt.Println(g, reflect.TypeOf(g))
	g = 2
	fmt.Println(g, reflect.TypeOf(g))
	g = 1e123
	fmt.Println(g, reflect.TypeOf(g))
	g = 'a'
	fmt.Println(g, reflect.TypeOf(g))

	fmt.Printf("%T\n", 0)
	fmt.Printf("%T\n", 0.0)
	fmt.Printf("%T\n", 0i)
	fmt.Printf("%T\n", '\000')
}
