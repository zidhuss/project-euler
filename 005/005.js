function gcd(a, b) {
  while (b !== 0) {
    c = b
    b = a % b
    a = c
  }
  return a
}

function lcm(a, b) {
  return a * b / gcd(a, b)
}

let result = 1
for (let i = 1; i < 20; i++)
  result = lcm(result, i)
console.log(result)
