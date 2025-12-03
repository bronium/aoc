BATTERY_LENGTH = 12

function findLargest(buffer, start, bank) {
    if (buffer.toString().length === BATTERY_LENGTH) {
        return buffer
    }

    const digitMap = Array.from({ length: 10 }, () => [])

    for (let i = start; i < bank.length - (BATTERY_LENGTH - buffer.toString().length - 1); i++) {
        digitMap[bank[i]].push(i)
    }

    let ind = 0
    const indices = digitMap.findLast((arr, i) => { ind = i; return arr.length > 0 })
    if (!indices) return 0

    buffer = buffer * 10 + ind
    results = indices.map((ind) => {
        return findLargest(buffer, ind + 1, bank)
    })

    return Math.max(...results)
}

const fs = require('node:fs');
const input = fs.readFileSync('./input.txt', 'utf8');

let sum = 0
input.split('\n').forEach((bank) => {
    sum += findLargest('', 0, bank.split('').map(num => parseInt(num)))
})

console.log("Sum: ", sum)