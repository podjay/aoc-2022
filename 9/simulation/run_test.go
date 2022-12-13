package simulation

import (
	"aoc-2022/9/move"
	"github.com/stretchr/testify/require"
	"testing"
)

// R 4
// U 4
// L 3
// D 1
// R 4
// D 1
// L 5
// R 2
func TestSimulation(t *testing.T) {
	moves := []move.It{
		{Direction: move.Right, Distance: 4},
		{Direction: move.Up, Distance: 4},
		{Direction: move.Left, Distance: 3},
		{Direction: move.Down, Distance: 1},
		{Direction: move.Right, Distance: 4},
		{Direction: move.Down, Distance: 1},
		{Direction: move.Left, Distance: 5},
		{Direction: move.Right, Distance: 2}}

	numberOfMoves := Run(moves)
	require.Equal(t, 13, numberOfMoves)
}

func TestSimulationDiagonal(t *testing.T) {
	moves := []move.It{
		{Direction: move.Right, Distance: 1},
		{Direction: move.Up, Distance: 1},
		{Direction: move.Right, Distance: 1},
		{Direction: move.Up, Distance: 1},
		{Direction: move.Right, Distance: 1},
		{Direction: move.Up, Distance: 1},
		{Direction: move.Right, Distance: 1},
		{Direction: move.Up, Distance: 1}}

	numberOfMoves := Run(moves)
	require.Equal(t, 13, numberOfMoves)
}
