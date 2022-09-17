package main

import "fmt"

func main() {
	input := []string{"Mumbai", "Hyderabad", "Calicut", "Chennai"}
	k := 9
	var value int
	count := 0

	for i := 0; i < len(input); i++ {
		value = len(input[i])
		if value >= k {
			count++
		}
	}
	fmt.Println(count)
}
