package main

import (
	"fmt"
	"sort"
)

type Graph map[string]map[string]bool

var m = make(map[string]int)

func main() {
	ages := map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Println(ages)

	ages = make(map[string]int)
	ages["alice"] = 31
	ages["charlie"] = 34
	fmt.Println(ages)

	ages["alice"] = 32
	fmt.Println(ages["alice"])

	delete(ages, "alice")
	fmt.Println(ages)

	// All of these operations are safe even if the element isn't in the map; a
	// map lookup using a key that isn't present returns the zero value for its
	// type.
	ages["bob"] = ages["bob"] + 1
	ages["bob"] += 1
	ages["bob"]++
	fmt.Println(ages)

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	// To enumerate the key/value pairs in order, we must sort	the keys
	// explicitly.
	names := make([]string, 0, len(ages))
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	var empty map[string]int
	fmt.Println(empty == nil)
	fmt.Println(len(empty) == 0)

	empty = make(map[string]int)
	fmt.Println(empty)
	fmt.Println(len(empty))
	empty["carol"] = 21
	fmt.Println(empty)

	age, ok := ages["bob"]
	if ok {
		fmt.Println(age)
	}

	if age, ok := ages["bob"]; ok {
		fmt.Println(age)
	}

	seen := make(map[string]bool)
	lines := []string{"q", "w", "e", "r", "t", "y", "q", "w", "e", "r", "q", "w"}
	for _, line := range lines {
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}

	var graph Graph
	graph = make(map[string]map[string]bool)
	graph.addEdge("a", "b")
	graph.addEdge("a", "c")
	graph.addEdge("b", "c")
	graph.addEdge("c", "d")
	fmt.Println(graph.hasEdge("a", "b"))
	fmt.Println(graph.hasEdge("b", "d"))
	fmt.Println(graph)
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}

func Count(list []string) int {
	return m[k(list)]
}

func (graph Graph) addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func (graph Graph) hasEdge(from, to string) bool {
	return graph[from][to]
}
