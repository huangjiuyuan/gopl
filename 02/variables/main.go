package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	var s1 string
	var i1, i2 int
	var b1, f1, s2 = true, 2.3, "four"
	fmt.Println(s1, i1, i2, b1, f1, s2)

	freq := rand.Float64() * 3.0
	t := 0.0
	fmt.Println(freq, t)

	i3 := 100
	var boiling float64 = 100
	var names []string
	var err error
	fmt.Println(i3, boiling, names, err)

	i4, i5 := 0, 1
	fmt.Println(i4, i5)

	i4, i5 = i5, i4
	fmt.Println(i4, i5)

	f, err := os.Open("main.go")
	if err != nil {
		fmt.Println(err)
	}
	f.Close()
}
