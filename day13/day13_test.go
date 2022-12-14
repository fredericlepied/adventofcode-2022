package day13

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func check(input []string, want int, t *testing.T) {
	got := ComputeResult(input)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func checkFirst(input string, car string, rest string, t *testing.T) {
	got0, got1 := GetFirst(input)

	if got0 != car || got1 != rest {
		t.Errorf("%s: got %s + %s, wanted %s + %s", input, got0, got1, car, rest)
	}
}

func TestComputeResult(t *testing.T) {
	// part 1
	checkFirst("[]", "[]", "", t)
	checkFirst("[1]", "1", "[]", t)
	checkFirst("[1,1,3,1,1]", "1", "[1,3,1,1]", t)
	checkFirst("[[[]]]", "[[]]", "[]", t)
	checkFirst("[[2,3,4]]", "[2,3,4]", "[]", t)
	checkFirst("[[1],4]", "[1]", "[4]", t)
	check(shortList, 13, t)
	longList := readFile("input.txt")
	check(longList, 2, t)

	// part 2
}

func readFile(fpath string) (lines []string) {
	readFile, err := os.Open(fpath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}
	readFile.Close()

	return lines
}

var shortList = []string{
	"[1,1,3,1,1]",
	"[1,1,5,1,1]",
	"",
	"[[1],[2,3,4]]",
	"[[1],4]",
	"",
	"[9]",
	"[[8,7,6]]",
	"",
	"[[4,4],4,4]",
	"[[4,4],4,4,4]",
	"",
	"[7,7,7,7]",
	"[7,7,7]",
	"",
	"[]",
	"[3]",
	"",
	"[[[]]]",
	"[[]]",
	"",
	"[1,[2,[3,[4,[5,6,7]]]],8,9]",
	"[1,[2,[3,[4,[5,6,0]]]],8,9]",
}
