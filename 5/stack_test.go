package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStack_Pop(t *testing.T) {
	stack := NewStack()

	stack.Push("H")
	stack.Push("E")

	if pop, ok := stack.Pop(); ok {
		require.Equal(t, pop, "E")
	} else {
		require.FailNow(t, "pop failed")
	}

	if pop, ok := stack.Pop(); ok {
		require.Equal(t, pop, "H")
	} else {
		require.FailNow(t, "pop failed")
	}
}

func TestStack_Reverse(t *testing.T) {
	stack := NewStack()

	for _, c := range "HELLO" {
		stack.Push(string(c))
	}

	stack.Reverse()

	output := ""
	for i := 0; i < len(stack.dll); i++ {
		output += stack.dll[i]
	}
	require.Equal(t, "OLLEH", output)
}
