function romanToInt(s) {
    const values = {
        I: 1,
        V: 5,
        X: 10,
        L: 50,
        C: 100,
        D: 500,
        M: 1000
    };

    let result = 0;

    for (let i = 0; i < s.length; i++) {
        if (!(s[i] in values)) {
            console.log("Ошибка: недопустимый символ " + s[i]);
            return -1;
        }

        let current = values[s[i]];

        if (i + 1 < s.length) {
            if (!(s[i + 1] in values)) {
                console.log("Ошибка: недопустимый символ " + s[i + 1]);
                return -1;
            }

            let next = values[s[i + 1]];

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

// Node.js
const readline = require("readline");

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

rl.question("Введите римское число (только заглавные): ", function(input) {
    let answer = romanToInt(input);

    if (answer !== -1) {
        console.log("Результат:", answer);
    }

    rl.close();
});