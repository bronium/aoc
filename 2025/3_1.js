const fs = require('node:fs');
const input = fs.readFileSync('./input.txt', 'utf8');

const data = input.split('\n').map((bank) => {
    return bank.split('').map((digit) => {
        return parseInt(digit)
    })
})

let sum = 0
data.forEach((bank) => {
    let largest = 0
    for (let i = 0; i < bank.length - 1; i++) {
        for (let j = i + 1; j < bank.length; j++) {
            const current = bank[i] * 10 + bank[j]
            if (current > largest) { largest = current }
        }
    }
    sum += largest
})

console.log(sum)