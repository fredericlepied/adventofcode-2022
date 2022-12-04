package lib

import "testing"
import (
	"bufio"
	"fmt"
	re "regexp"
	"strconv"
	"os"
)

func TestComputeScore(t *testing.T){

	got := ComputeScore(shortList)
	want := 2
	
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
	
	longList := ReadFile("input.txt")
	
	got = ComputeScore(longList)
	want = 507
	
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	got = ComputeScore2(shortList)
	want = 4
	
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
	
	got = ComputeScore2(longList)
	want = 897
	
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Map(vs []string, f func(string) (int, error)) []int {
    vsm := make([]int, len(vs))
    for i, v := range vs {
	    vsm[i], _ = f(v)
    }
    return vsm
}

func ReadFile(fpath string) (res [][]int) {
	longList := [][]int{}
	
	readFile, err := os.Open(fpath)
	
	if err != nil {
		fmt.Println(err)
		return longList
	}
	fileScanner := bufio.NewScanner(readFile)

	number_regexp := re.MustCompile("[,-]")
	
	fileScanner.Split(bufio.ScanLines)
	
	for fileScanner.Scan() {
		text := fileScanner.Text()
		list_str := number_regexp.Split(text, 4)
		list := Map(list_str, strconv.Atoi)
		longList = append(longList, list)
	}
	
	readFile.Close()	
	return longList
}

var shortList = [][]int {
	{2, 4, 6, 8},
	{2, 3, 4, 5},
	{5, 7, 7, 9},
	{2, 8, 3, 7},
	{6, 6, 4, 6},
	{2, 6, 4, 8},
}
