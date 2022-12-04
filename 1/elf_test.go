package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCanAddCalories(t *testing.T) {
	elf := NewElf()
	elf.AddCalorie(int64(10))
	require.Equal(t, int64(10), elf.GetCalorieCount())
}
