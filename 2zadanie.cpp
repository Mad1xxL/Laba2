#include <iostream>
#include <string>
#include <map>

int romanToInt(std::string s) {
    std::map<char, int> values = {
        {'I', 1}, {'V', 5}, {'X', 10}, {'L', 50},
        {'C', 100}, {'D', 500}, {'M', 1000}
    };

    int result = 0;

    for (int i = 0; i < s.size(); i++) {
        if (values.find(s[i]) == values.end()) {
            std::cout << "Ошибка: недопустимый символ " << s[i] << std::endl;
            return -1;
        }

        int current = values[s[i]];

        if (i + 1 < s.size()) {
            if (values.find(s[i + 1]) == values.end()) {
                std::cout << "Ошибка: недопустимый символ " << s[i + 1] << std::endl;
                return -1;
            }

            int next = values[s[i + 1]];

            if (current < next) {
                result += next - current;
                i++;
                continue;
            }
        }

        result += current;
    }

    return result;
}

int main() {
    std::string input;
    std::cout << "Введите римское число (только заглавные буквы): ";
    std::cin >> input;

    int answer = romanToInt(input);

    if (answer != -1) {
        std::cout << "Результат: " << answer << std::endl;
    }

    return 0;
}