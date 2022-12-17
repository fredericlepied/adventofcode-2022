package day16

import (
	. "github.com/fredericlepied/adventofcode-2022/utils"
	"testing"
)

func check(fname string, want int, t *testing.T, f func([]string) int) {
	input := ReadFile(fname)
	got := f(input)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestComputePart1(t *testing.T) {
	check("small.txt", 1651, t, ComputeResult)
	//check("input.txt", 2000000, t, ComputeResult)
}
