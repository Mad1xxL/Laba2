#include <iostream>
#include <string>
#include <map>

int romanToInt(const std::string& s) {
    std::map<char, int> values = {
        {'I', 1}, {'V', 5}, {'X', 10}, {'L', 50}, {'C', 100}, {'D', 500}, {'M', 1000}
    };

    int result = 0;
    for (size_t i = 0; i < s.size(); ++i) {
        int current = values[s[i]];
        int next = (i + 1 < s.size()) ? values[s[i + 1]] : 0;

        if (current < next) {
            result += next - current;
            ++i; // пропускаем следующий символ
        } else {
            result += current;
        }
    }
    return result;
}

int main() {
    std::string input;
    std::cout << "Введите римское число: ";
    std::cin >> input;

    // Переводим в верхний регистр
    for (char& c : input) c = std::toupper(c);

    std::cout << "Результат: " << romanToInt(input) << std::endl;
}