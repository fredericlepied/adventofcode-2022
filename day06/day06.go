// https://adventofcode.com/2022/day/6

package day06

import (
	"fmt"
)

func ComputeResult(text string) (start int) {
	fmt.Println("ComputeResult")
	start = 0
	for loop := 1; loop < len(text); loop++ {
		for idx := start; idx < loop; idx++ {
			if (text[idx] == text[loop]) {
				start = idx
				//fmt.Printf("duplicate at %d: %c\n", start, text[start])
			}
		}
		if (len(text[start:loop]) == 4) {
			start = loop + 1
			break
		}
	}
	return start
}
