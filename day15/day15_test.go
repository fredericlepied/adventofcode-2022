package day15

import (
	u "github.com/fredericlepied/adventofcode-2022/utils"
	"testing"
)

func check(fname string, y int, want int, t *testing.T, f func([]string, int) int) {
	input := u.ReadFile(fname)
	got := f(input, y)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestComputePart1(t *testing.T) {
	// part 1
	check("small.txt", 10, 26, t, ComputeResult)
	check("input.txt", 2000000, 4737443, t, ComputeResult)
}

func TestComputePart2(t *testing.T) {
	// part 2
	check("small.txt", 20, 56000011, t, ComputeResult2)
	check("input.txt", 4000000, 11482462818989, t, ComputeResult2)
}
