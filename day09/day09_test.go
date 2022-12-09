package day09

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

func check2(input []string, want int, t *testing.T) {
	got := ComputeResult2(input)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestComputeResult(t *testing.T) {
	// part 1
	check(shortList, 13, t)

	longList := readFile("input.txt")
	check(longList, 6236, t)

	// part 2
	check2(shortList, 1, t)
	check2(shortList2, 36, t)
	check2(longList, 2449, t)
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
	"R 4",
	"U 4",
	"L 3",
	"D 1",
	"R 4",
	"D 1",
	"L 5",
	"R 2",
}

var shortList2 = []string{
	"R 5",
	"U 8",
	"L 8",
	"D 3",
	"R 17",
	"D 10",
	"L 25",
	"U 20",
}
