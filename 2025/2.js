const fs = require('node:fs');
const input = fs.readFileSync('./input.txt', 'utf8');

const ranges = input.split(',')

let sum = 0;

function isInvalid(id) {
    let pointer = 0
    let step = 1

    while (step <= id.length / 2) {
        const part = id.slice(pointer, pointer + step)
        let valid = true
        for (let i = 0; i < id.length; i += step) {
            const current = id.slice(pointer + i, pointer + step + i)
            if (current !== part) {
                valid = false
                break
            }
        }
        if (valid) {
            return true
        }
        step++
    }

    return false
}

ranges.forEach((range) => {
    [first, last] = range.split('-')
    first = parseInt(first)
    last = parseInt(last)

    for (let i = first; i <= last; i++) {
        num = i.toString()
        if (isInvalid(num)) {
            sum += i
        }
    }
})

console.log(sum)
