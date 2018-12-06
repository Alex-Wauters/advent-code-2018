with open('input.txt') as f:
    lines = f.readlines()
d = dict({0: 1})
s = 0
while True:
    for x in lines:
        s += int(x)
        try:
            prev = d[s]
            print(s)
            quit()
        except KeyError:
            d[s] = 1
