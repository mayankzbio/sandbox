package concurrent

type Adder struct{}

func (r Adder) Sum(a, b int) int {
	// Some thread safe operation
	return a + b
}
