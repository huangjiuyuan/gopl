package main

import "fmt"

type Point struct {
	X, Y int
}

type address struct {
	hostname string
	port     int
}

func main() {
	// If all the fields of a struct are comparable, the struct itself is
	// comparable, so two expressions of that type may be compared using "=="
	// or "!=".
	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y)
	fmt.Println(p == q)

	// Comparable struct types, like other comparable types, may be used as the
	// key type of a map.
	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	fmt.Println(hits)
}
