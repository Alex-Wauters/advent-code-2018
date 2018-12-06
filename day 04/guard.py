import re

with open('input.txt') as f:
    lines = f.readlines()
lines.sort()

starts_shift = r'^.\d{4}-\d{2}-\d{2} (\d{2}):(\d{2}). Guard #(\d+) '
wakes_up = r'^.\d{4}-\d{2}-\d{2} (\d{2}):(\d{2}). wakes up'
falls_asleep = r'^.\d{4}-\d{2}-\d{2} (\d{2}):(\d{2}). falls asleep'

d = dict()


def get_minute(g_id, minute):
    try:
        d[g_id]
    except KeyError:
        d[g_id] = {minute: 0}
    try:
        return d[g_id][minute]
    except KeyError:
        d[g_id][minute] = 0
    return 0


def sum_max_maxcount(g_id):
    sum = 0
    max = 0
    maxid = 0
    for m in d[g_id]:
        sum = sum + d[g_id][m]
        if d[g_id][m] > max:
            max = d[g_id][m]
            maxid = m
    return sum, maxid, d[g_id][maxid]


last_asleep = 60
guard_id = -1
for l in lines:
    if re.match(starts_shift, l):
        if guard_id != -1:
            for k in range(last_asleep, 60):
                d[guard_id][k] = get_minute(guard_id, k) + 1
        matches = re.findall(starts_shift, l)
        last_asleep = 60
        guard_id = matches[0][2]
    elif re.match(wakes_up, l):
        matches = re.findall(wakes_up, l)
        for j in range(last_asleep, int(matches[0][1])):
            d[guard_id][j] = get_minute(guard_id, j) + 1
        last_asleep = 60
    elif re.match(falls_asleep, l):
        matches = re.findall(falls_asleep, l)
        last_asleep = int(matches[0][1])


max_minutes = 0
max_guard_id = -1
max_minute = 0
p2_guard_id = -1
p2_minute_id = -1
p2_minute_count = 0
for g in d:
    x_s, x_min, x_count = sum_max_maxcount(g)
    if x_s > max_minutes:
        max_minutes = x_s
        max_minute = x_min
        max_guard = g
    if x_count > p2_minute_count:
        p2_minute_id = x_min
        p2_guard_id = g
        p2_minute_count = x_count


print("Part 1", int(max_guard) * max_minute)
print("Part 2", int(p2_guard_id) * p2_minute_id)





