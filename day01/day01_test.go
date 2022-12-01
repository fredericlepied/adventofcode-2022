package lib

import "testing"
import (
	"bufio"
	"fmt"
	"strconv"
	"os"
)

func TestFindElf(t *testing.T){

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
}

func TestFindElfLong(t *testing.T){
	longList := [][]int{}

	readFile, err := os.Open("input.txt")
	
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	
	fileScanner.Split(bufio.ScanLines)
	
	calories := []int{}
	for fileScanner.Scan() {
		text := fileScanner.Text()
		if (text == "") {
			longList = append(longList, calories)
			calories = []int{}
		} else {
			val, _ := strconv.Atoi(text)
			calories = append(calories, val)
		}
	}
	
	readFile.Close()

	got := FindElf(longList)
	want := 68775
	
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}

	got = FindElf2(longList)
	want = 202585
	
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

var shortList = [][]int {
	{1000, 2000, 3000},
	{4000},
	{5000, 6000},
	{7000, 8000, 9000},
	{10000} }
