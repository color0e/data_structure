package main

import "fmt"

type Stack []int

func (s *Stack) pop() int {
	if len(*s) == 0 {
		return -1
	}
	tmp := *s
	result := tmp[len(*s)-1]
	*s = tmp[:len(*s)-1]
	return result
}

func (s *Stack) push(data int) {
	*s = append(*s, data)
}

func main() {
	stack := Stack{}
	stack.push(1)
	stack.push(2)
	stack.push(3)
	fmt.Println(stack)
	fmt.Println(stack.pop())
	fmt.Println(stack)
}
