const fs = require('fs')
const arrayProduct = (a, b) => a * b

function largestAdjacentProduct(digits, size) {
  let largest = digits.slice(0, size).reduce(arrayProduct)
  for (let i = size; i < digits.length; i++) {
    const current = digits.slice(i - size, i).reduce(arrayProduct)
    if (current > largest)
      largest = current
  }
  return largest
}

const digits = fs.readFileSync('./resources/digits.txt').toString()
  .replace(/\n/g, '').split('')
console.log(largestAdjacentProduct(digits, 13))
