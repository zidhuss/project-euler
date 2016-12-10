def sum_divisors(n):
    dsum = 0
    i = 2
    while i*i < n:
        if n % i == 0:
            dsum += i + (n // i)
        i += 1
    if i*i == n:
        dsum += i
    return dsum + 1


abundant = [i for i in range(12, 28113) if sum_divisors(i) > i]

sums = set([i+j for i in abundant for j in abundant if i+j < 28124])

print(sum([i for i in range(28124) if i not in sums]))
