const readline = require("readline");

function isBinary(s) {
    if (s.length === 0) return false;
    for (let c of s) {
        if (c !== '0' && c !== '1') return false;
    }
    return true;
}

const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

rl.question("Введите бинарную строчку: ", (s) => {
    if (!isBinary(s)) {
        console.log("Ошибка: строка должна содержать только 0 и 1!");
        rl.close();
        return;
    }

    rl.question("Введите количество 0, которые поменяются на 1: ", (kInput) => {
        let k = Number(kInput);

        if (!Number.isInteger(k) || k < 0) {
            console.log("k должен быть положительным числом!");
            rl.close();
            return;
        }

        let left = 0, zeros = 0, maxLen = 0;

        for (let right = 0; right < s.length; right++) {
            if (s[right] === '0') zeros++;

            while (zeros > k) {
                if (s[left] === '0') zeros--;
                left++;
            }

            maxLen = Math.max(maxLen, right - left + 1);
        }

        console.log("Максимальная длина:", maxLen);
        rl.close();
    });
});