package main

import (
	"fmt"

	"github.com/huangjiuyuan/gopl/02/lifetimeofvariables/other"
)

var global *int

func main() {
	f()
	g()
	fmt.Println(*global, global)

	fmt.Println(other.Global)
	other.Global = 100
	fmt.Println(other.Global)
}

func f() {
	var x int
	x = 1
	global = &x
}

func g() {
	y := new(int)
	*y = 1
}
