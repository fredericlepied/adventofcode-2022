// https://adventofcode.com/2022/day/15

package day15

import (
	"fmt"
	re "regexp"
	//"sort"
	"image"
	"strconv"
)

func displayImg(img [][]byte) {
	for idx := range img {
		fmt.Println(string(img[idx]))
	}
}

func displayLine(line map[int]byte) {
	minX, maxX := 10000000, -10000000
	for x := range line {
		minX = min(minX, x)
		maxX = max(maxX, x)
	}
	for x := minX; x <= maxX; x++ {
		if _, ok := line[x]; ok {
			fmt.Printf("%c", line[x])
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func s2i(str string) int {
	n, _ := strconv.Atoi(str)
	return n
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

func abs(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}

}

func set(x int, y int, val byte, img map[int]byte, targetY int) {
	if y == targetY {
		if v, ok := img[x]; !ok || v == '#' {
			img[x] = val
		}
	}
}

func ComputeResult(lines []string, targetY int) (result int) {
	fmt.Println("ComputeResult")

	line_regexp := re.MustCompile("Sensor at x=([-0-9]+), y=([-0-9]+): closest beacon is at x=([-0-9]+), y=([-0-9]+)")

	beacons := map[image.Point]image.Point{}
	for idx := 0; idx < len(lines); idx++ {
		line := lines[idx]
		res := line_regexp.FindStringSubmatch(line)
		beacons[image.Point{s2i(res[1]), s2i(res[2])}] = image.Point{s2i(res[3]), s2i(res[4])}
	}
	img := map[int]byte{}
	for sensor, beacon := range beacons {
		set(sensor.X, sensor.Y, 'S', img, targetY)
		set(beacon.X, beacon.Y, 'B', img, targetY)
		dist := abs(sensor.X-beacon.X) + abs(sensor.Y-beacon.Y)
		if targetY >= sensor.Y-dist && targetY <= sensor.Y+dist {
			for x := sensor.X - dist; x < sensor.X+dist; x++ {
				dist2 := abs(x-sensor.X) + abs(targetY-sensor.Y)
				if dist2 <= dist {
					set(x, targetY, '#', img, targetY)
				}
			}
		}
	}
	for idx := range img {
		if img[idx] == '#' {
			result++
		}
	}
	return result
}
