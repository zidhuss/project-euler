def largest_prime_factor(number):
    factor = 1
    i = 3
    while i * i <= number:
        if number == 1:
            return factor
        elif number % i != 0:
            i += 2
            continue
        else:
            factor = i
            while number % i == 0:
                number //= i
    return number

print(largest_prime_factor(600851475143))
