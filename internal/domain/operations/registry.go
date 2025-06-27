package operations

type BinaryOp func(int, int) int

var Registry = map[string]BinaryOp{
	"sum":      func(a, b int) int { return a + b },
	"multiply": func(a, b int) int { return a * b },
	///divide, mod ...
}
