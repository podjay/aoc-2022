package decode_test

import (
	"aoc-2022/6/decode"
	"testing"
)

func TestMessage(t *testing.T) {

	tests := []struct {
		name string
		args string
		want int
	}{
		{
			name: "input 1",
			args: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want: 19,
		},
		{
			name: "input 2",
			args: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want: 23,
		},
		{
			name: "input 3",
			args: "nppdvjthqldpwncqszvftbrmjlhg",
			want: 23,
		},
		{
			name: "input 4",
			args: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want: 29,
		},
		{
			name: "input 5",
			args: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want: 26,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decode.Message(tt.args); got != tt.want {
				t.Errorf("Message() = %v, want %v", got, tt.want)
			}
		})
	}
}
