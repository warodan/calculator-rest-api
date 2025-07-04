package operations

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOperations(t *testing.T) {
	tests := []struct {
		name   string
		op     string
		a, b   int
		expect int
	}{
		{"sum of 2 and 3", OpSum, 2, 3, 5},
		{"multiply 4 and 5", OpMultiply, 4, 5, 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn, ok := Registry[tt.op]
			require.True(t, ok)
			require.Equal(t, tt.expect, fn(tt.a, tt.b))
		})
	}
}
