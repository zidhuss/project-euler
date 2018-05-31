def gcd(a, b):
    while b != 0:
        a, b = b, a % b
    return a


num = 1
den = 1

for i in range(1, 10):
    for d in range(1, i):
        for n in range(1, d):
            a = n * 10 + i
            b = i * 10 + d
            if (n/d == a/b):
                num *= a
                den *= b

print(den // gcd(num, den))
