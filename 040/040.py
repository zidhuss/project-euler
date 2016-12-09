s = ''
i = 1

while len(s) <= 1000000:
    s += str(i)
    i += 1

print(int(s[1-1]) * int(s[10-1]) * int(s[100-1]) * int(s[1000-1]) *
      int(s[10000-1]) * int(s[100000-1]) * int(s[1-1]))
