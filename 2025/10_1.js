const fs = require('node:fs');
const input = fs.readFileSync('./input.txt', 'utf8');

const lines = input.split('\n').map((line) => {
    const parts = line.split(' ')
    const lights_raw = parts.shift()
    const joltage_raw = parts.pop()

    const lights = Array(lights_raw.length - 2)
    for (let i = 1; i < lights_raw.length - 1; i++) {
        if (lights_raw[i] === '#') lights[i - 1] = 1
        if (lights_raw[i] === '.') lights[i - 1] = 0
    }
    const buttons = parts.map((el) => el.substring(1, el.length - 1).split(',').map((num) => +num))
    const joltage = joltage_raw.substring(1, joltage_raw.length - 1).split(',').map((num) => +num)

    return { lights, buttons, joltage }
})

lines.forEach((line) => {
    let min_presses = Infinity
    function pressButton(lights, visited, ind, presses) {
        console.log(lights)
        if (presses >= min_presses) return

        for (button of line.buttons[ind]) {
            lights[button] = lights[button] === 1 ? 0 : 1
        }

        if (lights.every((el) => el === 1)) {
            min_presses = presses
            return
        }


        if (visited.has(lights)) return
        visited.add(lights)

        for (let i = 0; i < line.buttons.length; i++) {
            pressButton(lights, visited, i, presses + 1)
        }

        return
    }

    for (let i = 0; i < line.buttons.length; i++) {
        console.log("Pressing button ", line.buttons[i])
        pressButton([...line.lights], new Set([line.lights]), i, 1)
    }

    console.log(line)
    console.log("Min presses: ", min_presses)
    console.log("=============================\n")
})

