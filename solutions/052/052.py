def is_same_digits_range(n):
    is_same_digits = lambda a, b: sorted(str(a)) == sorted(str(b))
    for i in range(2, 7):
        if not is_same_digits(n, n*i):
            return False
    return True


i = 1
while not is_same_digits_range(i):
    i += 1
print(i)
