package main

import "fmt"

func main() {
	p := new(int)
	fmt.Println(*p, p)
	*p = 2
	fmt.Println(*p, p)

	i1 := newInt1()
	fmt.Println(*i1, i1)
	i2 := newInt2()
	fmt.Println(*i2, i2)

	i3 := new(int)
	i4 := new(int)
	fmt.Println(i3 == i4)
}

func newInt1() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}
