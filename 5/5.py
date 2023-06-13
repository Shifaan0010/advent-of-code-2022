from collections import defaultdict
import re

with open('5.txt', 'r') as file:
    stacks = defaultdict(list)

    for i in range(10):
        line = file.readline()

        for i in range(10):
            index = i * 4 + 1

            if index >= len(line):
                break
            
            # print(line[index], end=' ')

            if line[index].isalpha():
                stacks[i + 1].append(line[index])

        # print()
    
    for i in stacks:
        stacks[i] = stacks[i][::-1]
    
    print(*sorted(stacks.items()), sep='\n')
    
    for i, line in enumerate(file):
        move = re.match(r'move (\d+) from (\d+) to (\d+)', line)

        if move:
            count, from_index, to_index = map(int, move.groups())

            print(count, from_index, to_index)

            # for i in range(count):
            #     stacks[to_index].append(stacks[from_index].pop())

            stacks[from_index], stacks[to_index] = stacks[from_index][:-count], stacks[to_index] + stacks[from_index][-count:]
        else:
            print(line)

    print(*sorted(stacks.items()), sep='\n')

