def main():
    longest = 0
    result = 0

    for a in range(-1000, 1001):
        for b in range(-1000, 1001):
            n = 0
            while is_prime(n ** 2 + a * n + b):
                n += 1

            if n > longest:
                longest = n
                result = a * b

    print(result)


def is_prime(n: int):
    if n <= 1:
        return False

    if n <= 3:
        return True

    if (n % 2 == 0) or (n % 3 == 0):
        return False

    i = 5
    while i * i <= n:
        if (n % i == 0) or (n % (i + 2) == 0):
            return False
        i += 6
    return True


if __name__ == "__main__":
    main()
