import math

limit = 2000000
sievebound = (limit - 1) // 2
crosslimit = int((math.sqrt(limit) - 1) / 2)
sieve = [False for i in range(2, sievebound+2)]

for i in range(1, crosslimit):
    if not sieve[i]:
        # 2*i+1 is prime, mark mulitples
        for j in range(2*i*(i+1), sievebound, 2*i+1):
            sieve[j] = True

s = 2
for i in range(1, sievebound):
    if not sieve[i]:
        s += 2*i+1

print(s)
