package decode_test

import (
	"aoc-2022/6/decode"
	"testing"
)

func TestFind(t *testing.T) {

	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "input 1",
			args: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want: 5,
		},
		{
			name: "input 2",
			args: "nppdvjthqldpwncqszvftbrmjlhg",
			want: 6,
		},
		{
			name: "input 3",
			args: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want: 10,
		},
		{
			name: "input 4",
			args: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decode.Packet(tt.args); got != tt.want {
				t.Errorf("Packet() = %v, want %v", got, tt.want)
			}
		})
	}
}
