const fs = require('node:fs');
const input = fs.readFileSync('./input.txt', 'utf8');
const matrix = input.split('\n').map((row) => row.split(''))

const PAPER_ROLL = '@'
const WIDTH = matrix.length
const HEIGHT = matrix[0].length
const DIRECTIONS = [
    { x: -1, y: -1 },
    { x: 0, y: -1 },
    { x: 1, y: -1 },
    { x: -1, y: 0 },
    { x: 1, y: 0 },
    { x: -1, y: 1 },
    { x: 0, y: 1 },
    { x: 1, y: 1 },
]

function coordinateInBounds(x, y) {
    return (x >= 0 && x < WIDTH && y >= 0 && y < HEIGHT)
}

let forklift_counter = 0
for (let x = 0; x < WIDTH; x++) {
    for (let y = 0; y < HEIGHT; y++) {
        if (matrix[x][y] === PAPER_ROLL) {
            let paper_roll_counter = 0
            DIRECTIONS.forEach((direction) => {
                newX = x + direction.x
                newY = y + direction.y
                if (!coordinateInBounds(newX, newY)) return

                if (matrix[newX][newY] === PAPER_ROLL) {
                    paper_roll_counter++
                }
            })

            if (paper_roll_counter < 4) forklift_counter++
        }
    }
}

console.log("Forkliftable paper rolls: ", forklift_counter)
