package day14

import (
	"os"
	"testing"
)

func check(fname string, floor bool, want int, t *testing.T) {
	input, _ := os.ReadFile(fname)
	got := ComputeResult(input, floor)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestComputeResult(t *testing.T) {
	// part 1
	check("small.txt", false, 24, t)
	check("input.txt", false, 737, t)
	// part 2
	check("small.txt", true, 93, t)
	check("input.txt", true, 28145, t)
}
