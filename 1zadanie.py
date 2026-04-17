def is_binary(s: str) -> bool:
    if len(s) == 0:
        return False
    for c in s:
        if c != '0' and c != '1':
            return False
    return True


s = input("Введите бинарную строчку: ")

if not is_binary(s):
    print("Ошибка: строка должна содержать только 0 и 1!")
    exit()

try:
    k = int(input("Введите количество 0, которые поменяются на 1: "))
except ValueError:
    print("k должен быть положительным числом!")
    exit()

if k >= 0:
    left, zeros, max_len = 0, 0, 0

    for right in range(len(s)):
        if s[right] == '0':
            zeros += 1
        while zeros > k:
            if s[left] == '0':
                zeros -= 1
            left += 1
        if right - left + 1 > max_len:
            max_len = right - left + 1

    print("Максимальная длина:", max_len)
else:
    print("k должен быть положительным числом!")