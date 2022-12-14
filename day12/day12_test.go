package day12

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

func TestComputeResult(t *testing.T) {
	// part 1
	//check(veryShortList, 31, t)
	check(shortList, 31, t)

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

var veryShortList = []string{
	"aEbqponm",
}

var shortList = []string{
	"Sabqponm",
	"abcryxxl",
	"accszExk",
	"acctuvwj",
	"abdefghi",
}
