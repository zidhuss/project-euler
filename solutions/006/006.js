let sumSquare = 0
let squareSum = 0
for (let i = 1; i <= 100; i++) {
  squareSum += i * i
  sumSquare += i
}
console.log(sumSquare * sumSquare - squareSum)
