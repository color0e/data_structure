package main

import "fmt"

func main() {
	num := 10
	fmt.Println("[before] num:", num)
	pointer(&num)
	fmt.Println("[after] num:", num)

  num2:=1.5
  square(&num2)
  fmt.Println(num2)

  x:=1
  y:=2
  fmt.Println("[before] x:",x," y:",y)
  swap(&x,&y)
  fmt.Println("[after] x:",x," y:",y)
}

func pointer(num *int) {
	*num = 20
}

func square(x *float64) {
   *x = *x * *x
}

func swap(x *int,y *int){
  temp:=*x
  *x=*y
  *y=temp
}
