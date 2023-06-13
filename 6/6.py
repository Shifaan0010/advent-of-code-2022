from collections import Counter

with open('6.txt', 'r') as file:
    counts = Counter()

    chars = file.read()