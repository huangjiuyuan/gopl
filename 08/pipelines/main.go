package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	// Equivalent to:
	// 	go func() {
	// 		for {
	// 			x, ok := <-naturals
	// 			if !ok {
	// 				break
	// 			}
	// 			squares <- x * x
	// 		}
	// 		close(squares)
	// 	}()
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	// Equivalent to:
	// 	for {
	// 		x, ok := <-squares
	// 		if !ok {
	// 			break
	// 		}
	// 		fmt.Println(x)
	// 	}
	for x := range squares {
		fmt.Println(x)
	}
}
