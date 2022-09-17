package main

import "fmt"

func main() {
	num := [7]string{"Mumbai", "Delhi", "Australia", "Nigeria", "USA", "London", "Canada"}
	k := 2
	//fmt.Println(num)

	var s = num[k-1:]
	fmt.Println(s)

}
