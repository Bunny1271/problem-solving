package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	scores := []int{11, 13, 101, 14, 33, 141}
	var sum int
	var total int

	if len(scores) > 11 {
		log.Println("In-vaild scores")
		os.Exit(-1)
	}
	for i := 0; i < len(scores); i++ {
		sum = scores[i]
		total = total + sum
	}
	fmt.Println(total)
}
