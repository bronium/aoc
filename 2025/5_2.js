const fs = require('node:fs');
const input = fs.readFileSync('./input.txt', 'utf8');
const fresh = input.split('\n\n')[0].split('\n').map((range) => range.split('-').map(el => +el))

for (let mainIndex = 0; mainIndex < fresh.length; mainIndex++) {
    let [mainStart, mainEnd] = fresh[mainIndex]

    for (let targetIndex = 0; targetIndex < fresh.length; targetIndex++) {
        if (!fresh[targetIndex]) continue
        if (mainIndex === targetIndex) continue;

        const [targetStart, targetEnd] = fresh[targetIndex]

        if (mainStart < targetStart) {
            if (mainEnd >= targetStart && mainEnd <= targetEnd) {
                mainEnd = targetStart - 1
                fresh[mainIndex] = [mainStart, mainEnd]
                continue
            }
        }

        if (mainStart >= targetStart && mainStart <= targetEnd) {
            if (mainEnd <= targetEnd) {
                fresh[mainIndex] = null
                break
            }

            if (mainEnd > targetEnd) {
                mainStart = targetEnd + 1
                fresh[mainIndex] = [mainStart, mainEnd]
                continue
            }
        }
    }
}

let counter = 0
fresh.filter(el => el).forEach(([start, end]) => {
    counter += end - start + 1
})

console.log("Fresh ingredient IDs: ", counter)
