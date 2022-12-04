package main

import (
	file "aoc-2022"
	"strings"
)

func main() {
	scanner := file.Scanner("2022/3/3.txt")
	totalPriority := 0
	for scanner.Scan() {
		firstGroup := scanner.Text()
		scanner.Scan()
		secondGroup := scanner.Text()
		scanner.Scan()
		thirdGroup := scanner.Text()

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
