// https://adventofcode.com/2022/day/1

package lib

import (
    "fmt"
)

func FindElf(elts [][]int) (res int) {
	max := 0
	for i := range elts {
		sum := 0
		for j := range elts[i] {
			sum += elts[i][j]
		}
		if (sum > max) {
			max = sum
		}
	}
	return max
}

func FindElf2(elts [][]int) (res int) {
	max := []int{0, 0, 0}
	for i := range elts {
		sum := 0
		for j := range elts[i] {
			sum += elts[i][j]
		}
		for k := range max {
			if (sum > max[k]) {
				for l := 2; l > k; l-- {
					max[l] = max[l - 1]
				}
				max[k] = sum
				fmt.Println("Max", k, max[k])
				break
			}
		}
	}
	ret := 0
	for j := range max {
		ret += max[j]
	}
	return ret
}
