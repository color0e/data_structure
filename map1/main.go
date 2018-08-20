package main

import (
	"fmt"
	"sort"
)

func main() {
	codeCount := map[rune]int{}
	count("가나다라나바다라가가가가", codeCount)
	var keys sort.IntSlice
	for key := range codeCount {
		keys = append(keys, int(key))
	}
	sort.Sort(keys)
	for _, key := range keys {
		fmt.Println(string(key), codeCount[rune(key)])
	}
}

func count(str string, cC map[rune]int) {
	for _, v := range str {
		cC[v]++
	}
}
