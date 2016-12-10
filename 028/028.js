const a = n =>  Math.floor(n * (n + 2) / 4) + Math.floor((n % 4) / 3) + 1

const n = 1001
let sum = 0
for (let i = 1; a(i) <= n * n; i++)
  sum += a(i)
console.log(sum)
