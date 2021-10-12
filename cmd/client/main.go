package main

import (
	"fmt"
	"github.com/Entrio/cfsui/internal"
)

func main() {

	s := internal.NewStack()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)
	s.Push(5)
	s.Push("A")
	s.Push("B")
	s.Push("C")
	s.Push("D")
	s.Push("E")
	s.Push("F")
	s.Push("G")
	s.Push("H")
	s.Push("I")
	i := s.Pop().(string)
	fmt.Println(i)
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()
	five := s.Pop().(int)
	fmt.Println(five)

	someValue := "this is cool"
	s.Push(&someValue)

	peeked := s.Peek().(*string)
	fmt.Println(fmt.Sprintf("Peek ptr: %d, value %s", peeked, *peeked))
	*peeked = "pew pew"

	strPtr := s.Pop().(*string)
	fmt.Println(fmt.Sprintf("Pop ptr: %d, val: %s", strPtr, *strPtr))

	//ui := internal.NewUIManager()
	//go ui.Render()
	//<-ui.WaitForExit()
}
