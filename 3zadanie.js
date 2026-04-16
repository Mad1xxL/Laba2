const readline = require("readline");

function digitSum(x) {
    let sum = 0;
    while (x > 0) {
        sum += x % 10;
        x = Math.floor(x / 10);
    }
    return sum;
}

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

process.stdout.write("Введите количество символов N: ");

let tokens = [];

rl.on("line", (line) => {
    tokens.push(...line.trim().split(/\s+/));

    // если уже есть N
    if (tokens.length >= 1) {
        let N = Number(tokens[0]);

        // если уже ввели все числа — сразу закрываем
        if (Number.isInteger(N) && tokens.length >= N + 1) {
            rl.close();
        }
    }
});

rl.on("close", () => {
    let N = Number(tokens[0]);

    if (!Number.isInteger(N) || N <= 0) {
        console.log("Ошибка ввода");
        return;
    }

    if (tokens.length !== N + 1) {
        console.log("Ошибка ввода");
        return;
    }

    let count = 0;

    for (let i = 1; i <= N; i++) {
        let x = Number(tokens[i]);

        if (!Number.isInteger(x) || x <= 0) {
            console.log("Ошибка ввода");
            return;
        }

        if (digitSum(x) > 10) {
            count++;
        }
    }

    console.log("Количество чисел у которых сумма цифр больше 10:", count);
});