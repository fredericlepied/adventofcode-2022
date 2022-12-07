// https://adventofcode.com/2022/day/2

package lib

import (
	"fmt"
)

func ComputeRound(round []string) (res int) {
	lost := 0
	draw := 3
	won := 6
	rock := 1
	paper := 2
	scissors := 3
	// rock -> 1
	if round[0] == "A" {
		// rock
		if round[1] == "X" {
			return rock + draw
			// paper
		} else if round[1] == "Y" {
			return paper + won
			// scissors
		} else if round[1] == "Z" {
			return scissors + lost
		}
		// paper -> 2
	} else if round[0] == "B" {
		// rock
		if round[1] == "X" {
			return rock + lost
			// paper
		} else if round[1] == "Y" {
			return paper + draw
			// scissors
		} else if round[1] == "Z" {
			return scissors + won
		}
		// scissors -> 3
	} else if round[0] == "C" {
		// rock
		if round[1] == "X" {
			return rock + won
			// paper
		} else if round[1] == "Y" {
			return paper + lost
			// scissors
		} else if round[1] == "Z" {
			return scissors + draw
		}
	} else {
		fmt.Println("Invalid first character", round[0])
	}
	return 0
}

func ComputeScore(elts [][]string) (res int) {
	fmt.Println("ComputeScore")
	score := 0
	for idx := range elts {
		score += ComputeRound(elts[idx])
	}
	return score
}

func ComputeRound2(round []string) (res int) {
	// rock -> 1
	if round[0] == "A" {
		// lose
		if round[1] == "X" {
			round[1] = "Z"
			// draw
		} else if round[1] == "Y" {
			round[1] = "X"
			// win
		} else if round[1] == "Z" {
			round[1] = "Y"
		}
		// paper -> 2
	} else if round[0] == "B" {
		// lose
		if round[1] == "X" {
			round[1] = "X"
			// draw
		} else if round[1] == "Y" {
			round[1] = "Y"
			// win
		} else if round[1] == "Z" {
			round[1] = "Z"
		}
		// scissors -> 3
	} else if round[0] == "C" {
		// lose
		if round[1] == "X" {
			round[1] = "Y"
			// draw
		} else if round[1] == "Y" {
			round[1] = "Z"
			// win
		} else if round[1] == "Z" {
			round[1] = "X"
		}
	} else {
		fmt.Println("Invalid first character", round[0])
	}
	return ComputeRound(round)
}

func ComputeScore2(elts [][]string) (res int) {
	fmt.Println("ComputeScore2")
	score := 0
	for idx := range elts {
		score += ComputeRound2(elts[idx])
	}
	return score
}
