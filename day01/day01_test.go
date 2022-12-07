package lib

import "testing"
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TestFindElf(t *testing.T) {

	got := FindElf(shortList)
	want := 24000

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	got = FindElf2(shortList)
	want = 45000

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	longList := ReadCalories("input.txt")

	got = FindElf(longList)
	want = 68775

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	got = FindElf2(longList)
	want = 202585

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func ReadCalories(fpath string) (res [][]int) {
	longList := [][]int{}

	readFile, err := os.Open(fpath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	calories := []int{}
	for fileScanner.Scan() {
		text := fileScanner.Text()
		if text == "" {
			longList = append(longList, calories)
			calories = []int{}
		} else {
			val, _ := strconv.Atoi(text)
			calories = append(calories, val)
		}
	}

	readFile.Close()
	return longList
}

var shortList = [][]int{
	{1000, 2000, 3000},
	{4000},
	{5000, 6000},
	{7000, 8000, 9000},
	{10000},
}
