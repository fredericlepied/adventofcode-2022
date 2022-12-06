package day06

import "testing"
import (
	"bufio"
	"fmt"
	"os"
)

func check(input string, length int, want int, t *testing.T) {
	got := ComputeResult(input, length)

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestComputeResult(t *testing.T) {
	// part 1
	check("bvwbjplbgvbhsrlpgdmjqwftvncz", 4, 5, t)
	check("nppdvjthqldpwncqszvftbrmjlhg", 4, 6, t)
	check("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4, 10, t)
	check("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4, 11, t)

	longText := readFile("input.txt")
	check(longText, 4, 1909, t)
	
	// part 2
	check("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14, 19, t)
	check("bvwbjplbgvbhsrlpgdmjqwftvncz", 14, 23, t)
	check("nppdvjthqldpwncqszvftbrmjlhg", 14, 23, t)
	check("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14, 29, t)
	check("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14, 26, t)
	check(longText, 14, 3380, t)
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
