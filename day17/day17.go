// https://adventofcode.com/2022/day/17

package day17

import (
	"fmt"
	. "github.com/fredericlepied/adventofcode-2022/utils"
)

var debug = true

func b(str string) []byte {
	return []byte(str)

}

func displayChamber(debug bool, chamber [][]byte) {
	if !debug {
		return
	}
	for y := len(chamber) - 1; y >= 0; y-- {
		fmt.Println(string(chamber[y]))
	}
}

func processRock(rock []string, chamber [][]byte, directions string, round int) ([][]byte, int) {
	// check if there are 3 emtpy lines at the top bellow the rock
	newLines := 3 + len(rock)
	extra := newLines
	Debug(debug, "need", newLines, "empty lines")
	for y := len(chamber) - 1; y >= len(chamber)-extra; y-- {
		if Index('#', chamber[y]) != -1 || Index('-', chamber[y]) != -1 {
			Debug(debug, "adding", newLines, "new lines")
			for n := 0; n < newLines; n++ {
				chamber = append(chamber, b("|.......|"))
			}
			break
		}
		newLines--
	}
	displayChamber(debug, chamber)

	// movements
	x := 3
	y := len(chamber) - 1
	width := 0
	for line := range rock {
		width = Max(width, len(rock[line]))
	}
	level := 0
	for {
		current := y - level
		// horizontal move
		if directions[round%len(directions)] == '>' {
			Debug(debug, "right")
			// check that the rock can move right
			canMove := true
			for idx := 0; idx < len(rock); idx++ {
				if chamber[current][x+width] != '.' {
					canMove = false
					break
				}
			}
			if canMove {
				x++
				Debug(debug, "moved right", x)
			}
		} else {
			Debug(debug, "left")
			// check that the rock can move left
			canMove := true
			for idx := 0; idx < len(rock); idx++ {
				if chamber[current][x-1] != '.' {
					canMove = false
					break
				}
			}
			if canMove {
				x--
				Debug(debug, "moved left", x)
			}
		}
		Debug(debug, "pos", x, current, "round", round)

		// check if the rock can go down
		canMove := true
		for dx := 0; dx < len(rock[0]); dx++ {
			xx := x + dx
			yy := current - len(rock)
			fmt.Printf("checking %d,%d",
				xx,
				yy)
			fmt.Printf(" '%c' %c'\n",
				rock[len(rock)-1][dx], chamber[yy][xx])
			if rock[len(rock)-1][dx] == '#' && chamber[yy][xx] != '.' {
				canMove = false
				break
			}
		}
		round++
		if !canMove {
			for dy := 0; dy < len(rock); dy++ {
				for dx := 0; dx < len(rock[dy]); dx++ {
					if rock[dy][dx] == '#' {
						chamber[current-dy][x+dx] = '#'
					}
				}
			}
			break
		} else {
			Debug(debug, "down")
		}
		level++
	}
	return chamber, round
}

func ComputeResult(directions string) (result int) {
	Debug(debug, "ComputeResult")

	rock1 := []string{"####"}
	rock2 := []string{".#.", "###", ".#."}
	rock3 := []string{"..#", "..#", "###"}
	rock4 := []string{"#", "#", "#", "#"}
	rock5 := []string{"##", "##"}
	rocks := [][]string{rock1, rock2, rock3, rock4, rock5}

	chamber := [][]byte{b("+-------+"), b("|.......|"), b("|.......|")}

	round := 0
	for loop := 1; loop <= 2022; loop++ {
		if loop == 10 {
			debug = false
		}
		idx := (loop - 1) % 5
		rock := rocks[idx]
		Debug(debug, "rock", idx+1, "loop", loop)
		Debug(debug)
		for y := 0; y < len(rock); y++ {
			Debug(debug, rock[y])
		}
		Debug(debug)
		chamber, round = processRock(rock, chamber, directions, round)
		displayChamber(debug, chamber)
		Debug(debug)
	}
	Debug(debug)
	displayChamber(true, chamber)
	for y := len(chamber) - 1; y >= 0; y-- {
		if Index('#', chamber[y]) != -1 {
			return y
		}
	}
	return result
}
