const fs = require('node:fs');
const input = fs.readFileSync('./input.txt', 'utf8');

const NUMBER_OF_CONNECTIONS = 1000

const coords = input.split('\n').map(line => line.split(',').map(el => +el))
const SIZE = coords.length

function calculateDistance(p1, p2, p3, q1, q2, q3) {
    return Math.sqrt((p1 - q1) ** 2 + (p2 - q2) ** 2 + (p3 - q3) ** 2)
}

const distance_map = new Map()
for (let i = 0; i < SIZE - 1; i++) {
    for (let j = i + 1; j < SIZE; j++) {
        distance_map.set(calculateDistance(...coords[i], ...coords[j]), { x: i, y: j })
    }
}

const distances = [...distance_map.keys()].sort((a, b) => a - b);
const connections = Array.from({ length: SIZE }, () => [])

for (let i = 0; i < NUMBER_OF_CONNECTIONS; i++) {
    const pair = distance_map.get(distances[i])
    connections[pair.x].push(pair.y)
    connections[pair.y].push(pair.x)
}

const visited = Array(SIZE).fill(false)
function countConnectedBoxes(boxes) {
    if (boxes.length === 0) return 1

    let box_counter = 1
    for (box of boxes) {
        if (visited[box]) continue
        visited[box] = true
        box_counter += countConnectedBoxes(connections[box])
    }

    return box_counter
}

const sizes = []
connections.forEach((connection, ind) => {
    if (visited[ind]) return

    visited[ind] = true
    sizes.push(countConnectedBoxes(connection))
})
sizes.sort((a, b) => b - a)

let acc = 1
for (let i = 0; i < 3; i++) {
    acc *= sizes[i]
}

console.log("TOTAL: ", acc)
