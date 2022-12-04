// https://adventofcode.com/2022/day/2

package lib

import (
	"fmt"
)

func ComputeLine(line []int) (res int) {
	if (line[0] <= line[2] && line[1] >= line[3]) {
		return 1
	} else if (line[2] <= line[0] && line[3] >= line[1]) {
		return 1
	}
	return 0
}

func ComputeScore(elts [][]int) (res int) {
	fmt.Println("ComputeScore")
	score := 0
	
	for i := 0; i< len(elts); i++ {
		score += ComputeLine(elts[i])
	}

	return score
}
