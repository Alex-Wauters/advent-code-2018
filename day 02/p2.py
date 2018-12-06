with open('input.txt') as f:
    lines = f.readlines()
for l1 in lines:
    for l2 in lines:
        if sum(c1 != c2 for c1, c2 in zip(l1, l2)) == 1:
            print(l1 + '\n' + l2)

