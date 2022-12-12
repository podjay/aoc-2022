package forest

import (
	"bufio"
	"strconv"
)

func BuildForest(scanner *bufio.Scanner) ([][]Tree, error) {
	forest := make([][]Tree, 0)
	row := int64(0)
	for scanner.Scan() {
		text := scanner.Text()

		rowHeights := make([]Tree, 0)
		forest = append(forest, rowHeights)
		for _, value := range text {
			height, err := strconv.ParseInt(string(value), 10, 64)
			if err != nil {
				return nil, err
			}
			forest[row] = append(forest[row], Tree{Height: height})
		}
		row++
	}
	return forest, nil
}
