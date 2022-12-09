// https://adventofcode.com/2022/day/9

package day09

import (
	"fmt"
	"strconv"
	"strings"
)

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}

}

func moveTail(x int, y int, tx int, ty int) (int, int) {
	ax := abs(tx - x)
	ay := abs(ty - y)
	if ax <= 1 && ay <= 1 {
		return tx, ty
	}
	// only horizontal or vertical
	if ax+ay == 1 {
		tx += (x - tx) / 2
		ty += (y - ty) / 2
	} else { // diagonal or same
		if ax == 2 {
			tx += (x - tx) / 2
			ty += (y - ty)
		}
		if ay == 2 {
			tx += (x - tx)
			ty += (y - ty) / 2
		}
	}
	return tx, ty
}

func recordPosition(x int, y int, pos map[string]bool) {
	key := fmt.Sprintf("%d %d", x, y)
	pos[key] = true

}

func ComputeResult(lines []string) int {
	fmt.Println("ComputeResult")
	var positions map[string]bool
	positions = make(map[string]bool)
	x := 0
	y := 0
	tx := 0
	ty := 0
	recordPosition(tx, ty, positions)
	for idx := 0; idx < len(lines); idx++ {
		move := strings.Split(lines[idx], " ")
		// fmt.Println(move)
		number, _ := strconv.Atoi(move[1])
		for loop := 0; loop < number; loop++ {
			switch move[0] {
			case "U":
				y += 1
			case "D":
				y -= 1
			case "R":
				x += 1
			case "L":
				x -= 1
			}
			tx, ty = moveTail(x, y, tx, ty)
			// fmt.Println(x, y, tx, ty)
			recordPosition(tx, ty, positions)
		}
	}
	return len(positions)
}
