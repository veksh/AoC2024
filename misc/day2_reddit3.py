with open('2024_02.txt') as f:
    reports = [[int(num) for num in line.split()] for line in f.read().splitlines()]
    
def is_monotone(inputlist):
    return inputlist == sorted(inputlist) or inputlist[::-1] == sorted(inputlist)

def is_gradual(inputlist):
    return all(1<=abs(i-j)<=3 for i,j in zip(inputlist[:-1],inputlist[1:]))

def is_safe(inputlist):
    return is_gradual(inputlist) and is_monotone(inputlist)

answer = sum(any(is_safe(report[:i]+report[i+1:]) for i in range(len(report))) for report in reports)
print(answer)
