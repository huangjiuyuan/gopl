package main

import (
	"fmt"
)

func main() {
	i := 0
	b := itob(i)
	fmt.Println(b)
}

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func itob(i int) bool {
	return i != 0
}
