package utils

import (
	"fmt"
)

type LinkedList[T comparable] struct {
	Data T
	Next *LinkedList[T]
}

func NewLinkedList[T comparable](data T) *LinkedList[T] {
	l := new(LinkedList[T])

	l.Data = data
	l.Next = nil

	return l
}

func (self *LinkedList[T]) Len() (count int) {
	count++
	for node := self; node.Next != nil; node = node.Next {
		count++
	}
	return count
}

func (self *LinkedList[T]) InsertNext(node *LinkedList[T]) *LinkedList[T] {
	node.Next = self.Next
	self.Next = node
	return node
}

func (self *LinkedList[T]) Display() {
	node := self
	for ; node.Next != nil; node = node.Next {
		fmt.Print(node.Data, "->")
	}
	fmt.Println(node.Data)
}

// https://leetcode.com/problems/remove-duplicates-from-sorted-list-ii/
func (self *LinkedList[T]) RemoveDuplicates() *LinkedList[T] {
	prev := self
	first := self

	if self.Next != nil && self.Data == self.Next.Data {
		first = nil
	}

	for node := self; node != nil; node = node.Next {
		// detecting first duplicate
		if node.Next != nil && node.Data == node.Next.Data {
			next := node.Next
			for ; next.Next != nil && next.Data == node.Data; next = next.Next {

			}
			prev.Next = next
			node = prev
		} else {
			// if there were duplicates at the beginning
			if first == nil {
				first = node
			}
			prev = node
		}
	}
	return first
}
