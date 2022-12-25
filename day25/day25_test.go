package day25

import (
	. "github.com/fredericlepied/adventofcode-2022/utils"
	"testing"
)

func check(fname string, want string, t *testing.T, f func([]string) string) {
	input := ReadFile(fname)

	got := f(input)

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

type value struct {
	num   int
	snafu string
}

func TestPart1(t *testing.T) {
	var table = []value{
		{1, "1"},
		{2, "2"},
		{3, "1="},
		{4, "1-"},
		{5, "10"},
		{6, "11"},
		{7, "12"},
		{8, "2="},
		{9, "2-"},
		{10, "20"},
		{15, "1=0"},
		{20, "1-0"},
		{2022, "1=11-2"},
		{12345, "1-0---0"},
		{314159265, "1121-1110-1=0"},
	}
	for _, elt := range table {
		res := SnafuToInt(elt.snafu)
		if res != elt.num {
			t.Errorf("got %d, wanted %d", res, elt.num)
		}
		snafu := IntToSnafu(elt.num)
		if snafu != elt.snafu {
			t.Errorf("got %s, wanted %s for %d", snafu, elt.snafu, elt.num)
		}
	}
	check("small.txt", "2=-1=0", t, ComputeResult)
	check("input.txt", "2-=2==00-0==2=022=10", t, ComputeResult)
}
