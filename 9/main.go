package main

import (
	file "aoc-2022"
	"aoc-2022/9/move"
	"aoc-2022/9/simulation"
)

func main() {
	moves, err := move.Parse(file.Scanner("9/input.txt"))
	if err != nil {
		panic(err)
	}

	run := simulation.Run(moves)
	println(run)
}
