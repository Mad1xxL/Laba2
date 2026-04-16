package main

import "fmt"

func digitSum(x int) int {
	sum := 0
	for x > 0 {
		sum += x % 10
		x /= 10
	}
	return sum
}

func main() {
	var N int

	fmt.Print("Введите количество символов N: ")
	if _, err := fmt.Scan(&N); err != nil || N <= 0 {
		fmt.Print("Ошибка ввода")
		return
	}

	count := 0

	for i := 0; i < N; i++ {
		var x int
		if _, err := fmt.Scan(&x); err != nil || x <= 0 {
			fmt.Print("Ошибка ввода")
			return
		}

		if digitSum(x) > 10 {
			count++
		}
	}

	fmt.Println("Количество чисел у которых сумма цифр больше 10:", count)
}