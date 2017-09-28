package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w Wheel
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
	fmt.Println(w)

	var u Wheel
	u.Circle.Point.X = 8
	u.Circle.Point.Y = 8
	u.Circle.Radius = 5
	u.Spokes = 20
	fmt.Println(u)

	v := Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20,
	}
	fmt.Printf("%#v\n", v)
	v.X = 42
	fmt.Printf("%#v\n", v)

	// The outer struct type gains not just the fields of the embedded type
	// but its methods too. This mechanism is the main way that complex object
	// behaviors are composed from simpler ones. Composition is central to
	// object-oriented programming in Go.
	fmt.Println(v.GetPosition())
	fmt.Println(v.GetRadius())
}

func (p Point) GetPosition() (int, int) {
	return p.X, p.Y
}

func (c Circle) GetRadius() int {
	return c.Radius
}
