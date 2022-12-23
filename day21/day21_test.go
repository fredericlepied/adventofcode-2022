package day21

import (
	. "github.com/fredericlepied/adventofcode-2022/utils"
	"testing"
)

func check(prefix string, fname string, want int, t *testing.T, f func([]string) int) {
	input := ReadFile(fname)

	got := f(input)

	if got != want {
		t.Errorf("%s(%s): got %d, wanted %d", prefix, fname, got, want)
	}
}

func TestComputePart1(t *testing.T) {
	check("day21p1", "small.txt", 152, t, ComputeResult)
	check("day21p1", "input.txt", 54703080378102, t, ComputeResult)
	//check("day21p2", "small.txt", 301, t, ComputeResult2)
}
