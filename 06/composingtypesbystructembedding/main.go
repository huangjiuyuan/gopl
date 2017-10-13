package main

import (
	"fmt"
	"image/color"
	"math"
	"sync"
)

type Point struct{ X, Y float64 }

type ColoredPoint struct {
	Point
	Color color.RGBA
}

type PtrColoredPoint struct {
	*Point
	Color color.RGBA
}

func main() {
	// We can select the fields of ColoredPoint that were contributed by the
	// embedded Point without mentioning Point.
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	// We can call methods of the embedded	Point field using a receiver of
	// type ColoredPoint, even though ColoredPoint has no declared methods.
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point)) // "5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"
	// The methods of Point have been promoted to ColoredPoint. In this way,
	// embedding allows complex types with many methods to be built up by the
	// composition of several fields, each providing a few methods.
	// A ColoredPoint is not a Point, but it "has a" Point, and it has two
	// additional methods Distance and ScaleBy promoted from Point.

	// The type of an anonymous field may be a pointer to a named type, in
	// which case fields and methods are promoted indirectly from the
	// pointed-to object. Adding another level of indirection lets us share
	// common structures and vary the relation ships between objects
	// dynamically.
	pp := PtrColoredPoint{&Point{1, 1}, red}
	pq := PtrColoredPoint{&Point{5, 4}, blue}
	fmt.Println(pp.Distance(*pq.Point)) // "5"
	pq.Point = pp.Point
	// p and q now share the same Point
	pp.ScaleBy(2)
	fmt.Println(*pp.Point, *pq.Point) // "{2 2} {2 2}"
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

var (
	mu      sync.Mutex // guards mapping
	mapping = make(map[string]string)
)

func Lookup(key string) string {
	mu.Lock()
	v := mapping[key]
	mu.Unlock()
	return v
}

// The new variable gives more expressive names to the variables related to the
// cache, and because the sync.Mutex field is embedded within it, its Lock and
// Unlock methods are promoted to the unnamed struct type, allowing us to lock
// the cache with a self-explanatory syntax.
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func CacheLookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
