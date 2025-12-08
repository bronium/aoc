const fs = require('node:fs');
const input = fs.readFileSync('./input.txt', 'utf8');

const matrix = input.split('\n').map(l => l.split(''))

const indices = new Array(matrix[0].length).fill(0)
indices[matrix[0].findIndex(el => el ==='S')] = 1

for (let i = 1; i < matrix[0].length; i++) {
  for (let ind = 0; ind < indices.length; ind++) {
    const value = indices[ind]
    if (value > 0) {
      if (matrix[i][ind] === '^') {
        indices[ind - 1] += value
        indices[ind + 1] += value
        indices[ind] = 0
      }
    }
  }
}

console.log(indices.reduce((acc, el) => acc + el))
