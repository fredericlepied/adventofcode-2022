package day06

import "testing"
import (
	"bufio"
	"fmt"
	"os"
)

func check(input string, want int, t *testing.T) {
	got := ComputeResult(input)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestComputeResult(t *testing.T) {
	check("bvwbjplbgvbhsrlpgdmjqwftvncz", 5, t)
	check("nppdvjthqldpwncqszvftbrmjlhg", 6, t)
	check("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10, t)
	check("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11, t)

	longText := readFile("input.txt")
	check(longText, 1909, t)
}

func readFile(fpath string) (text string) {
	readFile, err := os.Open(fpath)

	if err != nil {
		fmt.Println(err)
		return ""
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	text = fileScanner.Text()

	readFile.Close()

	return text
}
