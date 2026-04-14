def roman_to_int(s):
    values = {
        'I': 1, 'V': 5, 'X': 10, 'L': 50,
        'C': 100, 'D': 500, 'M': 1000
    }

    s = s.upper()
    result = 0
    i = 0

    while i < len(s):
        if s[i] not in values:
            print("Ошибка: недопустимый символ", s[i])
            return -1

        current = values[s[i]]

        if i + 1 < len(s):
            next_value = values[s[i + 1]]

            if current < next_value:
                result += next_value - current
                i += 2
                continue

        result += current
        i += 1

    return result


input_str = input("Введите римское число: ")
answer = roman_to_int(input_str)

if answer != -1:
    print("Результат:", answer)