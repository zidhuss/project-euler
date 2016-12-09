const fs = require('fs')

const names = fs.readFileSync('./resources/names.txt').toString().split(',')
  .map(x => x.replace(/\"/g, '')).sort()

let sum = 0
for (let i = 0; i < names.length; i++) {
  sum += names[i].split('').map(x => x.charCodeAt(0) - 64)
    .reduce((a, b) => a + b) * (i + 1)
}
console.log(sum)
