import sys
import re
from functools import cache
from collections import defaultdict

MATERIALS = ('ore', 'clay', 'obsidian', 'geode')
TOTAL_TIME = 24
ROBOTS = (1, 0, 0, 0)
RESOURCES = (0, 0, 0, 0)

inp = sys.stdin.read()

blueprints = []

blueprint_regex = re.compile(r'Blueprint (\d+):')
material_regexes = {
    'ore': re.compile(r'Each ore robot costs (?P<ore>\d+) ore.'),
    'clay': re.compile(r'Each clay robot costs (?P<ore>\d+) ore.'),
    'obsidian': re.compile(r'Each obsidian robot costs (?P<ore>\d+) ore and (?P<clay>\d+) clay.'),
    'geode': re.compile(r'Each geode robot costs (?P<ore>\d+) ore and (?P<obsidian>\d+) obsidian.')
}

while match := re.search(blueprint_regex, inp):
    inp = inp[match.end():]

    blueprint = {}

    for material, regex in material_regexes.items():
        match = re.search(regex, inp)
        inp = inp[match.end():]

        b = defaultdict(int, {k: int(v) for k, v in match.groupdict().items()})

        blueprint[material] = tuple([b[mat] for mat in MATERIALS])

    blueprints.append(tuple(blueprint[mat] for mat in MATERIALS))

print(*blueprints, sep='\n')


def add(a, b):
    return tuple(a[i] + b[i] for i, _ in enumerate(a))


def mul(a, b):
    if type(b) == tuple:
        return tuple(a[i] * b[i] for i, _ in enumerate(a))
    elif type(b) == int:
        return tuple(a[i] * b for i, _ in enumerate(a))
    else:
        raise ValueError("abc")


def next_states(rem_time, robots, resources, blueprint):
    if rem_time == 0:
        return

    yield (rem_time - 1, robots, add(resources, robots))

    for i, _ in enumerate(MATERIALS):
        rem_resources = add(resources, mul(blueprint[i], -1))

        if all(res >= 0 for res in rem_resources):
            new_resources = add(rem_resources, robots)
            new_robots = add(robots, tuple(
                0 if idx != i else 1 for idx in range(len(robots))))

            yield (rem_time - 1, new_robots, new_resources)

def next_states_pruned(rem_time, robots, resources, blueprint):
    if rem_time == 0:
        return

    yield (rem_time - 1, robots, add(resources, robots))

    for i, _ in enumerate(MATERIALS):
        # skip if we have enough robots to create next type of robot
        if i < 3 and all(res >= 0 for res in add(robots, mul(blueprint[i+1], -1))):
            continue

        rem_resources = add(resources, mul(blueprint[i], -1))

        if all(res >= 0 for res in rem_resources):
            new_resources = add(rem_resources, robots)
            new_robots = add(robots, tuple(
                0 if idx != i else 1 for idx in range(len(robots))))
            
            yield (rem_time - 1, new_robots, new_resources)

def max_geodes(total_time, start_robots, start_resources, blueprint):
    @cache
    def calc_max_geodes(rem_time, robots, resources):
        if rem_time == 0:
            # print(rem_time, robots, resources)
            return resources[3]

        max_collected = max(calc_max_geodes(*state)
                            for state in next_states_pruned(rem_time, robots, resources, blueprint))

        return max_collected

    return calc_max_geodes(total_time, start_robots, start_resources)


def max_geodes_bfs(total_time, start_robots, start_resources, blueprint):
    start = (total_time, start_robots, start_resources)

    max_collected = [0] * (total_time + 1)

    queue = [start]
    while queue:
        rem_time, robots, resources = queue.pop(0)

        if resources[3] > max_collected[rem_time]:
            max_collected[rem_time] = resources[3]
            print(max_collected)

        for state in next_states(rem_time, robots, resources, blueprint):
            # no need to check if state has been visited since graph is a tree
            queue.append(state)

    return max_collected[0]


for blueprint in blueprints:
    print(max_geodes(TOTAL_TIME, ROBOTS, RESOURCES, blueprint))
# print(max_geodes_bfs(TOTAL_TIME, ROBOTS, RESOURCES, blueprints[0]))
