import re
import sys

from functools import cmp_to_key

def parse_array(s):
    stack = []

    i = 0
    while i < len(s):
        ch = s[i]

        if ch == '[':
            stack.append([])
        elif ch == ']':
            array = stack.pop()
            if stack:
                stack[-1].append(array)
            else:
                return array
        elif match := re.match(r'(\d+)', s[i:]):
            num = match.groups()[0]
            n = int(num)
            # print(n, len(num))

            stack[-1].append(n)

            i += len(num)
            continue
            
        i += 1

    return None

def order(packet1, packet2):
    i = 0

    while i < len(packet1) or i < len(packet2):
        try:
            a = packet1[i]
        except IndexError:
            return 1

        try:
            b = packet2[i]
        except IndexError:
            return -1

        if type(a) == int and type(b) == int:
            if packet1[i] < packet2[i]:
                return 1
            elif packet1[i] > packet2[i]:
                return -1
        else:
            if type(a) == int:
                a = [a]
            if type(b) == int:
                b = [b]
            
            ord = order(a, b)
            if ord != 0:
                return ord
        
        i += 1
    
    return 0

# pair = []
# index = 1

# indexSum = 0

packets = [[[2]], [[6]]]

for line in sys.stdin:
    if line.strip() == '':
        continue

    # pair.append(parse_array(line))

    # if len(pair) == 2:
    #     # print(pair)
    #     if order(pair) == 1:
    #         indexSum += index
    #         # print(index)
        
    #     pair = []
    #     index += 1

    packets.append(parse_array(line))

# print('Sum:', indexSum)

packets.sort(key=cmp_to_key(order), reverse=True)

# print(*packets, sep='\n')
key = (packets.index([[2]]) + 1) * (packets.index([[6]]) + 1)
print('Key:', key)
