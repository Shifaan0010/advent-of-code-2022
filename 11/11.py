import sys
import re
from functools import reduce
import math

class Monkey:
    def __init__(self, id, items, expr, test, neighbor1, neighbor2):
        self.id = id
        self.items = items
        # self.next_items = []
        self.expr = expr
        self.test = test
        self.neighbor1 = neighbor1
        self.neighbor2 = neighbor2

        self.throw_count = 0
    
    def __str__(self):
        return f'{self.id} {self.items} {self.throw_count} {self.expr} {self.test} {self.neighbor1.id} {self.neighbor2.id}'

    def operation(self, n):
        op = self.expr[1]
        a = n if self.expr[0] == 'old' else int(self.expr[0])
        b = n if self.expr[2] == 'old' else int(self.expr[2])

        if op == '+':
            return a + b
        elif op == '*':
            return a * b
        else:
            raise ValueError

    def catch_item(self, item):
        self.items.append(item)
    
    def throw_items(self, mod):
        for item in self.items:
            # new_worry_level = self.operation(item) // 3
            # monkey_to_throw = self.neighbor1 if new_worry_level % self.test == 0 else self.neighbor2
            new_worry_level = self.operation(item) % mod
            monkey_to_throw = self.neighbor1 if new_worry_level % self.test == 0 else self.neighbor2
            monkey_to_throw.catch_item(new_worry_level)

            self.throw_count += 1
        
        self.items = []
    
    # def next_round(self):
    #     self.items = self.next_items
    #     self.next_items = []

    @classmethod
    def read(cls):
        id = re.findall(r'\d+', input())[0]
        items = [int(s) for s in re.findall(r'\d+', input())]
        op = re.findall(r'(\w+|\*|\+)', input().split('=')[-1])
        test = int(re.findall(r'\d+', input())[0])
        neighbor1 = int(re.findall(r'\d+', input())[0])
        neighbor2 = int(re.findall(r'\d+', input())[0])

        # print(id, items, op, test, neighbor1, neighbor2)

        return cls(id, items, op, test, neighbor1, neighbor2)

monkeys = []
while sys.stdin:
    try:
        monkeys.append(Monkey.read())
        input()
    except EOFError:
        break

for monkey in monkeys:
    monkey.neighbor1 = monkeys[monkey.neighbor1]
    monkey.neighbor2 = monkeys[monkey.neighbor2]

print(*monkeys, sep='\n', end='\n\n')

mod = math.lcm(*(monkey.test for monkey in monkeys))

for round in range(10000):
    for monkey in monkeys:
        monkey.throw_items(mod)
    # for monkey in monkeys:
    #     monkey.next_round()

    # print()
    # print(f'Round {round + 1}')
    # print(*monkeys, sep='\n')
    # print(f'Round {round + 1}:', *(monkey.throw_count for monkey in monkeys))

shenanigans = reduce(int.__mul__, sorted([monkey.throw_count for monkey in monkeys], reverse=True)[:2])
print()
print(f'Shenanigans = {shenanigans}')