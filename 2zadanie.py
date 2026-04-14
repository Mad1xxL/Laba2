def roman_to_int(s: str) -> int:
    values = {
        'I': 1, 'V': 5, 'X': 10, 'L': 50,
        'C': 100, 'D': 500, 'M': 1000
    }

    result = 0
    i = 0

    while i < len(s):
        current = values[s[i]]
        next_val = values[s[i + 1]] if i + 1 < len(s) else 0

        if current < next_val:
            result += next_val - current
            i += 2
        else:
            result += current
            i += 1

    return result


input_str = input("Введите римское число: ").upper()
print("Результат:", roman_to_int(input_str))