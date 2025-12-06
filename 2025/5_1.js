const fs = require('node:fs');
const input = fs.readFileSync('./input.txt', 'utf8');

let [fresh, available] = input.split('\n\n')
fresh = fresh.split('\n').map((range) => range.split('-').map(el => +el))
available = available.split('\n').map(el => +el)


let counter = 0
for (fruit of available) {
    for (let [start, end] of fresh) {
        if (fruit >= start && fruit <= end) {
            counter++
            break
        }
    }
}

console.log(counter)
