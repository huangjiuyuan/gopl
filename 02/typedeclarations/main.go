package main

import (
	"fmt"

	"github.com/huangjiuyuan/gopl/02/typedeclarations/tempconv"
)

func main() {
	absoluteZeroF := tempconv.CToF(tempconv.AbsoluteZeroC)
	freezingF := tempconv.CToF(tempconv.FreezingC)
	boilingF := tempconv.CToF(tempconv.BoilingC)
	fmt.Println(absoluteZeroF, freezingF, boilingF)

	// A conversion from one type to another is allowed if both have the same
	// underlying type, or if both are unnamed pointer types that point to
	// variables of the same underlying type; these conversions change the type
	// but not the representation of the value. If x is assignable to T, a
	// conversion is permitted but is usually redundant.
	converted := tempconv.Celsius(boilingF)
	fmt.Println(converted)

	// The underlying type of a named type determines its structure and
	// representation, and also the set of intrinsic operations it supports,
	// which are the same as if the underlying type had been used directly. That
	// means that arithmetic operators work the same for Celsius and Fahrenheit
	// as they do for float64.
	fmt.Printf("%g\n", tempconv.BoilingC-tempconv.FreezingC)
	fmt.Printf("%g\n", boilingF-freezingF)

	// Comparison operators like == and < can also be used to compare a value
	// of a named type to another of the same type, or to a value of the
	// underlying type.
	var c tempconv.Celsius
	var f tempconv.Fahrenheit
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	fmt.Println(c == tempconv.Celsius(f))

	t := tempconv.FToC(212.0)
	fmt.Println(t.String())
	fmt.Printf("%v\n", t)
	fmt.Printf("%s\n", t)
	fmt.Println(t)
	fmt.Printf("%g\n", t)
	fmt.Println(float64(t))
}
