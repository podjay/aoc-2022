package puzzle1

import (
	file "aoc-2022"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolve(t *testing.T) {
	scanner := file.Scanner("test_input.txt")

	answer, err := Solve(scanner)
	require.NoError(t, err)
	require.Equal(t, 21, *answer)
}
