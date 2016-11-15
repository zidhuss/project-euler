function largestPrimeFactor(number) {
  let factor = 1
  let result = number
  for (let i = 2; i * i <= result; i++) {
    if (result === 1)
      return factor
    else if (result % i !== 0)
      continue
    else {
      factor = i
      while (result % i === 0)
        result /= i
    }
  }
  return result
}

console.log(largestPrimeFactor(600851475143))
