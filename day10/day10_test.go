package day10

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
	check(shortList, 0, t)

	check(shortList2, 13140, t)

	longList := readFile("input.txt")
	check(longList, 14540, t)

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
	"noop",
	"addx 3",
	"addx -5",
}

var shortList2 = []string{
	"addx 15",
	"addx -11",
	"addx 6",
	"addx -3",
	"addx 5",
	"addx -1",
	"addx -8",
	"addx 13",
	"addx 4",
	"noop",
	"addx -1",
	"addx 5",
	"addx -1",
	"addx 5",
	"addx -1",
	"addx 5",
	"addx -1",
	"addx 5",
	"addx -1",
	"addx -35",
	"addx 1",
	"addx 24",
	"addx -19",
	"addx 1",
	"addx 16",
	"addx -11",
	"noop",
	"noop",
	"addx 21",
	"addx -15",
	"noop",
	"noop",
	"addx -3",
	"addx 9",
	"addx 1",
	"addx -3",
	"addx 8",
	"addx 1",
	"addx 5",
	"noop",
	"noop",
	"noop",
	"noop",
	"noop",
	"addx -36",
	"noop",
	"addx 1",
	"addx 7",
	"noop",
	"noop",
	"noop",
	"addx 2",
	"addx 6",
	"noop",
	"noop",
	"noop",
	"noop",
	"noop",
	"addx 1",
	"noop",
	"noop",
	"addx 7",
	"addx 1",
	"noop",
	"addx -13",
	"addx 13",
	"addx 7",
	"noop",
	"addx 1",
	"addx -33",
	"noop",
	"noop",
	"noop",
	"addx 2",
	"noop",
	"noop",
	"noop",
	"addx 8",
	"noop",
	"addx -1",
	"addx 2",
	"addx 1",
	"noop",
	"addx 17",
	"addx -9",
	"addx 1",
	"addx 1",
	"addx -3",
	"addx 11",
	"noop",
	"noop",
	"addx 1",
	"noop",
	"addx 1",
	"noop",
	"noop",
	"addx -13",
	"addx -19",
	"addx 1",
	"addx 3",
	"addx 26",
	"addx -30",
	"addx 12",
	"addx -1",
	"addx 3",
	"addx 1",
	"noop",
	"noop",
	"noop",
	"addx -9",
	"addx 18",
	"addx 1",
	"addx 2",
	"noop",
	"noop",
	"addx 9",
	"noop",
	"noop",
	"noop",
	"addx -1",
	"addx 2",
	"addx -37",
	"addx 1",
	"addx 3",
	"noop",
	"addx 15",
	"addx -21",
	"addx 22",
	"addx -6",
	"addx 1",
	"noop",
	"addx 2",
	"addx 1",
	"noop",
	"addx -10",
	"noop",
	"noop",
	"addx 20",
	"addx 1",
	"addx 2",
	"addx 2",
	"addx -6",
	"addx -11",
	"noop",
	"noop",
	"noop",
}
