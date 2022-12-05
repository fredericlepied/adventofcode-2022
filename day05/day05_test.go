package day05

import "testing"
import (
	"bufio"
	"fmt"
	"os"
	re "regexp"
	"strconv"
	"strings"
)

func TestComputeResult(t *testing.T) {

	shortStacks := make([]Stack, 3)
	shortStacks[0].Push("Z")
	shortStacks[0].Push("N")
	shortStacks[1].Push("M")
	shortStacks[1].Push("C")
	shortStacks[1].Push("D")
	shortStacks[2].Push("P")
	got := ComputeResult(shortList, shortStacks)
	want := "CMZ"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}

	longList, longStacks := ReadFile("input.txt")
	got = ComputeResult(longList, longStacks)
	want = "TBVFVDZPN"

	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func Map(vs []string, f func(string) (int, error)) []int {
	vsm := make([]int, len(vs))
	for i, v := range vs {
		vsm[i], _ = f(v)
	}
	return vsm
}

func ReadFile(fpath string) (orders [][]int, stacks []Stack) {
	readFile, err := os.Open(fpath)

	if err != nil {
		fmt.Println(err)
		return orders, stacks
	}
	fileScanner := bufio.NewScanner(readFile)

	order_regexp := re.MustCompile("move (\\d+) from (\\d+) to (\\d+)")

	fileScanner.Split(bufio.ScanLines)

	state := "init"
	for fileScanner.Scan() {
		text := fileScanner.Text()
		if state == "init" {
			stacks = make([]Stack, (len(text)+1)/4)
			state = "stacks"
		}
		if state == "stacks" {
			if (strings.Contains(text, "[")) {
				for idx := 0; idx <= len(text)/4; idx++ {
					if (text[idx*4 + 1:idx*4 + 2] != " ") {
						stacks[idx].Push(text[idx*4 + 1:idx*4 + 2])
					}
				}
			} else {
				for idx := 0; idx < len(stacks); idx++ {
					stacks[idx].Reverse()
				}
				state = "order"
			}
		}
		if state == "order" {
			res := order_regexp.FindStringSubmatch(text)
			if (len(res) == 4) {
				order := []int {}
				for idx := 1; idx<4; idx++ {
					num, _ := strconv.Atoi(res[idx])
					order = append(order, num)
				}
				orders = append(orders, order)
			}
		}
	}

	readFile.Close()
	return orders, stacks
}

var shortList = [][]int{
	{1, 2, 1},
	{3, 1, 3},
	{2, 2, 1},
	{1, 1, 2},
}
