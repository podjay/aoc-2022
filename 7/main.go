package main

import (
	file "aoc-2022"
	"aoc-2022/7/puzzle"
	"errors"
	"fmt"
	"strconv"
)

func main() {
	partOne()
	partTwo()
}

func partTwo() {
	scanner := file.Scanner("7/input.txt")

	lines := make([]puzzle.Line, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if matches, ok := puzzle.ParseCommand(line); ok {
			// we have a command
			if matches[1] == "cd" {
				lines = append(lines, puzzle.NewChangeDirectory(matches[2]))
			}

			if (matches[1]) == "ls" {
				lines = append(lines, puzzle.NewListFiles())
			}
		} else {
			// it's a dir or file
			if matches, ok := puzzle.ParseDir(line); ok {
				lines = append(lines, puzzle.NewDirInput(matches[1]))
			} else {
				if matches, ok := puzzle.ParseFile(line); ok {
					fileSize, err := strconv.ParseInt(matches[1], 10, 64)
					if err != nil {
						panic(err)
					}
					filename := matches[2]
					lines = append(lines, puzzle.NewFileInput(filename, fileSize))

				} else {
					panic(errors.New(fmt.Sprintf("found a line that I could not map to dir or file: %s", line)))
				}
			}
		}
	}

	tree := puzzle.NewDirectory("/", nil)
	for _, line := range lines {
		var err error
		tree, err = line.ApplyInputToDirectory(tree)
		if err != nil {
			panic(err)
		}
	}

	root := tree.Root()

	// total space available = 70000000
	// required space = 30000000

	diskSpace := int64(70000000)
	spaceUsed := root.Solve(diskSpace, nil)

	spaceAvailable := diskSpace - spaceUsed

	requiredSpace := int64(30000000)

	spaceToFreeUp := requiredSpace - spaceAvailable

	var dirsWithMaxSize []puzzle.Directory
	root.Solve(diskSpace, &dirsWithMaxSize)

	smallest := diskSpace
	for _, directory := range dirsWithMaxSize {
		size := directory.Solve(diskSpace, nil)

		if size < smallest && size >= spaceToFreeUp {
			smallest = size
		}
	}
	println(smallest)
}

func partOne() {
	scanner := file.Scanner("7/input.txt")

	lines := make([]puzzle.Line, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if matches, ok := puzzle.ParseCommand(line); ok {
			// we have a command
			if matches[1] == "cd" {
				lines = append(lines, puzzle.NewChangeDirectory(matches[2]))
			}

			if (matches[1]) == "ls" {
				lines = append(lines, puzzle.NewListFiles())
			}
		} else {
			// it's a dir or file
			if matches, ok := puzzle.ParseDir(line); ok {
				lines = append(lines, puzzle.NewDirInput(matches[1]))
			} else {
				if matches, ok := puzzle.ParseFile(line); ok {
					fileSize, err := strconv.ParseInt(matches[1], 10, 64)
					if err != nil {
						panic(err)
					}
					filename := matches[2]
					lines = append(lines, puzzle.NewFileInput(filename, fileSize))

				} else {
					panic(errors.New(fmt.Sprintf("found a line that I could not map to dir or file: %s", line)))
				}
			}
		}
	}

	tree := puzzle.NewDirectory("/", nil)
	for _, line := range lines {
		var err error
		tree, err = line.ApplyInputToDirectory(tree)
		if err != nil {
			panic(err)
		}
	}

	root := tree.Root()

	var dirsWithMaxSize []puzzle.Directory
	root.Solve(100000, &dirsWithMaxSize)

	totalSize := int64(0)
	for _, directory := range dirsWithMaxSize {
		totalSize += directory.Solve(100000, nil)
	}
}
