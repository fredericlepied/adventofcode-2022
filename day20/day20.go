// https://adventofcode.com/2022/day/20

package day20

import (
	"fmt"
	. "github.com/fredericlepied/adventofcode-2022/utils"
)

var debug = true

func b(str string) []byte {
	return []byte(str)

}

func ComputeResult(lines []string) (result int) {
	fmt.Println("ComputeResult")

	numbers := NewCircularList(ToInt(lines[0]))
	list := []*CircularList[int]{numbers}
	node := numbers
	l := len(lines)
	for idx := 1; idx < l; idx++ {
		node = node.InsertNext(NewCircularElement(ToInt(lines[idx])))
		list = append(list, node)
	}

	l = numbers.Len()
	for idx := 0; idx < len(list); idx++ {
		n := list[idx]
		if n.Data > 0 {
			node := n.GetNext(n.Data % l)
			node.InsertNext(n)
		} else if n.Data < 0 {
			node := n.GetPrev((-n.Data) % l)
			node.InsertPrev(n)
		}
	}

	zero := numbers.Find(0)
	fmt.Println("length", l, "zero", zero)
	fmt.Println(zero.GetNext(1000%l).Data, zero.GetNext(2000%l).Data, zero.GetNext(3000%l).Data)
	return zero.GetNext(1000%l).Data + zero.GetNext(2000%l).Data + zero.GetNext(3000%l).Data
}
