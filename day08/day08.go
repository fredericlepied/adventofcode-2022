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
		return true
	}
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

func computeScenic(i int, j int, lines []string) int {
	val, _ := strconv.Atoi(lines[i][j : j+1])
	// right
	scenic_x1 := 0
	for x := j + 1; x < len(lines[i]); x++ {
		scenic_x1 += 1
		other, _ := strconv.Atoi(lines[i][x : x+1])
		if val <= other {
			break
		}
	}
	// left
	scenic_x2 := 0
	for x := j - 1; x >= 0; x-- {
		scenic_x2 += 1
		other, _ := strconv.Atoi(lines[i][x : x+1])
		if val <= other {
			break
		}
	}
	// up
	scenic_y1 := 0
	for y := i + 1; y < len(lines[i]); y++ {
		scenic_y1 += 1
		other, _ := strconv.Atoi(lines[y][j : j+1])
		if val <= other {
			break
		}
	}
	// down
	scenic_y2 := 0
	for y := i - 1; y >= 0; y-- {
		scenic_y2 += 1
		other, _ := strconv.Atoi(lines[y][j : j+1])
		if val <= other {
			break
		}
	}
	return scenic_x1 * scenic_x2 * scenic_y1 * scenic_y2
}

func ComputeResult2(lines []string) (total int) {
	fmt.Println("ComputeResult2")

	total = 0
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			scenic := computeScenic(i, j, lines)
			if scenic > total {
				total = scenic
			}
		}
	}
	return total
}
