const limit = 2000000
const sievebound = Math.floor((limit - 1) / 2)
const crosslimit = Math.floor(Math.sqrt(limit) - 1) / 2
const sieve = new Array(sievebound).fill(false)

for (let i = 1; i < crosslimit; i++)
  if (!sieve[i])
    for (let j = 2 * i * (i + 1); j < sievebound; j += 2 * i + 1)
      sieve[j] = true

sum = 2
for (let i = 1; i < sievebound; i++)
  if (!sieve[i])
    sum += 2 * i + 1

console.log(sum)
