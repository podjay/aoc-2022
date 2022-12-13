package simulation

import (
	"aoc-2022/9/move"
	"math"
)

func Run(moves []move.It) int {

	grid := [1000][1000]bool{}
	tail := move.Position{X: 500, Y: 500}
	head := move.Position{X: 500, Y: 500}
	for _, it := range moves {
		for i := 0; i < it.Distance; i++ {
			headLastPos := head
			switch it.Direction {
			case move.Up:
				head.Y++
			case move.Down:
				head.Y--
			case move.Left:
				head.X--
			case move.Right:
				head.X++
			}
			tail = moveTail(head, tail, headLastPos)
			grid = visit(tail, grid)
		}
	}

	return countVisits(grid)
	// 6877 -- too high
	// 6310 -- just right
	// 6310 -- too low
}

func countVisits(grid [1000][1000]bool) int {
	visits := 0
	for _, bools := range grid {
		for _, b := range bools {
			if b {
				visits++
			}
		}
	}
	return visits
}

func moveTail(head move.Position, tail move.Position, pos move.Position) move.Position {

	if distance(head, tail) > math.Sqrt(2) {
		tail.X = pos.X
		tail.Y = pos.Y
	}

	return tail
}

func distance(head move.Position, tail move.Position) float64 {
	return math.Sqrt(float64((tail.X-head.X)*(tail.X-head.X) + (tail.Y-head.Y)*(tail.Y-head.Y)))
}

func visit(tail move.Position, grid [1000][1000]bool) [1000][1000]bool {
	grid[tail.X][tail.Y] = true
	return grid
}
