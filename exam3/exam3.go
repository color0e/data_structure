package main

import "fmt"

type fruits struct {
	fruit []string
}

func (f *fruits) input(num int) {
	var fruit string
	for i := 0; i < num; i++ {
		fmt.Print("fruit[", i, "]:")
		fmt.Scanln(&fruit)
		f.fruit = append(f.fruit, fruit)
	}
}

func (f *fruits) print() {
	for _, fruit := range f.fruit {
		fmt.Println(fruit)
	}
}

func main() {
	var num int
	f := &fruits{}
	fmt.Print("fruit num:")
	fmt.Scanln(&num)
	f.input(num)
	f.print()

}
