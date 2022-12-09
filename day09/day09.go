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
	if ax == 2 && ay == 2 {
		tx += (x - tx) / 2
		ty += (y - ty) / 2
	} else if ax == 2 {
		tx += (x - tx) / 2
		ty += (y - ty)
	} else if ay == 2 {
		tx += (x - tx)
		ty += (y - ty) / 2
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

func ComputeResult2(lines []string) int {
	fmt.Println("ComputeResult2")
	var positions map[string]bool
	positions = make(map[string]bool)
	x := 0
	y := 0
	var tx [9]int
	var ty [9]int
	last := len(tx) - 1
	recordPosition(tx[last], ty[last], positions)
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
			tx[0], ty[0] = moveTail(x, y, tx[0], ty[0])
			for n := 1; n < len(tx); n++ {
				tx[n], ty[n] = moveTail(tx[n-1], ty[n-1], tx[n], ty[n])
			}
			recordPosition(tx[last], ty[last], positions)
		}
	}
	return len(positions)
}
