package main

import (
	"fmt"
	"sort"
)

func main() {
	str := "가나다라가"
	check, list := hasDupeRune(str)
	fmt.Println("중복여부:", check)
	var keys sort.IntSlice
	for key := range list {
		keys = append(keys, int(key))
	}
	sort.Sort(keys)
	for _, key := range keys {
		fmt.Println(string(key), ":", list[rune(key)])
	}
}

//글자 중복을 검사하는 함수작성
func hasDupeRune(str string) (bool, map[rune]int) {
	result := false
	list := map[rune]int{}
	check := map[rune]bool{}
	for _, v := range str {
		list[v]++
		if check[v] {
			result = true
		}
		check[v] = true
	}
	return result, list
}
