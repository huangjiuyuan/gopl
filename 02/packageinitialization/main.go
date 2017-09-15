package main

import (
	"fmt"

	"github.com/huangjiuyuan/gopl/02/packageinitialization/popcount"
)

var a = b + c
var b = f()
var c = 1

func main() {
	fmt.Println(a, b, c)
	fmt.Println(popcount.PopCount(255))
}

func f() int {
	return c + 1
}
