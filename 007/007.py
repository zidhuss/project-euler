n = 10001
primes = [2]
current = 3
while len(primes) < n:
    for prime in primes:
        if current % prime == 0:
            break
        elif prime ** 2 > current:
            primes.append(current)
            break
    current += 2
print(primes[-1])
