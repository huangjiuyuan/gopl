package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	var t [3]int
	fmt.Println(t[0])
	fmt.Println(t[len(t)-1])

	// Print the indices and elements.
	for i, v := range t {
		fmt.Printf("%d %d\n", i, v)
	}
	// Print the elements only.
	for _, v := range t {
		fmt.Printf("%d\n", v)
	}

	var q [3]int = [3]int{1, 2, 3}
	var r [3]int = [3]int{1, 2}
	fmt.Println(q[2], r[2])

	// In an array literal, if an ellipsis "..." appears in place of the
	// length, the array length is determined by the number of initializers.
	p := [...]int{1, 2, 3, 4}
	fmt.Printf("%T\n", p)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c)
	d := [3]int{1, 2}
	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(d))

	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	var p1 *[8]byte = &[8]byte{}
	var p2 *[8]byte = &[8]byte{}
	zero1(p1)
	zero2(p2)
	fmt.Println(p1)
	fmt.Println(p2)

	checksum()
}

func zero1(ptr *[8]byte) {
	for i := range ptr {
		ptr[i] = 1
	}
}

func zero2(ptr *[8]byte) {
	*ptr = [8]byte{2, 2, 2, 2}
}

func checksum() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter 1 argument")
		return
	}
	if os.Args[1] == strconv.Itoa(384) {
		fmt.Printf("%x\n", sha512.Sum512([]byte("a")))
	} else if os.Args[1] == strconv.Itoa(512) {
		fmt.Printf("%x\n", sha512.Sum384([]byte("a")))
	} else {
		fmt.Printf("%x\n", sha256.Sum256([]byte("a")))
	}
}
