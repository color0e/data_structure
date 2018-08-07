package main

import (
	"fmt"
)

func main() {
	answer := solution("124313")
	fmt.Println(answer)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func solution(s string) int {
	// [실행] 버튼을 누르면 출력 값을 볼 수 있습니다.
	max := 1
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i:j+1] == Reverse(s[i:j+1]) {
				if max < (j - i + 1) {
					max = (j - i + 1)
				}
			}
		}
	}

	return max
}
