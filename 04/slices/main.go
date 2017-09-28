package main

import (
	"fmt"
)

func main() {
	nums := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := nums[2:5]
	s2 := nums[4:7]
	fmt.Println(s1, s2)
	nums[4] = 1024
	fmt.Println(s1, s2)

	s3 := s2[:5]
	fmt.Println(s3)

	v1 := "string"
	v2 := v1[2:6]
	fmt.Println(v1, v2)
	v1 = "golang"
	fmt.Println(v1, v2)

	b1 := "string"
	b2 := b1[2:6]
	fmt.Println(b1, b2)
	b1 = "golang"
	fmt.Println(b1, b2)

	a := [...]int{0, 1, 2, 3, 4, 5}
	reverse(a[:])
	fmt.Println(a)

	fmt.Println(equal(s1, s2))

	var s []int
	fmt.Println(len(s), s == nil)
	s = nil
	fmt.Println(len(s), s == nil)
	s = []int(nil)
	fmt.Println(len(s), s == nil)
	s = []int{}
	fmt.Println(len(s), s == nil)

	s = make([]int, 4)
	fmt.Println(len(s), cap(s))
	s = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(len(s), cap(s))

	s = make([]int, 4, 6)
	fmt.Println(len(s), cap(s))
	s = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(len(s), cap(s))
}

// Since a slice contains a pointer to an element of an array, passing a slice
// to a function permits the function to modify the underlying array elements.
// In other words, copying a slice creates an alias for the underlying array.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// Unlike arrays, slices are not comparable, so we cannot use == to test whether
// two slices contain the same elements. We must do the comparison ourselves.
func equal(x, y []int) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
