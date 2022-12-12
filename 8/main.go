package main

import (
	file "aoc-2022"
	"aoc-2022/8/puzzle2"
)

func main() {
	scanner := file.Scanner("8/input.txt")
	//answer, err := puzzle1.Solve(scanner)
	//if err != nil {
	//	panic(err)
	//}
	//println(*answer)

	answer2, err := puzzle2.Solve2(scanner)
	if err != nil {
		panic(err)
	}
	println(*answer2)
}
