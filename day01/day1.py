#!/usr/bin/python3

from collections import Counter

# open(0) for stdin
pairs = [list(map(int, line.split('   '))) for line in open('input.txt').read().splitlines()]
s = [sorted([p[i] for p in pairs]) for i in [0, 1]]
print("p1: ", sum([abs(p[0] - p[1]) for p in zip(s[0], s[1])]))

l = [p[0] for p in pairs]
r = Counter(p[1] for p in pairs)
print("p2: ", sum([n*r[n] for n in l]))