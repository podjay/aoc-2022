package decode

import (
	"strings"
)

func Packet(signal string) int {
	array := strings.Split(signal, "")
	for i := 0; i < len(array)-4; i++ {
		first := array[i]
		second := array[i+1]
		third := array[i+2]
		fourth := array[i+3]
		if first == second ||
			first == third ||
			first == fourth ||
			second == third ||
			second == fourth ||
			third == fourth {
			continue
		} else {
			return i + 4
		}
	}

	return 0
}
