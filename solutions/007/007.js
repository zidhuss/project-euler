const n = 10001
const primes = [2]

for (let x = 3; primes.length < n; x += 2)
  for (let i = 0; i < primes.length; i++)
    if (x % primes[i] === 0)
      break
    else if (primes[i] * primes[i] > x) {
      primes[primes.length] = x
      break
    }

console.log(primes[primes.length - 1])
