package main

import (
	file "aoc-2022"
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

var (
	pattern = regexp.MustCompile("move (\\d+) from (\\d+) to (\\d+)")
)

func main() {
	scanner := file.Scanner("5/input.txt")

	stacks := makeStacks(scanner)

	// Skip to moves
	scanner.Scan()
	scanner.Scan()

	moves := buildMoves(scanner)
	doMoves(moves, stacks)

	result := ""
	for _, stack := range stacks {
		pop, _ := stack.Pop()
		result += pop
	}
	fmt.Sprintf("Result: %s", result)
}

func doMoves(moves []Move, stacks []*Stack) {
	for _, move := range moves {
		tempStack := NewStack()
		for i := 0; i < move.number; i++ {
			pop, _ := stacks[move.from-1].Pop()
			tempStack.Push(pop)
		}

		for i := 0; i < move.number; i++ {
			pop, _ := tempStack.Pop()
			stacks[(move.to - 1)].Push(pop)
		}
	}
}

func buildMoves(scanner *bufio.Scanner) []Move {
	moves := make([]Move, 0)
	for scanner.Scan() {
		moves = append(moves, buildMove(scanner))
	}

	return moves
}

type Move struct {
	from   int
	to     int
	number int
}

func buildMove(scanner *bufio.Scanner) Move {
	matches := pattern.FindStringSubmatch(scanner.Text())
	from, err := strconv.ParseInt(matches[2], 10, 32)
	if err != nil {
		panic(err)
	}
	to, _ := strconv.ParseInt(matches[3], 10, 32)
	if err != nil {
		panic(err)
	}

	number, _ := strconv.ParseInt(matches[1], 10, 32)
	if err != nil {
		panic(err)
	}

	return Move{
		from:   int(from),
		to:     int(to),
		number: int(number),
	}
}

func makeStacks(scanner *bufio.Scanner) []*Stack {
	stacks := make([]*Stack, 9)
	for i, _ := range stacks {
		stacks[i] = NewStack()
	}

	for i := 0; i < 8; i++ {
		scanner.Scan()
		for i := 0; i <= 8; i++ {
			line := []rune(scanner.Text())
			requiredPosition := (i * 4) + 1
			if len(line) >= requiredPosition {
				letter := string(line[requiredPosition])
				if letter != " " {
					stacks[i].Push(letter)
				}
			}
		}
	}

	for i := range stacks {
		stacks[i].Reverse()
	}
	return stacks
}
