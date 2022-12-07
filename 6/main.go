package main

import (
	file "aoc-2022"
	"aoc-2022/6/decode"
)

func main() {
	scanner := file.Scanner("6/input.txt")
	scanner.Scan()
	message := scanner.Text()
	number := decode.Packet(message)
	println(number)

	number = decode.Message(message)
	println(number)
}
