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

type Model struct {
	X int
	Y int
	Z int
}

func (self *Model) Add(m Model) Model {
	return Model{self.X + m.X, self.Y + m.Y, self.Z + m.Z}
}

func ComputeResult(lines []string) (result int) {
	fmt.Println("ComputeResult")

	cubes := map[Model]bool{}

	for idx := range lines {
		line := strings.Split(lines[idx], ",")
		cubes[Model{ToInt(line[0]), ToInt(line[1]), ToInt(line[2])}] = true
	}
	directions := []Model{Model{1, 0, 0}, Model{-1, 0, 0}, Model{0, 1, 0}, Model{0, -1, 0}, Model{0, 0, 1}, Model{0, 0, -1}}
	result = 6 * len(cubes)
	for cube := range cubes {
		for i := range directions {
			if _, ok := cubes[cube.Add(directions[i])]; ok {
				result--
			}
		}
	}
	return result
}
