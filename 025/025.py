fibo = [0, 1]
while len(str(fibo[-1])) < 1000:
    fibo.append(fibo[-2] + fibo[-1])
print(len(fibo) - 1)
