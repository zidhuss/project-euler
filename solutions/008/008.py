def product_list(numbers):
    product = 1
    for number in numbers:
        if number == 0:  # Stop on first 0
            return 0
        product *= number
    return product


def largest_adjacent_product(digits, size):
    largest = product_list(digits[:size])
    for i in range(size + 1, len(digits)):
        current = product_list(digits[i-size:i])
        if current > largest:
            largest = current
    return largest


with open('./resources/digits.txt', 'r') as number_file:
    all_numbers = number_file.read()

all_numbers = all_numbers.replace('\n', '')
print(largest_adjacent_product([int(i) for i in all_numbers], 13))
