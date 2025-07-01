package operations

type BinaryOp func(int, int) int

const (
	OpSum      = "sum"
	OpMultiply = "multiply"
)

var Registry = map[string]BinaryOp{
	OpSum:      func(a, b int) int { return a + b },
	OpMultiply: func(a, b int) int { return a * b },
	///divide, mod ...
}
