def roman_to_int(s):
    values = {
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000
    }

    result = 0
    repeat = 1
    i = 0

    while i < len(s):
        if s[i] not in values:
            print("Ошибка: недопустимый символ")
            return -1

        if i > 0 and s[i] == s[i - 1]:
            repeat += 1

            if repeat > 4:
                print("Ошибка: недопустимый формат ввода")
                return -1
        else:
            repeat = 1

        current = values[s[i]]

        if i + 1 < len(s):
            if s[i + 1] not in values:
                print("Ошибка: недопустимый символ")
                return -1

            next_value = values[s[i + 1]]

            if current < next_value:
                result += next_value - current
                i += 2
                repeat = 1
                continue

        result += current
        i += 1

    return result


input_str = input("Введите римское число: ")
answer = roman_to_int(input_str)

if answer != -1:
    print("Результат:", answer)