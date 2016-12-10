is_palindrome = lambda s: str(s) == str(s)[::-1]
sum = 0
for i in range(999999, 0, -2):
    if is_palindrome(i) and is_palindrome(bin(i)[2:]):
        sum += i

print(sum)
