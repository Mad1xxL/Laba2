package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	values := map[rune]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}

	result := 0
	runes := []rune(s)

	for i := 0; i < len(runes); i++ {
		current := values[runes[i]]

		next := 0
		if i+1 < len(runes) {
			next = values[runes[i+1]]
		}

		if current < next {
			result += next - current
			i++
		} else {
			result += current
		}
	}

	return result
}

func main() {
	var input string
	fmt.Print("Введите римское число: ")
	fmt.Scanln(&input)

	input = strings.ToUpper(input)

	fmt.Println("Результат:", romanToInt(input))
}