package concurrent

func Sum() {}

type Adder struct{}

func (r Adder) Sum(a, b int) int {
	// Some thread safe operation
	return a + b
}

type Arc struct{}

func (r Arc) Sum(a, b string) string {
	// Some thread safe operation
	return a + b
}
