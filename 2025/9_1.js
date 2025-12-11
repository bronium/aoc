const fs = require('node:fs');
const input = fs.readFileSync('./input.txt', 'utf8');

const coordinates = input.split('\n').map(line => line.split(',').map(el => +el))

let max = 0
for (let i = 0; i < coordinates.length - 1; i++) {
    for (let j = i + 1; j < coordinates.length; j++) {
        const width = Math.abs(coordinates[i][0] - coordinates[j][0]) + 1
        const heigth = Math.abs(coordinates[i][1] - coordinates[j][1]) + 1
        const area = width * heigth

        if (area > max) max = area
    }
}

console.log(max)
