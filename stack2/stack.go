package main

type Stack []int

type Stack2 []string

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

func (s *Stack) last() int {
	tmp := *s
	result := tmp[len(*s)-1]
	return result
}

func (s *Stack2) pop() string {
	if len(*s) == 0 {
		return "false"
	}
	tmp := *s
	result := tmp[len(*s)-1]
	*s = tmp[:len(*s)-1]
	return result
}

func (s *Stack2) push(data string) {
	*s = append(*s, data)
}

func (s *Stack2) last() string {
	tmp := *s
	result := tmp[len(*s)-1]
	return result
}
