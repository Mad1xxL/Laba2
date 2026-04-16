def digit_sum(x):
    s = 0
    while x > 0:
        s += x % 10
        x //= 10
    return s


try:
    print("Введите количество символов N:", end=" ")
    N = int(input())

    if N <= 0:
        print("Ошибка ввода")
    else:
        numbers = list(map(int, input().split()))

        if len(numbers) != N:
            print("Ошибка ввода")
        else:
            count = 0
            ok = True

            for x in numbers:
                if x <= 0:
                    print("Ошибка ввода")
                    ok = False
                    break

                if digit_sum(x) > 10:
                    count += 1

            if ok:
                print("Количество чисел у которых сумма цифр больше 10:", count)

except:
    print("Ошибка ввода")