def count(s, n):
    for c in s:
        if s.count(c) == n:
            return True
    return False

with open('input.txt') as f:
    lines = f.readlines()
twos = 0
threes = 0
for l in lines:
    twos += count(l, 2)
    threes += count(l, 3)
print(twos * threes)


