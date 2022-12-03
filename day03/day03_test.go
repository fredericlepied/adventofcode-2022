package lib

import "testing"
import (
	"bufio"
	"fmt"
	"os"
)

func TestComputeScore(t *testing.T){

	got := ComputeLine("AaAa")
	want := 28
	
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
	
	got = ComputeScore(shortList)
	want = 157
	
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
	
	longList := ReadFile("input.txt")
	
	got = ComputeScore(longList)
	want = 8349
	
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	got = ComputeScore2(shortList)
	want = 70
	
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
	
	got = ComputeScore2(longList)
	want = 2681
	
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func ReadFile(fpath string) (res []string) {
	longList := []string{}
	
	readFile, err := os.Open(fpath)
	
	if err != nil {
		fmt.Println(err)
		return longList
	}
	fileScanner := bufio.NewScanner(readFile)
	
	fileScanner.Split(bufio.ScanLines)
	
	for fileScanner.Scan() {
		text := fileScanner.Text()
		longList = append(longList, text)
	}
	
	readFile.Close()	
	return longList
}

var shortList = []string {
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}
