// https://adventofcode.com/2022/day/15

package day15

import (
	"fmt"
	. "github.com/fredericlepied/adventofcode-2022/utils"
	re "regexp"
)

type Coord struct {
	X int
	Y int
}

func (self Coord) distance(coord Coord) int {
	return Abs(self.X-coord.X) + Abs(self.Y-coord.Y)
}

func displayImg(img [][]byte) {
	for idx := range img {
		fmt.Println(string(img[idx]))
	}
}

func displayLine(line map[int]byte) {
	minX, maxX := 10000000, -10000000
	for x := range line {
		minX = Min(minX, x)
		maxX = Max(maxX, x)
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

func set(x int, y int, val byte, img map[int]byte, targetY int) {
	if y == targetY {
		if v, ok := img[x]; !ok || v == '#' {
			img[x] = val
		}
	}
}

func initBeacons(lines []string, beacons map[Coord]Coord) {
	line_regexp := re.MustCompile("Sensor at x=([-0-9]+), y=([-0-9]+): closest beacon is at x=([-0-9]+), y=([-0-9]+)")

	for idx := 0; idx < len(lines); idx++ {
		line := lines[idx]
		res := line_regexp.FindStringSubmatch(line)
		beacons[Coord{ToInt(res[1]), ToInt(res[2])}] = Coord{ToInt(res[3]), ToInt(res[4])}
	}
}

func ComputeResult(lines []string, targetY int) (result int) {
	fmt.Println("ComputeResult")

	beacons := map[Coord]Coord{}
	initBeacons(lines, beacons)

	img := map[int]byte{}
	for sensor, beacon := range beacons {
		set(sensor.X, sensor.Y, 'S', img, targetY)
		set(beacon.X, beacon.Y, 'B', img, targetY)
		dist := sensor.distance(beacon)
		if targetY >= sensor.Y-dist && targetY <= sensor.Y+dist {
			for x := sensor.X - dist; x < sensor.X+dist; x++ {
				dist2 := Abs(x-sensor.X) + Abs(targetY-sensor.Y)
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

func initBeacons2(lines []string, beacons map[Coord]Coord, distances map[Coord]int) {
	line_regexp := re.MustCompile("Sensor at x=([-0-9]+), y=([-0-9]+): closest beacon is at x=([-0-9]+), y=([-0-9]+)")

	for idx := 0; idx < len(lines); idx++ {
		line := lines[idx]
		res := line_regexp.FindStringSubmatch(line)
		sensor := Coord{ToInt(res[1]), ToInt(res[2])}
		beacon := Coord{ToInt(res[3]), ToInt(res[4])}
		dist := sensor.distance(beacon)
		beacons[sensor] = beacon
		distances[sensor] = dist
	}
}

func canContainUnseenPoint(
	sensor Coord,
	distances map[Coord]int,
	minVal Coord,
	maxVal Coord) bool {
	corners := []Coord{
		{X: minVal.X, Y: minVal.Y},
		{X: minVal.X, Y: maxVal.Y},
		{X: maxVal.X, Y: minVal.Y},
		{X: maxVal.X, Y: maxVal.Y},
	}
	for _, corner := range corners {
		distance := sensor.distance(corner)
		if distance > distances[sensor] {
			return true
		}
	}
	return false
}

func findUnseenPoint(
	beacons map[Coord]Coord,
	distances map[Coord]int,
	minVal Coord,
	maxVal Coord) Coord {
	if minVal == maxVal {
		return minVal
	}
	mid := Coord{(minVal.X + maxVal.X) / 2, (minVal.Y + maxVal.Y) / 2}

	quadrants := make([][]Coord, 4)
	quadrants[0] = []Coord{minVal, mid}
	quadrants[1] = []Coord{Coord{X: mid.X + 1, Y: minVal.Y}, Coord{X: maxVal.X, Y: mid.Y}}
	quadrants[2] = []Coord{Coord{X: minVal.X, Y: mid.Y + 1}, Coord{X: mid.X, Y: maxVal.Y}}
	quadrants[3] = []Coord{Coord{X: mid.X + 1, Y: mid.Y + 1}, maxVal}

	for _, quadrant := range quadrants {
		if quadrant[0].X > quadrant[1].X || quadrant[0].Y > quadrant[1].Y {
			continue
		}

		allPairsCanContain := true
		for sensor := range beacons {
			if !canContainUnseenPoint(sensor, distances, quadrant[0], quadrant[1]) {
				allPairsCanContain = false
				break
			}
		}
		if allPairsCanContain {
			k := findUnseenPoint(beacons, distances, quadrant[0], quadrant[1])
			if k.X != -1 || k.Y != -1 {
				return k
			}
		}
	}
	return Coord{X: -1, Y: -1}
}

func ComputeResult2(lines []string, maxN int) (result int) {
	fmt.Println("ComputeResult2")

	beacons := map[Coord]Coord{}
	distances := map[Coord]int{}
	initBeacons2(lines, beacons, distances)

	point := findUnseenPoint(beacons, distances, Coord{0, 0}, Coord{maxN, maxN})
	return point.X*4000000 + point.Y
}
