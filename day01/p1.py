#!/usr/bin/python3

pairs = [list(map(int, line.split('   '))) for line in open(0).read().splitlines()]
s = [sorted([p[i] for p in pairs]) for i in [0, 1]]
# print(sum([abs(p[0] - p[1]) for p in s01]))
print(s)
print(sum([abs(p[0] - p[1]) for p in zip(s[0], s[1])]))
