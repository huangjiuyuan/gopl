package main

import (
	"fmt"
)

func main() {
	medals := []string{"gold", "silver", "bronze"}
	medals[0] = "Gold"
	medals[1] = "Silver"
	medals[2] = "Bronze"
	fmt.Println(medals)
}
