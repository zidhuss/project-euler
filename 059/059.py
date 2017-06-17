from operator import xor

with open("./resources/cipher.txt") as f:
    txt = [int(i) for i in f.read().strip().split(',')]

s = 0
key = 'god'
for i, x in enumerate(txt):
    s += xor(x, ord(key[i % 3]))

print(s)
