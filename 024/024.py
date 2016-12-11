from itertools import permutations

for i, p in enumerate(permutations(range(10))):
    if i == 1000000 - 1:
        print(''.join(map(str, p)))
        break
