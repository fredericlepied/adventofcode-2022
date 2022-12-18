// https://adventofcode.com/2022/day/18

package day18

import (
	"fmt"
	. "github.com/fredericlepied/adventofcode-2022/utils"
	"strings"
)

var debug = true

func b(str string) []byte {
	return []byte(str)

}

type Cube struct {
	X int
	Y int
	Z int
}

func (self *Cube) Add(m Cube) Cube {
	return Cube{self.X + m.X, self.Y + m.Y, self.Z + m.Z}
}

func ComputeResult(lines []string) (result int) {
	fmt.Println("ComputeResult")

	cubes := map[Cube]bool{}

	for idx := range lines {
		line := strings.Split(lines[idx], ",")
		cubes[Cube{ToInt(line[0]), ToInt(line[1]), ToInt(line[2])}] = true
	}
	directions := []Cube{Cube{1, 0, 0}, Cube{-1, 0, 0}, Cube{0, 1, 0}, Cube{0, -1, 0}, Cube{0, 0, 1}, Cube{0, 0, -1}}
	result = 6 * len(cubes)
	for cube := range cubes {
		for i := range directions {
			if checkCube(cube.Add(directions[i]), cubes) {
				result--
			}
		}
	}
	return result
}

func checkCube(cube Cube, cubes map[Cube]bool) bool {
	if _, ok := cubes[cube]; ok {
		return true
	}
	return false
}

func ComputeResult2(lines []string) (result int) {
	fmt.Println("ComputeResult2")

	cubes := map[Cube]bool{}

	max := Cube{-100000, -100000, -100000}
	min := Cube{100000, 100000, 100000}

	for idx := range lines {
		line := strings.Split(lines[idx], ",")
		cube := Cube{ToInt(line[0]), ToInt(line[1]), ToInt(line[2])}
		cubes[cube] = true
		max.X = Max(max.X, cube.X)
		max.Y = Max(max.Y, cube.Y)
		max.Z = Max(max.Z, cube.Z)
		min.X = Min(min.X, cube.X)
		min.Y = Min(min.Y, cube.Y)
		min.Z = Min(min.Z, cube.Z)
	}
	Debug(debug, min, max)
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			for z := min.Z; z <= max.Z; z++ {
				if checkCube(Cube{x, y, z}, cubes) {
					result++
					break
				}
			}
		}
	}
	for x := min.X; x <= max.X; x++ {
		for y := min.Y; y <= max.Y; y++ {
			for z := max.Z; z >= min.Z; z-- {
				if checkCube(Cube{x, y, z}, cubes) {
					result++
					break
				}
			}
		}
	}
	for x := min.X; x <= max.X; x++ {
		for z := min.Z; z <= max.Z; z++ {
			for y := min.Y; y <= max.Y; y++ {
				if checkCube(Cube{x, y, z}, cubes) {
					result++
					break
				}
			}
		}
	}
	for x := min.X; x <= max.X; x++ {
		for z := min.Z; z <= max.Z; z++ {
			for y := max.Y; y >= min.Y; y-- {
				if checkCube(Cube{x, y, z}, cubes) {
					result++
					break
				}
			}
		}
	}
	for z := min.Z; z <= max.Z; z++ {
		for y := min.Y; y <= max.Y; y++ {
			for x := min.X; x <= max.X; x++ {
				if checkCube(Cube{x, y, z}, cubes) {
					result++
					break
				}
			}
		}
	}
	for z := min.Z; z <= max.Z; z++ {
		for y := min.Y; y <= max.Y; y++ {
			for x := max.X; x >= min.X; x-- {
				if checkCube(Cube{x, y, z}, cubes) {
					result++
					break
				}
			}
		}
	}
	return result
}
