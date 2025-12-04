const BATTERY_LENGTH = 12
const DIGIT_MAP = Array.from({ length: 10 }, () => [])

let cache
let bank

function clearDigitMap() {
    DIGIT_MAP.forEach((arr) => arr.length = 0)
}

function findLargest(buffer, start) {
    if (cache[buffer]) return cache[buffer]
    if (buffer.toString().length === BATTERY_LENGTH) return buffer

    clearDigitMap()

    for (let i = start; i < bank.length - (BATTERY_LENGTH - buffer.toString().length - 1); i++) {
        DIGIT_MAP[bank[i]].push(i)
    }

    let ind = 0
    const indices = DIGIT_MAP.findLast((arr, i) => { ind = i; return arr.length > 0 })
    if (!indices) return 0

    buffer = buffer * 10 + ind
    let max = 0
    indices.forEach((ind) => {
        const current = findLargest(buffer, ind + 1)
        if (current > max) {
            max = current
        }
    })

    cache[buffer] = max
    return max
}

const fs = require('node:fs');
const input = fs.readFileSync('./input.txt', 'utf8');

let sum = 0
input.split('\n').forEach((line) => {
    cache = []
    bank = line.split('').map(num => parseInt(num))
    sum += findLargest('', 0)
})

console.log("Sum: ", sum)