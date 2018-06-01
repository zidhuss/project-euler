function largestPrimeFactor(number) {
  let factor = 1
  let result = number
  for (let i = 3; i * i <= result; i += 2) {
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
