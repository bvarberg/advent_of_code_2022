package part1

import (
	"aoc05/stack"
	"fmt"
)

func Run() {
	s := stack.NewStringStack("A", "B", "C")
	fmt.Println(s)

	s.Push("D")
	fmt.Println(s)

	s.Pop()
	s.Pop()
	fmt.Println(s)

	top, ok := s.Top()
	if !ok {
		panic("No top crate")
	}
	fmt.Println(top)
	fmt.Println(s)

	s.Pop()
	s.Pop()
	fmt.Println(s)
}
