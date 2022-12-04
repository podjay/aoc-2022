package main

import (
	file "aoc-2022"
	"github.com/bits-and-blooms/bitset"
	"strconv"
	"strings"
)

func main() {
	scanner := file.Scanner("4/input.txt")
	fullyContained := 0
	for scanner.Scan() {
		pair := strings.Split(scanner.Text(), ",")
		firstSectionRange := toBoundary(pair[0])
		secondSectionRange := toBoundary(pair[1])

		if intersection := firstSectionRange.Union(secondSectionRange); intersection.Equal(firstSectionRange) || intersection.Equal(secondSectionRange) {
			fullyContained++
		}
	}
	println(fullyContained)
}

func toBoundary(pair string) *bitset.BitSet {
	bounds := strings.Split(pair, "-")
	lowerBound, _ := strconv.ParseInt(bounds[0], 10, 64)
	upperBound, _ := strconv.ParseInt(bounds[1], 10, 64)

	b := bitset.New(100) // assumed max length
	for i := lowerBound; i <= upperBound; i++ {
		b.Set(uint(i))
	}
	return b

}
