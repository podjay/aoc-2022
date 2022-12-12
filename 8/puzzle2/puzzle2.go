package puzzle2

import (
	"aoc-2022/8/forest"
	"bufio"
)

func Solve2(scanner *bufio.Scanner) (*int64, error) {
	grid, err := forest.BuildForest(scanner)
	if err != nil {
		return nil, err
	}

	topScore := int64(0)
	for i, trees := range grid {
		for j, _ := range trees {
			numVisibleDownCount := numVisibleDown(i, j, grid)
			numVisibleUpCount := numVisibleUp(i, j, grid)
			numVisibleLeftCount := numVisibleLeft(i, j, grid)
			numVisibleRightCount := numVisibleRight(i, j, grid)

			coefficient := int64(numVisibleDownCount * numVisibleUpCount * numVisibleLeftCount * numVisibleRightCount)
			if coefficient > topScore {
				topScore = coefficient
			}
		}
	}

	return &topScore, nil
}

func numVisibleDown(i int, j int, forest [][]forest.Tree) int {
	if i == len(forest)-1 {
		return 0
	}
	count := 0
	for x := i; x < len(forest)-1; x++ {
		count++
		treeOne := x + 1
		campHeight := forest[i][j].Height
		viewHeight := int64(0)
		neighbouringHeight := forest[treeOne][j].Height
		if neighbouringHeight >= campHeight || viewHeight >= campHeight {
			return count
		}
		viewHeight = neighbouringHeight
	}
	return count
}

func numVisibleRight(i int, j int, forest [][]forest.Tree) int {
	if j == len(forest)-1 {
		return 0
	}
	count := 0
	for y := j; y < len(forest)-1; y++ {
		count++
		treeOne := y + 1
		campHeight := forest[i][j].Height
		viewHeight := int64(0)
		neighbouringHeight := forest[i][treeOne].Height
		if neighbouringHeight >= campHeight || viewHeight >= campHeight {
			return count
		}
		viewHeight = neighbouringHeight
	}
	return count
}

func numVisibleUp(i int, j int, forest [][]forest.Tree) int {
	if i == 0 {
		return 0
	}
	count := 0
	for x := i; x >= 1; x-- {
		count++
		treeOne := x - 1
		campHeight := forest[i][j].Height
		viewHeight := int64(0)
		neighbouringHeight := forest[treeOne][j].Height
		if neighbouringHeight >= campHeight || viewHeight >= campHeight {
			return count
		}
	}
	return count
}

func numVisibleLeft(i int, j int, forest [][]forest.Tree) int {
	if j == 0 {
		return 0
	}
	count := 0
	for y := j; y >= 1; y-- {
		count++
		treeOne := y - 1
		campHeight := forest[i][j].Height
		viewHeight := int64(0)
		neighbouringHeight := forest[i][treeOne].Height
		if treeOne < 0 || neighbouringHeight >= campHeight || viewHeight >= campHeight {
			return count
		}
	}
	return count
}
