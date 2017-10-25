print(max([sum([int(i) for i in str(a ** b)]) for a in range(1, 101) for b in range(1, 101)]))
