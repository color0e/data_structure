package main

import "fmt"

func main() {
	nums := make([]int,10);

	for i,_ := range nums {
		fmt.Scan(&nums[i])
	}
  sum:=calc_sum(nums)
  avg:=sum/10
  fmt.Println(avg)

}

func calc_sum(array []int) int{
  sum:=0;
  for _,num := range array{
    sum=sum+num;
  }
  return sum
}
