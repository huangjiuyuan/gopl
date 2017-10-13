package main

import (
	"fmt"
	"math"
	"time"
)

type Point struct{ X, Y float64 }

type Path []Point

type Rocket struct{ Name string }

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	distanceFromP := p.Distance // method value
	fmt.Println(distanceFromP(q))
	var origin Point
	fmt.Println(distanceFromP(origin))
	scaleP := p.ScaleBy // method value
	scaleP(2)
	scaleP(3)
	scaleP(10)
	fmt.Println(p)

	r := new(Rocket)
	r.Name = "Very Fast"
	time.AfterFunc(1*time.Second, r.Launch)
	time.Sleep(2 * time.Second)

	// Related to the method value is the method expression. When calling a
	// method, as opposed to an ordinary function, we must supply the receiver
	// in a special way using the selector syntax. A method expression, written
	// T.f or (*T).f where T is a type, yields a function value with a regular
	// first parameter taking the place of the receiver.
	distance := Point.Distance   // method expression
	fmt.Println(distance(p, q))  // "5"
	fmt.Printf("%T\n", distance) // "func(Point, Point) float64"
	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (r *Rocket) Launch() {
	fmt.Println(r.Name)
}

func (p Point) Add(q Point) Point {
	return Point{p.X + q.X, p.Y + q.Y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.X - q.X, p.Y - q.Y}
}

func (path Path) TranslateBy(offset Point, add bool) {
	var op func(p, q Point) Point
	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}
	for i := range path {
		// Call either path[i].Add(offset) or path[i].Sub(offset).
		path[i] = op(path[i], offset)
	}
}
