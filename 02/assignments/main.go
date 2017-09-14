package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	var x int
	var p bool
	person := person{}
	count := []int{0, 1, 2, 3, 4}

	x = 1
	p = true
	person.name = "bob"
	count[x] = count[x] * 5
	fmt.Println(x, p, person, count[x])

	count[x] *= 5
	fmt.Println(count[x])

	v := 1
	v++
	fmt.Println(v)
	v--
	fmt.Println(v)
}
