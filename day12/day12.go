// https://adventofcode.com/2022/day/12

package day12

import (
	"fmt"
)

var possiblemoves = [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func visited(x int, y int, path [][2]int) bool {
	for idx := range path {
		if x == path[idx][0] && y == path[idx][1] {
			return true
		}
	}
	return false
}

func getHigh(x int, y int, lines []string) int {
	val := lines[y][x]
	if val == 'S' {
		return 0
	}
	if val >= 'a' && val <= 'z' {
		return int(val - 'a' + 1)
	} else if val == 'E' {
		return int('z' - 'a' + 2)
	} else {
		return 1000000
	}
}

func getMoves(moves [][2]int, x int, y int, lines []string) [][2]int {
	for idx := range possiblemoves {
		newX := x + possiblemoves[idx][0]
		newY := y + possiblemoves[idx][1]

		// filter out invalid moves
		if newY < 0 || newY >= len(lines) {
			continue
		}
		if newX < 0 || newX >= len(lines[newY]) {
			continue
		}
		// filter lower point
		if getHigh(newX, newY, lines) > getHigh(x, y, lines)+1 {
			continue
		}
		// filter out if already on the list
		if visited(newX, newY, moves) {
			continue
		}
		move := [2]int{newX, newY}
		moves = append(moves, move)
	}
	return moves
}

func findTarget(lines []string) (x int, y int) {
	for y = range lines {
		for x = range lines[y] {
			if lines[y][x] == 'E' {
				return x, y
			}
		}
	}
	return x, y
}

func ComputeResult(lines []string) int {
	targetX, targetY := findTarget(lines)
	fmt.Println("ComputeResult", targetX, targetY)
	moves := getMoves(nil, 0, 0, lines)
	length := 0
	finished := false
	for !finished {
		futureMoves := [][2]int{}
		length += 1
		for i := range moves {
			move := moves[i]
			futureMoves = getMoves(futureMoves, move[0], move[1], lines)
			finished = move[0] == targetX && move[1] == targetY
			lines[move[1]] = lines[move[1]][:move[0]] + "+" + lines[move[1]][move[0]+1:]
			if finished {
				break
			}
		}
		moves = futureMoves
		if !finished && len(moves) == 0 {
			fmt.Println("not found")
			return -1
		}
		fmt.Print("\033[H\033[2J")
		fmt.Println()
		for i := range lines {
			fmt.Println(lines[i])
		}
	}
	return length
}
