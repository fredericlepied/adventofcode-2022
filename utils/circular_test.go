package utils

import (
	"fmt"
	"testing"
)

func assertEqual(prefix string, got any, want any, t *testing.T) {
	if got != want {
		t.Errorf("%s: got %d, wanted %d", prefix, got, want)
	}

}

func TestCircularList(t *testing.T) {

	circle := NewCircularList(1)
	fmt.Println()
	circle.Display()
	circle.DisplayReverse()

	node := circle.InsertNext(NewCircularElement(2))
	fmt.Println()
	circle.Display()
	circle.DisplayReverse()

	for idx := 3; idx <= 10; idx++ {
		node = node.InsertNext(NewCircularElement(idx))
		fmt.Println()
		circle.Display()
		circle.DisplayReverse()
	}

	assertEqual("Len", circle.Len(), 10, t)
	assertEqual("Prev", circle.Prev, node, t)

	node = circle.GetNext(1)
	prev := node.Prev
	node.Remove()
	fmt.Println()
	circle.Display()
	circle.DisplayReverse()
	assertEqual("Len2", circle.Len(), 9, t)

	prev.InsertNext(node)
	fmt.Println()
	circle.Display()
	circle.DisplayReverse()
	assertEqual("Len3", circle.Len(), 10, t)
}

func TestCircularListReverse(t *testing.T) {

	circle := NewCircularList(10)
	fmt.Println()
	circle.Display()
	circle.DisplayReverse()

	node := circle.InsertPrev(NewCircularElement(9))
	fmt.Println()
	circle.Display()
	circle.DisplayReverse()

	for idx := 8; idx >= 1; idx-- {
		node = node.InsertPrev(NewCircularElement(idx))
		fmt.Println()
		circle.Display()
		circle.DisplayReverse()
	}

	assertEqual("Len", circle.Len(), 10, t)
	assertEqual("Prev", circle.Next, node, t)
}
