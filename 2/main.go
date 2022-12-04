package main

import (
	file "aoc-2022"
	"strings"
)

var lookup = map[string]Action{
	"A": Rock{},
	"B": Paper{},
	"C": Scissors{},
}

var requiredOutcomes = map[string]Outcome{
	"X": Loss{},
	"Y": Draw{},
	"Z": Win{},
}

func main() {
	scanner := file.Scanner("2/2.txt")
	score := 0
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		opponentAction := lookup[text[0]]
		requiredOutcome := requiredOutcomes[text[1]]
		requiredAction := opponentAction.RequiredActionFor(requiredOutcome)
		score += requiredAction.Points() + requiredOutcome.Points()
	}
	println(score)
}
