#include <iostream>
#include <string>

bool isBinary(const std::string& s) {
    for (char c : s)
        if (c != '0' && c != '1') return false;
    return !s.empty();
}

int main() 
{
    std::string s;
    std::cout << "Введите бинарную строчку: ";
    std::cin >> s;

    if (!isBinary(s)) {
        std::cout << "Ошибка: строка должна содержать только 0 и 1! ";
        return 0;
    }

    int k;
    std::cout << "Введите количество 0, которые поменяются на 1: ";
    if (std::cin >> k && k >= 0) {
 
    int left = 0, zeros = 0, maxLen = 0;

    for (int right = 0; right < s.size(); right++) {
        if (s[right] == '0') zeros++;
        while (zeros > k) {
            if (s[left] == '0') zeros--;
            left++;
        }

        if (right - left + 1 > maxLen)
            maxLen = right - left + 1;
    }
    
    std::cout << "Максимальная длина: " << maxLen << std::endl;
    }
    else std::cout << "k должен быть положительным числом!" << std::endl;
    return 0;
}