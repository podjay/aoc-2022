package puzzle1

import (
	"aoc-2022/8/forest"
	"bufio"
)

func Solve(scanner *bufio.Scanner) (*int, error) {
	grid, err := forest.BuildForest(scanner)
	if err != nil {
		return nil, err
	}

	markPerimeterAsVisible(&grid)
	lookLeft(&grid)
	lookRight(&grid)
	lookDown(&grid)
	lookUp(&grid)
	visibleTrees := allVisibleTrees(&grid)
	return &visibleTrees, nil
}

func allVisibleTrees(i *[][]forest.Tree) int {
	count := 0
	for _, trees := range *i {
		for _, tree := range trees {
			if tree.IsVisible {
				count++
			}
		}
	}
	return count
}

func lookDown(pforest *[][]forest.Tree) {
	tallestTree := int64(-1)
	for i := 0; i < len(*pforest); i++ {
		for j := 0; j < len(*pforest); j++ {
			if (*pforest)[j][i].Height > tallestTree {
				(*pforest)[j][i].IsVisible = true
				tallestTree = (*pforest)[j][i].Height
			}
		}
		tallestTree = 0
	}
}

func lookUp(pforest *[][]forest.Tree) {
	tallestTree := int64(-1)
	for i := len(*pforest) - 1; i >= 0; i-- {
		for j := len(*pforest) - 1; j >= 0; j-- {
			if (*pforest)[j][i].Height > tallestTree {
				(*pforest)[j][i].IsVisible = true
				tallestTree = (*pforest)[j][i].Height
			}
		}
		tallestTree = 0
	}
}

func lookRight(pforest *[][]forest.Tree) {
	tallestTree := int64(-1)

	for i := len(*pforest) - 1; i >= 0; i-- {
		for j := 0; j < len(*pforest); j++ {
			if (*pforest)[i][j].Height > tallestTree {
				(*pforest)[i][j].IsVisible = true
				tallestTree = (*pforest)[i][j].Height
			}
		}
		tallestTree = 0
	}
}

func lookLeft(pforest *[][]forest.Tree) {
	tallestTree := int64(-1)
	for i := len(*pforest) - 1; i >= 0; i-- {
		for j := len(*pforest) - 1; j >= 0; j-- {
			if (*pforest)[i][j].Height > tallestTree {
				(*pforest)[i][j].IsVisible = true
				tallestTree = (*pforest)[i][j].Height
			}
		}
		tallestTree = 0
	}
}

func markPerimeterAsVisible(pforest *[][]forest.Tree) {
	forest := *pforest
	firstRow := forest[0][0:]
	markNode(&firstRow)
	lastRow := forest[len(forest)-1][0:]
	markNode(&lastRow)

	for i, _ := range forest {
		forest[i][0].IsVisible = true
		forest[i][len(forest)-1].IsVisible = true
	}
}

func markNode(trees *[]forest.Tree) {
	for i, _ := range *trees {
		(*trees)[i].IsVisible = true
	}
}
