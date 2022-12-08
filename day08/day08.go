// https://adventofcode.com/2022/day/8

package day08

import (
	"fmt"
	"strconv"
)

func checkVisibility(i int, j int, lines []string) bool {
	val, _ := strconv.Atoi(lines[i][j : j+1])
	visible := true
	for x := j + 1; x < len(lines[i]); x++ {
		other, _ := strconv.Atoi(lines[i][x : x+1])
		if val <= other {
			visible = false
			break
		}
	}
	if visible {
		fmt.Println("vis +x", i, j, val)
		return true
	}
	// x negative
	visible = true
	for x := j - 1; x >= 0; x-- {
		other, _ := strconv.Atoi(lines[i][x : x+1])
		if val <= other {
			visible = false
			break
		}
	}
	if visible {
		fmt.Println("vis -x", i, j, val)
		return true
	}
	// y positive
	visible = true
	for y := i + 1; y < len(lines[i]); y++ {
		other, _ := strconv.Atoi(lines[y][j : j+1])
		if val <= other {
			visible = false
			break
		}
	}
	if visible {
		fmt.Println("vis +y", i, j, val)
		return true
	}
	// y negative
	visible = true
	for y := i - 1; y >= 0; y-- {
		other, _ := strconv.Atoi(lines[y][j : j+1])
		if val <= other {
			visible = false
			break
		}
	}
	if visible {
		fmt.Println("vis -y", i, j, val)
		return true
	}
	fmt.Println("no vis", i, j, val)
	return false
}

func ComputeResult(lines []string) (total int) {
	fmt.Println("ComputeResult")

	total = len(lines)*4 - 4
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			if checkVisibility(i, j, lines) {
				total += 1
			}
		}
	}
	return total
}
