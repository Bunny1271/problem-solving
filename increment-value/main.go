package main

import (
	"fmt"
	"strconv"
)

func main() {
	num := []string{"Golang", "123", "Data", "5"}
	k := 4

	for i, value := range num {
		s, err := strconv.Atoi(value)
		if err != nil {
			continue
		}
		s += k
		//fmt.Println(s)
		num[i] = strconv.Itoa(s)
		//strconv.Itoa(s)

	}
	fmt.Println(num)

}
