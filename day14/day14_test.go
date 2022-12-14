package day14

import (
	"os"
	"testing"
)

func check(fname string, want int, t *testing.T) {
	input, _ := os.ReadFile(fname)
	got := ComputeResult(input)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestComputeResult(t *testing.T) {
	// part 1
	check("small.txt", 24, t)
	check("input.txt", 737, t)
}
