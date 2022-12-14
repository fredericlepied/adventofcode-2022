// https://adventofcode.com/2022/day/14

package day14

import (
	"fmt"
	"image"
	"strconv"
	"strings"
)

func displayImg(img [][]byte) {
	for idx := range img {
		fmt.Println(string(img[idx]))
	}
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a int, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}

const base = 300
const width = 400

// 480 -> 540
func ComputeResult(input []byte, floor bool) (count int) {
	fmt.Println("ComputeResult")
	img := [][]byte{}
	maxY := 0

	for idx := 0; idx < 180; idx++ {
		line := []byte{}
		for x := 0; x < width; x++ {
			line = append(line, '.')
		}
		img = append(img, line)
	}

	for _, line := range strings.Split(strings.Trim(string(input), "\n"), "\n") {
		queue := []image.Point{}
		coords := strings.Split(line, " -> ")
		for _, coord := range coords {
			res := strings.Split(coord, ",")
			x, _ := strconv.Atoi(res[0])
			y, _ := strconv.Atoi(res[1])
			img[y][x-base] = '#'
			queue = append(queue, image.Point{x, y})
		}
		prev := queue[0]
		maxY = prev.Y
		for idx := 1; idx < len(queue); idx++ {
			cur := queue[idx]
			maxY = max(maxY, cur.Y)
			//fmt.Println("line", cur, "->", prev)
			if prev.X == cur.X {
				for y := min(prev.Y, cur.Y); y < max(prev.Y, cur.Y); y++ {
					//fmt.Println("pt", prev.X, y)
					img[y][prev.X-base] = '#'
				}
			} else {
				for x := min(prev.X, cur.X); x < max(prev.X, cur.X); x++ {
					//fmt.Println("pt", x, prev.Y)
					img[prev.Y][x-base] = '#'
				}
			}
			prev = cur
		}
	}

	if floor {
		for x := 0; x < width; x++ {
			img[maxY+2][x] = '#'
		}
		maxY += 2
	}

	for {
		x := 500 - base
		laid := false
		for y := 0; y < maxY+1; y++ {
			if img[y+1][x] != '.' {
				if img[y+1][x-1] == '.' {
					x--
				} else if img[y+1][x+1] == '.' {
					x++
				} else {
					img[y][x] = 'o'
					laid = true
					//fmt.Println("laid1", x, y)
					break
				}
			}
		}
		if !laid {
			break
		}
		if img[0][x] != '.' {
			count++
			break
		}
		count++
	}

	displayImg(img)
	return count
}
