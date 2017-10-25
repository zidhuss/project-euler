with open('./resources/numbers.txt', 'r') as file:
    print(str(sum(([int(line.strip()) for line in file.readlines()])))[:10])
