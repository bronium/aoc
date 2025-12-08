const fs = require('node:fs');
const input = fs.readFileSync('./input.txt', 'utf8');

const OPERATIONS = {
  '+': (op1, op2) => op1 + op2,
  '*': (op1, op2) => op1 * op2
}

function transpose(matrix) {
  return matrix[0].map((_, ind) => matrix.map(row => row[ind]));
}

const data = input.split('\n')
const commands = data.pop().trim().split(/\s+/)
let numbers = transpose(data.map((row) => row.trim().split(/\s+/).map((num) => +num)))

const result = commands.reduce((acc, op, ind) => {
  const start = numbers[ind].pop()
  return acc + numbers[ind].reduce((acc, value) => OPERATIONS[op](acc, value), start)
}, 0)

console.log(result)
