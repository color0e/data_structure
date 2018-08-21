package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	ExampleWriteto()
	ExampleReadfrom()
}

func ExampleWriteto() {
	lines := []string{
		"limjangsoon@naver.com",
		"limjangsoon92@gmail.com",
		"limjangsoon@nate.com",
	}
	if err := Writeto(os.Stdout, lines); err != nil {
		fmt.Println(err)
	}
}

func ExampleReadfrom() {
	r := strings.NewReader("ptl\ncolor0e\nlimjangsoon")
	var lines []string
	if err := Readfrom(r, &lines); err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
}
