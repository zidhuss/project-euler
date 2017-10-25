def is_palindrome(to_check):
    return str(to_check)[::-1] == str(to_check)


def largest_palindrome_product(last, first):
    largest = 0
    for a in range(last, first, -1):
        for b in range(last, first, -1):
            calculated = a * b
            if calculated > largest and is_palindrome(calculated):
                largest = calculated
    return largest

print(largest_palindrome_product(999, 100))
