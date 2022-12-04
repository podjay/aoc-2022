package file

import (
	"bufio"
	"os"
)

func Scanner(path string) *bufio.Scanner {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	return scanner
}
