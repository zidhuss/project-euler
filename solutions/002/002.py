fibo = [0, 1]
for i in range(2, 35):
    fibo.append(fibo[i - 2] + fibo[i - 1])
print(sum(filter(lambda x: x % 2 == 0, fibo)))
