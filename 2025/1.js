const fs = require('node:fs');
const input = fs.readFileSync('./input.txt', 'utf8');

let position = 50
let counter = 0
let counter_part_one = 0
const instructions = input.split('\n')

instructions.forEach((instr) => {
    const dir = instr[0]
    const num = parseInt(instr.slice(1))

    const original_position = position

    if (dir === 'L') {
        position -= num
        if (position === 0) {
            counter++
        } else if (position < 0) {
            let increment = (Math.abs(Math.floor((position - 1) / 100)))
            if (original_position === 0) {
                increment -= 1
            }

            counter += increment
        }

    } else if (dir === 'R') {
        position += num
        counter += Math.abs(Math.floor(position / 100))
    }

    position = (position % 100 + 100) % 100

    if (position === 0) {
        counter_part_one++
    }
})

console.log("Part 1: ", counter_part_one)
console.log("Part 2: ", counter)
