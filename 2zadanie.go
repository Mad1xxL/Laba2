package main

import "fmt"

func romanToInt(s string) int {
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	repeat := 1

	for i := 0; i < len(s); i++ {
		current, ok := values[s[i]]
		if !ok {
			fmt.Println("Ошибка: недопустимый символ")
			return -1
		}

		if i > 0 && s[i] == s[i-1] {
			repeat++

			if repeat > 4 {
				fmt.Println("Ошибка: недопустимый формат ввода")
				return -1
			}
		} else {
			repeat = 1
		}

		if i+1 < len(s) {
			next, ok := values[s[i+1]]
			if !ok {
				fmt.Println("Ошибка: недопустимый символ")
				return -1
			}

			if current < next {
				result += next - current
				i++
				repeat = 1
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