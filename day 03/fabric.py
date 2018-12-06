import re

pattern = r'^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$'  # "#3 @ 138,913: 22x20"
with open('input.txt') as f:
    lines = f.readlines()
f = dict()
for l in lines:
    matches = re.findall(pattern, l)
    for x in range(0, int(matches[0][3])):
        for y in range(0, int(matches[0][4])):
            c = str(x + int(matches[0][1])) + "," + str(y + int(matches[0][2]))
            try:
                f[c] = f[c] + 1
            except KeyError:
                f[c] = 1
count = 0
for p in f:
    count += f[p] > 1
print("Part 1:", count)


def no_overlaps(line):
    m = re.findall(pattern, line)
    for x_o in range(0, int(m[0][3])):
        for y_o in range(0, int(m[0][4])):
            c_o = str(x_o + int(m[0][1])) + "," + str(y_o + int(m[0][2]))
            if f[c_o] > 1:
                return False
    return True


for l in lines:
    if no_overlaps(l):
        print("Part 2", l)
