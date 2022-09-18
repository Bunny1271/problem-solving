package main

import (
	"fmt"
	"unicode"
)

func main() {
	input := "aKHIL rEDDY MARRIsdsds"
	final := 0
	current := 0
	var i int
	var char int32

	start := 0
	end := 0
	for i, char = range input {
		if unicode.IsLetter(char) {
			if unicode.IsUpper(char) {
				current += 1
				continue
			}

			resetCounts(&current, &final, &start, &end, i)
			continue
		}
		resetCounts(&current, &final, &start, &end, i)
	}
	if final < current {
		final = current
		start = i - final
		end = i
	}
	fmt.Println(final)
	fmt.Println(input[start:end])
}

func resetCounts(current, final, start, end *int, i int) {
	if *current > 0 {
		if *final < *current {
			*final = *current
			*start = i - *final
			*end = i
		}
		*current = 0
	}
}
