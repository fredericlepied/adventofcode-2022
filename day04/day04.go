// https://adventofcode.com/2022/day/4

package day04

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

func ComputeScoreHelper(elts [][]int, f func([]int) (int)) (res int) {
	score := 0
	
	for i := 0; i< len(elts); i++ {
		score += f(elts[i])
	}

	return score
}

func ComputeScore(elts [][]int) (res int) {
	fmt.Println("ComputeScore2")
	return ComputeScoreHelper(elts, ComputeLine)
}


func ComputeLine2(line []int) (res int) {
	if (line[0] >= line[2] && line[0] <= line[3]) {
		return 1
	} else if (line[1] >= line[2] && line[1] <= line[3]) {
		return 1
	} else 	if (line[0] <= line[2] && line[1] >= line[3]) {
		return 1
	}
	return 0
}

func ComputeScore2(elts [][]int) (res int) {
	fmt.Println("ComputeScore2")
	return ComputeScoreHelper(elts, ComputeLine2)
}
