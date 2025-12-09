const fs = require('node:fs');
const input = fs.readFileSync('./input.txt', 'utf8');

const coords = input.split('\n').map(line => line.split(',').map(el => +el))
const SIZE = coords.length

function calculateDistance(p1, p2, p3, q1, q2, q3) {
    return Math.sqrt((p1 - q1) ** 2 + (p2 - q2) ** 2 + (p3 - q3) ** 2)
}

const distance_map = new Map()
for (let i = 0; i < SIZE - 1; i++) {
    for (let j = i + 1; j < SIZE; j++) {
        distance_map.set(calculateDistance(coords[i][0], coords[i][1], coords[i][2], coords[j][0], coords[j][1], coords[j][2]), [i, j])
    }
}

const distances = [...distance_map.keys()].sort((a, b) => a - b);

const circuits = []
for (let j = 0; j < distances.length; j++) {
    if (circuits[0] && circuits[0].size === SIZE) {
        const last_boxes = distance_map.get(distances[j - 1])
        const x1 = coords[last_boxes[0]][0]
        const x2 = coords[last_boxes[1]][0]
        console.assert(x1 * x2 === 6844224, "Result: ", x1 * x2)
        break
    }

    const connection = distance_map.get(distances[j])

    let left_connection, right_connection
    for (let i = 0; i < circuits.length; i++) {
        if (left_connection && right_connection) break

        const circuit = circuits[i]

        if (circuit.has(connection[0])) {
            left_connection = i
            continue
        }

        if (circuit.has(connection[1])) {
            right_connection = i
            continue
        }
    }


    if (left_connection >= 0 && right_connection >= 0) {
        circuits[left_connection] = circuits[left_connection].union(circuits[right_connection])
        circuits.splice(right_connection, 1)
        continue
    }

    if (left_connection >= 0) {
        circuits[left_connection].add(connection[1])
        continue
    }

    if (right_connection >= 0) {
        circuits[right_connection].add(connection[0])
        continue
    }


    circuits.push(new Set(connection))
}
