package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func main() {
	names := []string{"Lily", "Chris", "Sandy", "Ben"}
	sort.Sort(StringSlice(names))
	fmt.Println(names)
}

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
