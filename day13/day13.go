// https://adventofcode.com/2022/day/13

package day13

import (
	"encoding/json"
	"fmt"
	"sort"
)

func compare(a, b any) int {
	as, aok := a.([]any)
	bs, bok := b.([]any)

	switch {
	case !aok && !bok:
		return int(a.(float64) - b.(float64))
	case !aok:
		as = []any{a}
	case !bok:
		bs = []any{b}
	}

	for i := 0; i < len(as) && i < len(bs); i++ {
		if c := compare(as[i], bs[i]); c != 0 {
			return c
		}
	}
	return len(as) - len(bs)
}

func ComputeResult(lines []string) (number int) {
	fmt.Println("ComputeResult")
	for idx := 0; idx <= len(lines)/3; idx++ {
		var left, right any
		json.Unmarshal([]byte(lines[3*idx]), &left)
		json.Unmarshal([]byte(lines[3*idx+1]), &right)
		fmt.Println(idx+1, left, right)
		if compare(left, right) <= 0 {
			fmt.Println("right order", idx+1)
			number += idx + 1
		}
		fmt.Println()
	}
	fmt.Println("=>", number)
	fmt.Println()
	return number
}

func ComputeResult2(lines []string) (number int) {
	fmt.Println("ComputeResult2")
	pkts := []any{}
	for idx := 0; idx <= len(lines)/3; idx++ {
		var left, right any
		json.Unmarshal([]byte(lines[3*idx]), &left)
		json.Unmarshal([]byte(lines[3*idx+1]), &right)
		pkts = append(pkts, left)
		pkts = append(pkts, right)
	}
	pkts = append(pkts, []any{[]any{2.}}, []any{[]any{6.}})
	sort.Slice(pkts, func(i, j int) bool { return compare(pkts[i], pkts[j]) < 0 })

	number = 1
	for i, p := range pkts {
		if fmt.Sprint(p) == "[[2]]" || fmt.Sprint(p) == "[[6]]" {
			number *= i + 1
		}
	}
	return number
}
