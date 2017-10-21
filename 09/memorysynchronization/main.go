package main

import (
	"fmt"
	"time"
)

func main() {
	var x, y int
	go func() {
		x = 1                // A1
		fmt.Println("y:", y) // A2
	}()
	go func() {
		y = 1                // B1
		fmt.Println("x:", x) // B2
	}()
	time.Sleep(500 * time.Millisecond)
}
