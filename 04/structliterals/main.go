package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func main() {
	p := Point{1, 2}
	fmt.Println(p)

	fmt.Println(Scale(Point{1, 2}, 5))

	pp1 := &Point{1, 2}
	fmt.Println(pp1)

	pp2 := new(Point)
	*pp2 = Point{1, 2}
	fmt.Println(pp2)
}

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}
