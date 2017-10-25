import math


def find_triplet_product(s):
    r = int(math.sqrt(s))
    for m in range(r, 2, -1):
        for n in range(m, 1, -1):
            a = m ** 2 - n ** 2
            b = 2 * (m * n)
            c = (m * m) + (n * n)
            if a + b + c == s:
                print(a*b*c)
                return


find_triplet_product(1000)
