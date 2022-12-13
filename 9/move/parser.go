package move

import (
	"bufio"
	"strconv"
	"strings"
)

func Parse(scanner *bufio.Scanner) ([]It, error) {
	var moves []It

	for scanner.Scan() {
		text := scanner.Text()
		split := strings.Split(text, " ")
		direction := fromString(split[0])
		height, err := strconv.ParseInt(split[1], 10, 64)
		if err != nil {
			return nil, err
		}
		moves = append(moves, It{
			Direction: direction,
			Distance:  int(height),
		})
	}

	return moves, nil
}

func fromString(s string) Direction {
	switch s {
	case "U":
		return Up
	case "D":
		return Down
	case "L":
		return Left
	case "R":
		return Right
	default:
		panic("No direction found")
	}
}
