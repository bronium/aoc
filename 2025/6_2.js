const fs = require('node:fs');
const input = fs.readFileSync('./input.txt', 'utf8');

const OPERATIONS = {
  '+': (op1, op2) => op1 + op2,
  '*': (op1, op2) => op1 * op2
}

const data = input.split('\n')
const commands = data.pop()

const numbers = []
let result = 0
for (let i = commands.length - 1; i >= 0; i--) {
  let buffer = ''
  for (line of data) {
    buffer += line[i]
  }
  numbers.push(+buffer)

  if (commands[i] === '+' || commands[i] === '*') {
    result += numbers.reduce((acc, num) => OPERATIONS[commands[i]](acc, num))

    numbers.length = 0
    i--
  }
}

console.log(result)
