package day17

import (
	. "github.com/fredericlepied/adventofcode-2022/utils"
	"testing"
)

func check(fname string, want int, t *testing.T, f func(string) int) {
	input := ReadFile(fname)

	// remove trailing newline
	got := f(input[0])

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestComputePart1(t *testing.T) {
	check("small.txt", 3068, t, ComputeResult)
	//check("input.txt", 2000000, t, ComputeResult)
}
