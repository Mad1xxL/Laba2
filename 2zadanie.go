package main

import (
	"fmt"
	"strings"
)

func romanToInt(s string) int {
	values := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50,
		'C': 100, 'D': 500, 'M': 1000,
	}

	s = strings.ToUpper(s)
	result := 0

	for i := 0; i < len(s); i++ {
		current, ok := values[s[i]]
		if !ok {
			fmt.Println("Ошибка: недопустимый символ", string(s[i]))
			return -1
		}

		if i+1 < len(s) {
			next := values[s[i+1]]

			if current < next {
				result += next - current
				i++
				continue
			}
		}

		result += current
	}

	return result
}

func main() {
	var input string
	fmt.Print("Введите римское число: ")
	fmt.Scan(&input)

	answer := romanToInt(input)
	if answer != -1 {
		fmt.Println("Результат:", answer)
	}
}