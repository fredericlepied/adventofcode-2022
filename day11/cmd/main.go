package main

import (
	"fmt"
	d "github.com/fredericlepied/adventofcode-2022/day11"
	u "github.com/fredericlepied/adventofcode-2022/utils"
)

func main() {
	longList := u.ReadFile("input.txt")
	result := d.ComputeResult(longList, 1, 10000)
	fmt.Println("Result=", result)
}
