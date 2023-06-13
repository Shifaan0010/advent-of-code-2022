import sys
from collections import defaultdict, deque
from functools import cache
import itertools
import heapq
import math


def log(foo):
    def bar(*args, **kwargs):
        print(
            f'{foo.__name__}({", ".join(repr(a) for a in itertools.chain(args, [f"k={kwargs[k]}" for k in kwargs]))})')
        return foo(*args, **kwargs)
    return bar


class Blizzard:
    def __init__(self, pos_dirs, height, width):
        self.positions = defaultdict(list)
        for pos in pos_dirs:
            self.positions[pos[:2]].append(pos[2])

        self.height = height
        self.width = width

        # print(self.positions)
        # print(list(self.pos_dir_iter()))

    def in_bounds(self, pos):
        r, c = pos

        return ((r == 0 and c == 1) or
                (r == self.height - 1 and c == self.width - 2) or
                ((r >= 1 and r <= self.height - 2) and
                    (c >= 1 and c <= self.width - 2)))

    def contains_blizzard_at(self, pos):
        return pos in self.positions

    def pos_dir_iter(self):
        return (
            (*pos, d)
            for pos in self.positions
            for d in self.positions[pos]
        )

    @staticmethod
    def move(pos_dir, height, width):
        r, c, d = pos_dir

        if d == '<':
            c = (c - 1 - 1) % (width - 2) + 1
        elif d == '>':
            c = (c - 1 + 1) % (width - 2) + 1
        elif d == '^':
            r = (r - 1 - 1) % (height - 2) + 1
        elif d == 'v':
            r = (r - 1 + 1) % (height - 2) + 1

        return (r, c, d)

    @classmethod
    def next_pos(cls, blizzard):
        return cls(
            (
                Blizzard.move(pos_dir, blizzard.height, blizzard.width)
                for pos_dir in blizzard.pos_dir_iter()
            ),
            blizzard.height, blizzard.width
        )

    @staticmethod
    @cache
    @log
    def at_time(blizzard, time=0):
        # print(time)
        if time == 0:
            return blizzard
        else:
            return Blizzard.next_pos(Blizzard.at_time(blizzard, time - 1))

    def __hash__(self):
        return hash((
            frozenset(self.pos_dir_iter()),
            self.height, self.width
        ))

    def __eq__(self, other):
        return (frozenset(self.pos_dir_iter()), self.height, self.width) == (frozenset(other.pos_dir_iter()), other.height, other.width)

    def __str__(self):
        grid = [
            [
                '.'
                if (r == 0 and c == 1) or
                (r == self.height - 1 and c == self.width - 2) or
                ((r >= 1 and r <= self.height - 2) and
                    (c >= 1 and c <= self.width - 2))
                else '#'

                for c in range(self.width)
            ]
            for r in range(self.height)
        ]

        for (r, c), dirs in self.positions.items():
            if len(dirs) > 0:
                ch = ''

                if len(dirs) == 1:
                    ch = dirs[0]
                elif len(dirs) < 16:
                    ch = hex(len(dirs))[2:]
                else:
                    ch = '*'

                grid[r][c] = ch

        return '\n'.join(''.join(row) for row in grid)


def neighbor_pos(pos):
    (r, c) = pos

    return (
        (r + 1, c),
        (r, c + 1),
        (r, c),
        (r, c - 1),
        (r - 1, c),
    )

def valid_neighbors(state, blizzard):
    time, (r, c) = state
    next_blizzard = Blizzard.at_time(blizzard, time + 1)

    neighbors = [(time + 1, pos) for pos in neighbor_pos(state[1])]

    return [neighbor for neighbor in neighbors
            if next_blizzard.in_bounds(neighbor[1]) and not next_blizzard.contains_blizzard_at(neighbor[1])]


def bfs(start, end, blizzard, time):
    min_time = math.inf

    queue = deque([(time, start)])
    visited = set((time, start))

    while queue:
        state = queue.popleft()

        if state[1] == end:
            min_time = state[0]
            break

        for next_state in valid_neighbors(state, blizzard):
            if next_state not in visited:
                visited.add(next_state)
                queue.append(next_state)

        # print(state, queue)

    return min_time


def dist(pos1, pos2):
    return abs(pos1[0] - pos2[0]) + abs(pos1[1] - pos2[1])


def bfs_astar(start, end, blizzard, time):
    min_time = math.inf

    prio_q = [(time + dist(start, end), time, start)]
    visited = set([(time, start)])

    while prio_q:
        h, t, pos = heapq.heappop(prio_q)

        if pos == end:
            min_time = t
            break

        for next_state in valid_neighbors((t, pos), blizzard):
            if next_state not in visited:
                visited.add(next_state)
                heapq.heappush(
                    prio_q, (next_state[0] + dist(next_state[1], end), *next_state))

        # print(state, queue)

    return min_time

# sets are much faster
def bfs2(start, end, blizzard, time):
    # states = {(time, start)}
    # visited = {(time, start)}

    # while not any(s[1] == end for s in states):
    #     time += 1
    #     blizzard = Blizzard.next_pos(blizzard)

    #     states = {(time, pos) for state in states for pos in neighbor_pos(state[1]) if blizzard.in_bounds(pos) and not blizzard.contains_blizzard_at(pos)} - visited
    #     visited |= states


    # return time, blizzard
    states = {start}

    while end not in states:
        blizzard = Blizzard.next_pos(blizzard)

        states = {s for state in states for s in neighbor_pos(state) if blizzard.in_bounds(s)} - blizzard.positions.keys()

        time += 1

    return time, blizzard


lines = sys.stdin.readlines()

height = len(lines)
width = len(lines[0]) - 1

blz_pos = [(r, c, tile) for r, line in enumerate(lines)
           for c, tile in enumerate(line)
           if tile in '<>^v']

blizzard = Blizzard(blz_pos, height, width)

start_pos = (0, 1)
end_pos = (height - 1, width - 2)

# print(height, width)

# for t in range(10):
#     print()
#     print(f'Time = {t}')
#     print(Blizzard.at_time(blizzard, t))

print('Part 1')

time_1, blz = bfs2(start_pos, end_pos, blizzard, 0)
print(f'Min time = {time_1}')

print()

print('Part 2')

time_2, blz = bfs2(end_pos, start_pos, blz, time_1)
print(time_2)
time_3, blz = bfs2(start_pos, end_pos, blz, time_2)

print(f'Min time = {time_3}')
