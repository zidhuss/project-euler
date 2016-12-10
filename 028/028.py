a = lambda n: n * (n + 2) // 4 + (n % 4) // 3 + 1

n = 1001
sum = 0
i = 1
while a(i) <= n*n:
    sum += a(i)
    i += 1

print(sum)
