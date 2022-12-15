package day15

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func check(fname string, y int, want int, t *testing.T) {
	input := readFile(fname)
	got := ComputeResult(input, y)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestComputeResult(t *testing.T) {
	// part 1
	check("small.txt", 10, 26, t)
	check("input.txt", 2000000, 4737443, t)
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
