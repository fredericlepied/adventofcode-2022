package day15

import (
	u "github.com/fredericlepied/adventofcode-2022/utils"
	"testing"
)

func check(fname string, y int, want int, t *testing.T) {
	input := u.ReadFile(fname)
	got := ComputeResult(input, y)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func check2(fname string, max int, want int, t *testing.T) {
	input := u.ReadFile(fname)
	got := ComputeResult2(input, max)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestComputePart1(t *testing.T) {
	// part 1
	check("small.txt", 10, 26, t)
	check("input.txt", 2000000, 4737443, t)
}

func TestComputePart2(t *testing.T) {
	// part 2
	check2("small.txt", 20, 56000011, t)
	check2("input.txt", 4000000, 11482462818989, t)
}
