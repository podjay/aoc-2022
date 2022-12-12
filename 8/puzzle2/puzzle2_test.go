package puzzle2

import (
	file "aoc-2022"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSolve2(t *testing.T) {
	scanner := file.Scanner("../test_input.txt")

	answer, err := Solve2(scanner)
	require.NoError(t, err)
	require.Equal(t, int64(8), *answer)
}
