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


amicable_sum = 0
for a in range(2, 1000):
    b = sum_divisors(a)
    if b > a:
        if sum_divisors(b) == a:
            amicable_sum += a + b

print(amicable_sum)
