const fs = require('node:fs');
const input = fs.readFileSync('./input.txt', 'utf8');

const coordinates = input.split('\n').map(line => line.split(',').map(el => +el))

const verticals = []
function populateVerticals() {
    const addIfVertical = (p1, p2) => {
        if (p1[0] === p2[0]) {
            const bigger = p1[1] > p2[1] ? p1[1] : p2[1]
            const smaller = p1[1] > p2[1] ? p2[1] : p1[1]
            verticals.push([p1[0], smaller, bigger])
        }
    }

    addIfVertical(coordinates[coordinates.length - 1], coordinates[0])
    for (let i = 0; i < coordinates.length - 1; i++) {
        addIfVertical(coordinates[i], coordinates[i + 1])
    }
}
populateVerticals()

const horizontals = []
function populateHorizontals() {
    const addIfHorizontal = (p1, p2) => {
        if (p1[1] === p2[1]) {
            const bigger = p1[0] > p2[0] ? p1[0] : p2[0]
            const smaller = p1[0] > p2[0] ? p2[0] : p1[0]
            horizontals.push([p1[1], smaller, bigger])
        }
    }

    addIfHorizontal(coordinates[coordinates.length - 1], coordinates[0])
    for (let i = 0; i < coordinates.length - 1; i++) {
        addIfHorizontal(coordinates[i], coordinates[i + 1])
    }
}
populateHorizontals()

function isHorizontalInsideShape(x1, y1, x2, y2) {
    if (y1 !== y2) console.error("SOMETHING IS WRONG")

    const baseline = y1
    const filtered_verticals = verticals.filter((vertical) => {
        return baseline >= vertical[1] && baseline <= vertical[2]
    }).sort((vert1, vert2) => vert1[0] - vert2[0])

    const clean_verticals = filtered_verticals.filter((vertical, ind) => {
        if (ind === 0 || ind === filtered_verticals.length - 1) return true

        return vertical[1] !== baseline && vertical[2] !== baseline
    })

    const separators = clean_verticals.map((vertical) => vertical[0])

    for (let i = 0; i < separators.length - 1; i += 2) {
        if (x1 >= separators[i] && x2 <= separators[i + 1]) return true
    }

    return false
}

function isVerticalInsideShape(x1, y1, x2, y2) {

    if (x1 !== x2) console.error("SOMETHING IS WRONG")

    const baseline = x1
    const filtered_horizontals = horizontals.filter((horizontal) => {
        return baseline >= horizontal[1] && baseline <= horizontal[2]
    }).sort((vert1, vert2) => vert1[0] > vert2[0])

    const clean_horizontals = filtered_horizontals.filter((horizontal, ind) => {
        if (ind === 0 || ind === filtered_horizontals.length - 1) return true

        return horizontal[1] !== baseline && horizontal[2] !== baseline
    })

    const separators = clean_horizontals.map((vertical) => vertical[0])

    for (let i = 0; i < separators.length - 1; i += 2) {
        if (y1 >= separators[i] && y2 <= separators[i + 1]) return true
    }

    return false
}

function isEmptyInside(minX, maxX, minY, maxY) {
    for (coord of coordinates) {
        if (coord[0] > minX && coord[0] < maxX && coord[1] > minY && coord[1] < maxY) return false
    }
    return true
}

function validArea([x1, y1], [x2, y2]) {
    if (x1 === x2 || y1 === y2) return true

    let minX, maxX, minY, maxY
    minX = Math.min(x1, x2)
    maxX = Math.max(x1, x2)
    minY = Math.min(y1, y2)
    maxY = Math.max(y1, y2)

    if (!isEmptyInside(minX, maxX, minY, maxY)) return false

    if (!isHorizontalInsideShape(minX, maxY, maxX, maxY)) return false
    if (!isHorizontalInsideShape(minX, minY, maxX, minY)) return false

    if (!isVerticalInsideShape(minX, minY, minX, maxY)) return false
    if (!isVerticalInsideShape(maxX, minY, maxX, maxY)) return false

    return true
}

let max = 0
for (let i = 0; i < coordinates.length - 1; i++) {
    for (let j = i + 1; j < coordinates.length; j++) {
        const width = Math.abs(coordinates[i][0] - coordinates[j][0]) + 1
        const heigth = Math.abs(coordinates[i][1] - coordinates[j][1]) + 1
        const area = width * heigth

        if (area > max) {
            if (validArea(coordinates[i], coordinates[j])) {
                max = area
            }
        }
    }
}

console.log(max)