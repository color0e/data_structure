package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Eval("3 + 2"))
	fmt.Println(Eval("3 + ( 2 * 2 )"))
	fmt.Println(Eval("3 * ( 2 + 7 )"))
}

func Eval(expr string) int {
	ops := Stack2{}
	nums := Stack{}

	reduce := func(higher string) {
		for len(ops) > 0 {
			op := ops.last()
			if strings.Index(higher, op) < 0 {
				return
			}
			ops.pop()
			if op == "(" {
				return
			}
			b, a := nums.pop(), nums.pop()
			switch op {
			case "+":
				nums.push(a + b)
			case "-":
				nums.push(a - b)
			case "*":
				nums.push(a * b)
			case "/":
				nums.push(a / b)
			}
		}
	}

	for _, token := range strings.Split(expr, " ") {
		switch token {
		case "(":
			ops.push(token)
		case "+", "-":
			reduce("+-*/")
			ops.push(token)
		case "*", "/":
			reduce("*/")
			ops.push(token)
		case ")":
			reduce("+-*/(")
		default:
			num, _ := strconv.Atoi(token)
			nums.push(num)
		}
	}
	reduce("+-*/")
	return nums[0]
}
