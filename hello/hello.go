package main

import "fmt"

func main(){
	test := [5]int{1, 2, 3, 4, 5}

	num := []int{}

	copy(test[:3], num)

	fmt.Println(num)
}