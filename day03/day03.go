// https://adventofcode.com/2022/day/3

package lib

import (
	"fmt"
	"strings"
)

func ComputeChar(chr byte) (res int) {
	if (chr >= 'a' && chr <= 'z'){
		return int(chr) - 'a' + 1
	} else if (chr >= 'A' && chr <= 'Z'){
		return int(chr) - 'A' + 27
	}
	return 0
}

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
		sum += ComputeChar(chr)
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

func ComputeLines2(lines []string) (res int) {
	sum := 0
	line := lines[0]
	for i := 0; i < len(line); i++ {
		slice := line[i:i+1]
		if (strings.Contains(lines[1], slice) && strings.Contains(lines[2], slice) ) {
			chr := line[i]
			sum += ComputeChar(chr)
			break
		}
	}
	return sum	
}

func ComputeScore2(elts []string) (res int) {
	fmt.Println("ComputeScore2")
	score := 0
	for i := 0; i < len(elts) / 3; i++ {
		score += ComputeLines2(elts[3*i:3*i+3])		
	}
	return score
}
