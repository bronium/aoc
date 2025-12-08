const fs = require('node:fs');
const input = fs.readFileSync('./input.txt', 'utf8');

const matrix = input.split('\n').map(l => l.split(''))
const coordinates = new Set([matrix[0].findIndex(el => el ==='S')])

let counter = 0
for (let i = 1; i < matrix[0].length; i++) {
  for (coord of coordinates) {
    if (matrix[i][coord] === '^') {
      counter++
      coordinates.add(coord - 1)
      coordinates.add(coord + 1)
      coordinates.delete(coord)
    }
  }
}

console.log(counter)
