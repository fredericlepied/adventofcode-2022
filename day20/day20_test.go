package day20

import (
	. "github.com/fredericlepied/adventofcode-2022/utils"
	"testing"
)

func check(fname string, want int, t *testing.T, f func([]string) int) {
	input := ReadFile(fname)

	got := f(input)

	if got != want {
		t.Errorf("%s got %d, wanted %d", fname, got, want)
	}
}

func TestPart1(t *testing.T) {
	check("small.txt", 3, t, ComputeResult)
	check("input.txt", 3, t, ComputeResult)
}
