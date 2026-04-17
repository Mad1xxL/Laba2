package main

import "fmt"

func isBinary(s string) bool {
	for _, c := range s {
		if c != '0' && c != '1' {
			return false
		}
	}
	return len(s) > 0
}

func main() {
	var s string
	fmt.Print("Введите бинарную строчку: ")
	fmt.Scan(&s)

	if !isBinary(s) {
		fmt.Print("Ошибка: строка должна содержать только 0 и 1!")
		return
	}

	var k int
	fmt.Print("Введите количество 0, которые поменяются на 1: ")
	_, err := fmt.Scan(&k)

	if err == nil && k >= 0 {
		left, zeros, maxLen := 0, 0, 0

		for right := 0; right < len(s); right++ {
			if s[right] == '0' {
				zeros++
			}
			for zeros > k {
				if s[left] == '0' {
					zeros--
				}
				left++
			}
			if right-left+1 > maxLen {
				maxLen = right - left + 1
			}
		}

		fmt.Println("Максимальная длина:", maxLen)
	} else	{ fmt.Println("k должен быть положительным числом!")}
}