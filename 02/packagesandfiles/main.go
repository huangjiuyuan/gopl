package main

import (
	"fmt"

	"github.com/huangjiuyuan/gopl/02/packagesandfiles/tempconv"
)

func main() {
	fmt.Printf("Brrrr! %v\n", tempconv.AbsoluteZeroC)
	fmt.Println(tempconv.CToF(tempconv.BoilingC))
}
