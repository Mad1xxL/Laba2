function romanToInt(s) {
    const values = {
        I: 1, V: 5, X: 10, L: 50,
        C: 100, D: 500, M: 1000
    };

    let result = 0;

    for (let i = 0; i < s.length; i++) {
        const current = values[s[i]];
        const next = i + 1 < s.length ? values[s[i + 1]] : 0;

        if (current < next) {
            result += next - current;
            i++;
        } else {
            result += current;
        }
    }

    return result;
}

// Ввод через Node.js
const readline = require('readline').createInterface({
    input: process.stdin,
    output: process.stdout
});

readline.question('Введите римское число: ', (input) => {
    input = input.toUpperCase();
    console.log('Результат:', romanToInt(input));
    readline.close();
});