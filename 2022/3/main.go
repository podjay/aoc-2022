package main

import (
	file "aoc-2022"
	"strings"
)

// Lowercase item types a through z have priorities 1 through 26.
// Uppercase item types A through Z have priorities 27 through 52.

var priorities = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,

	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

func main() {
	scanner := file.Scanner("2022/3/3.txt")
	totalPriority := 0
	for scanner.Scan() {
		firstGroup := scanner.Text()
		scanner.Scan()
		secondGroup := scanner.Text()
		scanner.Scan()
		thirdGroup := scanner.Text()
		//left, right := halfString(text)
		commonCharacter := findCommonCharacter(firstGroup, secondGroup, thirdGroup)

		totalPriority += priorities[commonCharacter]
	}
	println(totalPriority)
}

func findCommonCharacter(first string, second string, third string) string {
	for _, char := range first {
		firstSubStr := string(char)
		if strings.Contains(second, firstSubStr) {
			if strings.Contains(third, firstSubStr) {
				return firstSubStr
			}
		}
	}
	panic("no match found")
}
