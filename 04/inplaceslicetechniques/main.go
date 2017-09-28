package main

import "fmt"

func main() {
	data := []string{"one", "", "two", "three"}
	nonempty := nonempty1(data)
	fmt.Printf("%q %q\n", nonempty, nonempty[:4])
	fmt.Printf("%q\n", data)

	data = []string{"one", "", "two", "three"}
	nonempty = nonempty2(data)
	fmt.Printf("%q %q\n", nonempty, nonempty[:4])
	fmt.Printf("%q\n", data)

	stack := []int{0, 1, 2, 3, 4}
	stack = append(stack, 5)
	fmt.Println(stack)
	top := stack[len(stack)-1]
	fmt.Println(top)
	stack = stack[:len(stack)-1]
	fmt.Println(stack)

	s := []int{5, 6, 7, 8, 9}
	r1 := remove1(s, 2)
	fmt.Println(s, r1, r1[:5])
	s = []int{5, 6, 7, 8, 9}
	r2 := remove2(s, 2)
	fmt.Println(s, r2, r2[:5])
}

// nonempty1 returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func nonempty1(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func remove1(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
