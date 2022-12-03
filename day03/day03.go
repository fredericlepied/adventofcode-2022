// https://adventofcode.com/2022/day/2

package lib

import (
	"fmt"
	"strings"
)

func ComputeLine(line string) (res int) {
	sum := 0
	l := len(line) / 2
	second_part := line[l:]
	for i := 0; i < l; i++ {
		if (! strings.Contains(second_part, line[i:i+1])) {
			continue
		}
		if (strings.Contains(line[:i], line[i:i+1])) {
			continue
		}
		chr := line[i]
		if (chr >= 'a' && chr <= 'z'){
			sum += int(chr) - 'a' + 1
		} else if (chr >= 'A' && chr <= 'Z'){
			sum += int(chr) - 'A' + 27
		}
	}
	return sum	
}

func ComputeScore(elts []string) (res int) {
	fmt.Println("ComputeScore")
	score := 0
	for idx := range elts {
		score += ComputeLine(elts[idx])		
	}
	return score
}
