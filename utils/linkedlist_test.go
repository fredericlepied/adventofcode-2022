package utils

import (
	"fmt"
	"testing"
)

func TestLinkedList(t *testing.T) {

	list := NewLinkedList(1)
	fmt.Println("list")
	list.Display()
	assertEqual("Len", list.Len(), 1, t)

	list.InsertNext(NewLinkedList(2)).InsertNext(NewLinkedList(3)).InsertNext(NewLinkedList(3)).InsertNext(NewLinkedList(4)).InsertNext(NewLinkedList(4)).InsertNext(NewLinkedList(5))
	fmt.Println("node")
	list.Display()
	assertEqual("Len2", list.Len(), 7, t)

	cleaned := list.RemoveDuplicates()
	fmt.Println("cleaned")
	cleaned.Display()
	assertEqual("Len3", cleaned.Len(), 3, t)
}
