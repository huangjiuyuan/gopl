package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }

func main() {
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	p := Point{1, 2}
	pptr := &p
	pptr.ScaleBy(2)
	fmt.Println(p)

	p = Point{1, 2}
	(&p).ScaleBy(2)
	fmt.Println(p)

	// If the receiver p is a variable of type Point but the method requires a
	// *Point receiver, we can use this shorthand and the compiler will perform
	// an implicit &p on the variable. This works only for variables.
	p = Point{1, 2}
	p.ScaleBy(2)
	fmt.Println(p)

	// We cannot call a *Point method on a non-addressable Point receiver,
	// because there's no way to obtain the address of a temporary value.
	//
	//	 Point{1, 2}.ScaleBy(2)
	//
	// But we can call a Point method like Point.Distance with a *Point
	// receiver, because there is a way to obtain the value from the address:
	// just load the value pointed to by the receiver. The compiler inserts an
	// implicit * operation for us.
	q := Point{5, 8}
	fmt.Println(pptr.Distance(q))
	fmt.Println((*pptr).Distance(q))

	// In every valid method call expression, exactly one of these three
	// statements is true.
	// Either the receiver argument has the same type as the receiver
	// parameter, for example both have type T or both have type *T.
	Point{1, 2}.Distance(q) // Point
	pptr.ScaleBy(2)         // *Point
	// Or the receiver argument is a variable of type T and the receiver
	// parameter has type *T. The compiler implicitly takes the address of the
	// variable.
	p.ScaleBy(2) // implicit (&p)
	// Or the receiver argument has type *T and the receiver parameter has type
	// T. The compiler implicitly dereferences the receiver, in other words,
	// loads the value.
	pptr.Distance(q) // implicit (*pptr)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
