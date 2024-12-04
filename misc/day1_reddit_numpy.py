# https://www.reddit.com/r/adventofcode/comments/1h3vp6n/2024_day_1_solutions/

from numpy import loadtxt, sort, isin

A, B = sort(loadtxt('in.txt', int).T)
print(sum(abs(A - B)),
      sum(isin(B, A) * B))
# alt
from numpy import loadtxt, sort
A, B = sort(loadtxt(filename, int).T)
print(sum(abs(A - B)),
      sum(b*np.sum(A == b) for b in B))
# alt: https://numpy.org/doc/stable/reference/generated/numpy.c_.html
print(np.sum((A==c_[B])*A))
print(((B[None, :] == A[:, None]) * B).sum())
print(sum(abs(A - B)), np.sum((A==c_[B])*A))

# also: w/o numpy
# map() can take multiple iterables: https://docs.python.org/3/library/functions.html#map
data = [*map(int, open('in.txt').read().split())]
A, B = sorted(data[0::2]), sorted(data[1::2])
print(sum(map(lambda a, b: abs(a-b), A, B)),
      sum(map(lambda a: a * B.count(a), A)))
