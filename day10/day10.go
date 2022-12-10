// https://adventofcode.com/2022/day/10

package day10

import (
	"fmt"
	"strconv"
	"strings"
)

func incCycles(reg int, cycles int, signal int) (int, int, int) {
	cycles += 1
	if (cycles-20)%40 == 0 {
		strength := cycles * reg
		// fmt.Println("cycle", cycles, reg, strength)
		signal += strength
	}
	return reg, cycles, signal
}

func noop(reg int, cycles int, signal int) (int, int, int) {
	reg, cycles, signal = incCycles(reg, cycles, signal)
	return reg, cycles, signal
}

func addx(x int, reg int, cycles int, signal int) (int, int, int) {
	reg, cycles, signal = incCycles(reg, cycles, signal)
	reg, cycles, signal = incCycles(reg, cycles, signal)
	reg += x
	// fmt.Println("add", x, reg, cycles)
	return reg, cycles, signal
}

func ComputeResult(lines []string) (signal int) {
	fmt.Println("ComputeResult")
	reg := 1
	cycles := 0
	for idx := 0; idx < len(lines); idx++ {
		inst := strings.Split(lines[idx], " ")
		switch inst[0] {
		case "noop":
			reg, cycles, signal = noop(reg, cycles, signal)
		case "addx":
			x, _ := strconv.Atoi(inst[1])
			reg, cycles, signal = addx(x, reg, cycles, signal)
		}
	}
	fmt.Println("result", reg, cycles, signal)
	return signal
}
