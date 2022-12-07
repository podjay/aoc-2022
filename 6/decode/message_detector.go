package decode

import "strings"

var DistinctChars = 14

func Message(signal string) int {
	array := strings.Split(signal, "")

	for i := range array {
		m := make(map[string]bool)
		for j := 0; j < DistinctChars; j++ {
			s := array[i+j]
			if _, ok := m[s]; ok {
				break
			} else {
				m[s] = true
				if j == (DistinctChars - 1) {
					return i + j + 1
				}
			}
		}
	}
	return 0
}
