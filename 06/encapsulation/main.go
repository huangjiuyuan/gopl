package main

// The Counter type permits clients to increment the counter or to reset it to
// zero, but not to set it to some arbitrary value.
type Counter struct{ n int }

func main() {
	c := Counter{0}
	c.Increment()
	c.Reset()
}

func (c *Counter) N() int {
	return c.n
}
func (c *Counter) Increment() {
	c.n++
}
func (c *Counter) Reset() {
	c.n = 0
}
