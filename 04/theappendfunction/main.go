package main

import "fmt"

type IntSlice struct {
	ptr      *int
	len, cap int
}

func main() {
	var runes []rune
	for _, r := range "Hello, world" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)

	fmt.Println([]rune("Hello, world"))

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	a := []int{0, 1, 2, 3, 4, 5}
	b := make([]int, 4, 8)
	copy(b, a)
	fmt.Println(a, b)
	a[0] = 100
	fmt.Println(a, b)

	var v []int
	v = append(v, 1)
	v = append(v, 2, 3)
	v = append(v, 4, 5, 6)
	v = append(v, v...)
	fmt.Println(v)
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
