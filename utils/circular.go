package utils

import (
	"fmt"
)

type CircularList[T comparable] struct {
	Data T
	Next *CircularList[T]
	Prev *CircularList[T]
}

func NewCircularList[T comparable](data T) *CircularList[T] {
	c := new(CircularList[T])

	c.Data = data
	c.Next = c
	c.Prev = c

	return c
}

func NewCircularElement[T comparable](data T) *CircularList[T] {
	c := new(CircularList[T])

	c.Data = data
	c.Next = c
	c.Prev = c

	return c
}

func (c *CircularList[T]) InsertNext(node *CircularList[T]) *CircularList[T] {
	if node.Next != node {
		node.Remove()
	}
	node.Prev = c
	node.Next = c.Next
	c.Next.Prev = node
	c.Next = node
	if c.Prev == c {
		c.Prev = node
	}
	return node
}

func (c *CircularList[T]) InsertPrev(node *CircularList[T]) *CircularList[T] {
	if node.Next != node {
		node.Remove()
	}
	node.Next = c
	node.Prev = c.Prev
	c.Prev.Next = node
	c.Prev = node
	if c.Next == c {
		c.Next = node
	}
	return node
}

func (c *CircularList[T]) GetNext(count int) (ret *CircularList[T]) {
	ret = c
	for idx := 0; idx < count; idx++ {
		ret = ret.Next
	}
	return ret
}

func (c *CircularList[T]) GetPrev(count int) (ret *CircularList[T]) {
	ret = c
	for idx := 0; idx < count; idx++ {
		ret = ret.Prev
	}
	return ret
}

func (c *CircularList[T]) Len() (count int) {
	count++
	for node := c; node.Next != c; node = node.Next {
		count++
	}
	return count
}

func (c *CircularList[T]) Display() {
	node := c
	for ; node.Next != c; node = node.Next {
		fmt.Print(node.Data, "->")
	}
	fmt.Println(node.Data)
}

func (c *CircularList[T]) DisplayReverse() {
	node := c
	for ; node.Prev != c; node = node.Prev {
		fmt.Print(node.Data, "->")
	}
	fmt.Println(node.Data)
}

func (self *CircularList[T]) Remove() *CircularList[T] {
	next := self.Next
	prev := self.Prev
	prev.Next = next
	next.Prev = prev
	// list of 2
	if next.Next == self {
		next.Next = next
		next.Prev = next
		return self
	}
	self.Next = self
	self.Prev = self
	return self
}

func (self *CircularList[T]) Find(data T) *CircularList[T] {
	for node := self; node.Next != self; node = node.Next {
		if node.Data == data {
			return node
		}
	}
	return nil
}
