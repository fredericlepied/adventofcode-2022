package stack

import "testing"

func TestStack(t *testing.T){
	
	var stack Stack
	stack.Push("&")
	stack.Push("&")
	stack.Push("&")
	stack.Pop()
	got := len(stack)
	want := 2
	
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
