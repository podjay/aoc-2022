package main

import (
	file "aoc-2022"
	"fmt"
	"sort"
	"strconv"
)

func main() {
	scanner := file.Scanner("2022/1/1.txt")

	elves := make([]Elf, 0)
	var elf = NewElf()
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			elves = append(elves, elf)
			elf = NewElf()
		} else {
			parseCalorie, err := strconv.ParseInt(text, 10, 64)
			if err != nil {
				panic(err)
			}
			elf.AddCalorie(parseCalorie)
		}
	}

	// Sort by age, keeping original order or equal elements.
	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i].GetCalorieCount() > elves[j].GetCalorieCount()
	})

	println(fmt.Sprintf("Top 3: %d", elves[0].GetCalorieCount()+elves[1].GetCalorieCount()+elves[2].GetCalorieCount()))
}
