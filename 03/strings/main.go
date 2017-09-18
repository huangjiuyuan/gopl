package main

import "fmt"

func main() {
	s := "hello, world"
	s = "hello from the other side"
	fmt.Println(len(s))
	fmt.Println(s[0], s[7])

	fmt.Println(s[0:5])
	fmt.Println(s[:5])
	fmt.Println(s[6:])
	fmt.Println(s[:])

	fmt.Println("goodbye" + s[5:])

	l := "left foot"
	t := l
	l += ", right foot"
	fmt.Println(l)
	fmt.Println(t)
}
