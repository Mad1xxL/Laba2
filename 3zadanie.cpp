#include <iostream>

int digitSum(int x) {
    int sum = 0;
    while (x > 0) {
        sum += x % 10;
        x /= 10;
    }
    return sum;
}

int main() {
    int N;
    std::cout << "Введите количество символов N: ";
    std::cin >> N;

    if (!std::cin || N <= 0) {
        std::cout << "Ошибка ввода";
        return 0;
    }

    int count = 0;

    for (int i = 0; i < N; i++) {
        int x;
        std::cin >> x;

        if (!std::cin || x <= 0) {
            std::cout << "Ошибка ввода";
            return 0;
        }

        if (digitSum(x) > 10) {
            count++;
        }
    }

    std::cout << "Количество чисел у которых сумма цифр больше 10: " << count << std::endl;
    return 0;
}