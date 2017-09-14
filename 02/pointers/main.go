package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	x := 1
	p := &x
	fmt.Println(*p)
	*p = 2
	fmt.Println(x)

	var a, b int
	fmt.Println(&a == &a, &a == &b, &a == nil)

	var f = intPtr()
	fmt.Println(f, intPtr() == intPtr())

	v := 1
	incr(&v)
	fmt.Println(incr(&v))

	var n = flag.Bool("n", false, "omit trailing newline")
	var sep = flag.String("s", " ", "separator")
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

func intPtr() *int {
	v := 1
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
