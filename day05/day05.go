// https://adventofcode.com/2022/day/5

package day05

import (
	"fmt"
	stack "github.com/fredericlepied/adventofcode-2022/stack"
)

func ComputeResult(orders [][]int, stacks []stack.Stack) (res string) {
	fmt.Println("ComputeResult")
	for idx := 0; idx < len(orders); idx++ {
		order := orders[idx]
		for number := 0; number < order[0]; number++ {
			item, _ := stacks[order[1] - 1].Pop()
			stacks[order[2] - 1].Push(item)
		}
	}
	res = ""
	for idx := 0; idx < len(stacks); idx++ {
		if (stacks[idx].IsEmpty()) {
			res += " "
		} else {
			item, _ := stacks[idx].Pop()
			res += item
		}
	}
	return res
}

func ComputeResult2(orders [][]int, stacks []stack.Stack) (res string) {
	fmt.Println("ComputeResult2")
	for idx := 0; idx < len(orders); idx++ {
		order := orders[idx]
		var local_stack stack.Stack
		for number := 0; number < order[0]; number++ {
			item, _ := stacks[order[1] - 1].Pop()
			local_stack.Push(item)
		}
		for number := 0; number < order[0]; number++ {
			item, _ := local_stack.Pop()
			stacks[order[2] - 1].Push(item)
		}
	}
	res = ""
	for idx := 0; idx < len(stacks); idx++ {
		if (stacks[idx].IsEmpty()) {
			res += " "
		} else {
			item, _ := stacks[idx].Pop()
			res += item
		}
	}
	return res
}
