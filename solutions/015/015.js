function binomialCoefficient(n, k) {
  if (k < 0 || k > n)
    return 0
  if (k === 0 || k === n)
    return 1
  x =  n - k < k ? n - k : k
  let c = 1
  for (let i = 0; i < x; i++)
    c = c * (n - i) / (i + 1)
  return c
}

console.log(binomialCoefficient(40, 20))
