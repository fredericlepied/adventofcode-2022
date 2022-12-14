package day12

import (
	"os"
	"testing"
)

func check(input []byte, want int, want2 int, t *testing.T) {
	got, got2 := ComputeResult(input)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
	if got2 != want2 {
		t.Errorf("part2: got %d, wanted %d", got2, want2)
	}
}

func TestComputeResult(t *testing.T) {
	// part 1
	small, _ := os.ReadFile("small.txt")
	input, _ := os.ReadFile("input.txt")
	check(small, 31, 29, t)
	check(input, 449, 443, t)
}
