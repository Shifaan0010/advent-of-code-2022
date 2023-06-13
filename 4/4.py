import re

def read_ranges(line):
    return [int(s) for s in re.findall(r'\d+', line)]

def fully_contains(l1, r1, l2, r2):
    return (l1 <= l2 and r1 >= r2) or (l2 <= l1 and r2 >= r1)

def overlaps(l1, r1, l2, r2):
    return (l1 <= r2 and r1 >= l2) or (l2 <= r1 and r2 >= l1)

with open('4.txt', 'r') as file:
    print(sum(1 for line in file if overlaps(*read_ranges(line))))
