package lib

import "testing"
import (
	"bufio"
	"fmt"
	"strings"
	"os"
)

func TestFindElf(t *testing.T){

	got := ComputeScore(shortList)
	want := 15
	
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	longList := ReadFile("input.txt")
	
	got = ComputeScore(longList)
	want = 12679
	
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
	
	got = ComputeScore2(shortList)
	want = 12
	
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
	
	got = ComputeScore2(longList)
	want = 14470
	
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func ReadFile(fpath string) (res [][]string) {
	longList := [][]string{}

	readFile, err := os.Open(fpath)
	
	if err != nil {
		fmt.Println(err)
		return longList
	}
	fileScanner := bufio.NewScanner(readFile)
	
	fileScanner.Split(bufio.ScanLines)
	
	for fileScanner.Scan() {
		text := fileScanner.Text()
		res := strings.Split(text, " ")
		longList = append(longList, res)
	}
	
	readFile.Close()	
	return longList
}

var shortList = [][]string {
	{"A", "Y"},
	{"B", "X"},
	{"C", "Z"},
}
